package gp

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestDestructModel(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPredict(t *testing.T) {
	type args struct {
		modelName string
		input     []interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []interface{}
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := Predict(tt.args.modelName, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Predict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Predict() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func registerTFModel(modelName string) Model {
	exportDir := "testdata/saved_model_half_plus_two_cpu/000001"
	tags := []string{"serve"}
	if ok := RegisterTFModelWithParamName(modelName, exportDir, tags, "y", "x"); ok {
		return GetModel(modelName)
	}
	return nil
}

func Test_RegisterTFModel(t *testing.T) {
	modelName := "test"
	registerTFModel(modelName)

	rand.Seed(time.Now().UnixNano())
	inputS := generateSliceFloat2(int(rand.Int31n(20)) + 1)
	output, err := Predict(modelName, inputS)
	if err != nil {
		t.Errorf("Predict err: %v", err)
	}
	t.Logf("input:%v, output:%v", inputS, output)

}

func Benchmark_RegisterTFModel(b *testing.B) {
	modelName := "test"
	registerTFModel(modelName)

	n := b.N
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		inputS := generateSliceFloat2(int(rand.Int31n(20)) + 1)
		output, err := Predict(modelName, inputS)
		_ = output
		_ = err
	}
}
func TestReload(t *testing.T) {
	type args struct {
		modelName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Reload(tt.args.modelName); (err != nil) != tt.wantErr {
				t.Errorf("Reload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func generateSliceFloat2(size int) (s []float32) {
	rand.Seed(time.Now().UnixNano())
	s = make([]float32, size)
	for i := 0; i < size; i++ {
		s[i] = float32(rand.Int63n(200))
	}
	return s
}
