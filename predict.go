package gp

import (
	"sync"

	"github.com/xwi88/gp/tf"
)

var gpLock sync.RWMutex
var globalModels map[string]Model

type Model interface {
	Load() error
	Predict(...interface{}) (interface{}, error) // input maybe single slice; output as the same
	Destruct() error
}

// ModelOptions used to register and config the model
type ModelOptions struct {
}

// RegisterTFModel register TFModel with default output and input param key
func RegisterTFModel(name, path string, tags []string) bool {
	return RegisterTFModelWithParamName(name, path, tags, "StatefulPartitionedCall", "serving_default_input")
}

// GetModel get model by model name
func GetModel(name string) Model {
	return globalModels[name]
}

// RegisterTFModelWithParamName register TFModel with output and variadic input param key
func RegisterTFModelWithParamName(name, path string, tags []string, outputParamKey string, inputParamKey ...string) bool {
	gpLock.Lock()
	defer gpLock.Unlock()

	if globalModels == nil {
		globalModels = make(map[string]Model)
	}
	if _, exist := globalModels[name]; exist {
		return true
	}

	m, err := tf.RegisterWithParamName(name, path, tags, outputParamKey, inputParamKey...)
	if err == nil {
		globalModels[name] = m
		return true
	}
	return false
}

// Predict give predict val with model name and variadic input data
func Predict(name string, input ...interface{}) (output interface{}, err error) {
	if m, exist := globalModels[name]; exist {
		output, err = m.Predict(input)
		return
	}
	return
}

// Reload model by name
func Reload(name string) error {
	if m, exist := globalModels[name]; exist {
		return m.Load()
	}
	return nil
}

// DestructModel release model memory
func DestructModel() {
	gpLock.Lock()
	defer gpLock.Unlock()
	for k, v := range globalModels {
		if err := v.Destruct(); err == nil {
			delete(globalModels, k)
		}
	}
}
