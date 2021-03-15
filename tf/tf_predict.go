package tf

import (
	"errors"
	"sync"

	tfg "github.com/tensorflow/tensorflow/tensorflow/go"
)

type Model struct {
	name    string              // memory store tf model name, for supporting multi model
	path    string              // load model use, model dir(exportDir)
	tags    []string            // load model use, model tags
	options *tfg.SessionOptions // load model use, session options
	model   *tfg.SavedModel     // load and save tf model

	inputParamKey  string // required
	outputParamKey string // required

	count int // stats: load count
	lock  sync.RWMutex
}

// New according the input params to generate the special tf model
func New(name, exportDir string, tags []string, inputParamKey, outputParamKey string) *Model {
	return &Model{
		name:           name,
		path:           exportDir,
		tags:           tags,
		inputParamKey:  inputParamKey,
		outputParamKey: outputParamKey,
	}
}

// Predict tf predict
func (m *Model) Predict(data interface{}) (ret interface{}, err error) {
	if data == nil {
		return nil, errors.New("nil input")
	}

	inputData, err := tfg.NewTensor(data)
	if err != nil {
		return nil, err
	}
	input := map[tfg.Output]*tfg.Tensor{
		m.model.Graph.Operation(m.inputParamKey).Output(0): inputData,
	}
	output := []tfg.Output{
		m.model.Graph.Operation(m.outputParamKey).Output(0),
	}
	rt, err := m.model.Session.Run(input, output, nil)
	if err != nil {
		return nil, err
	}
	ret = rt[0].Value() // WARN: only result
	return
}

// Load load tf model from special path
func (m *Model) Load() error {
	m.lock.Lock()
	defer m.lock.Unlock()

	// TODO 1. judge model file exist
	// TODO 2. check others

	tfModel, err := tfg.LoadSavedModel(m.path, m.tags, m.options)
	if err != nil {
		return err
	}
	m.model = tfModel
	m.count++
	return nil
}

// Register register and load model
func Register(name, exportDir string, tags []string) (*Model, error) {
	return RegisterWithParamName(name, exportDir, tags, "serving_default_input", "StatefulPartitionedCall")
}

// Register register and load model
func RegisterWithParamName(name, exportDir string, tags []string, inputParamKey, outputParamKey string) (*Model, error) {
	m := New(name, exportDir, tags, inputParamKey, outputParamKey)
	return m, m.Load()
}

func (m *Model) Destruct() error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m == nil {
		return nil
	}

	m.tags = nil
	m.options = nil
	m.model = nil
	return nil
}

func (m *Model) Name() string {
	return m.name
}

func (m *Model) Path() string {
	return m.path
}
