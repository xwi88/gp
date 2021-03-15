package gp

import (
	"sync"

	"github.com/xwi88/gp/tf"
)

var gpLock sync.RWMutex
var globalModels map[string]Model

type Model interface {
	Load() error
	Predict(interface{}) (interface{}, error) // input maybe single, slice; output is the same
	Destruct() error
}

// ModelOptions used to register and config the model
type ModelOptions struct {
}

//  default param name: input, output
func RegisterTFModel(name, path string, tags []string) bool {
	RegisterTFModelWithParamName(name, path, tags, "serving_default_input", "StatefulPartitionedCall")
	return false
}

func GetModel(modelName string) Model {
	return globalModels[modelName]
}

func RegisterTFModelWithParamName(name, path string, tags []string, inputParamKey, outputParamKey string) bool {
	gpLock.Lock()
	defer gpLock.Unlock()

	if globalModels == nil {
		globalModels = make(map[string]Model)
	}
	if _, exist := globalModels[name]; exist {
		return true
	}

	m, err := tf.RegisterWithParamName(name, path, tags, inputParamKey, outputParamKey)
	if err == nil {
		globalModels[name] = m
		return true
	}
	// 	TODO register failed
	return false
}

func Predict(modelName string, input interface{}) (output interface{}, err error) {
	if m, exist := globalModels[modelName]; exist {
		output, err = m.Predict(input)
		return
	}
	// TODO log

	return
}

// Reload reload model
func Reload(modelName string) error {
	if m, exist := globalModels[modelName]; exist {
		return m.Load()
	}
	//  TODO
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
