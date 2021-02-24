package gp

import (
	"sync"

	"github.com/xwi88/gp/tf"
)

var gpLock sync.RWMutex
var globalModels map[string]Model

type Model interface {
	Load() error
	Predict(interface{}) ([]interface{}, error)
	Destruct() error
}

//  default param name: input, output
func RegisterTFModel(name, path string, tags []string) bool {
	gpLock.Lock()
	defer gpLock.Unlock()

	if globalModels == nil {
		globalModels = make(map[string]Model)
	}
	if _, exist := globalModels[name]; exist {
		return true
	}

	m, err := tf.Register(name, path, tags)
	if err == nil {
		globalModels[name] = m
		return true
	}
	// 	TODO register failed
	return false
}

func Predict(modelName string, input []interface{}) (output []interface{}, err error) {
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
