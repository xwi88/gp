# gp

go predict with tensorflow, pytorch and others 

## Tensorflow

### Ref

- [tensorflow](https://github.com/tensorflow/tensorflow)
- [tensorflow-serving](https://github.com/tensorflow/serving)

### Go API

>https://www.tensorflow.org/install/lang_go

### C Lib

- [download](https://www.tensorflow.org/install/lang_c#download)
- [libtensorflow-cpu-darwin-x86_64-2.3.0](https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-darwin-x86_64-2.3.0.tar.gz)
- [libtensorflow-cpu-linux-x86_64-2.3.0](https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.3.0.tar.gz)

```bash
#export LIB_TENSORFLOW_FILE=libtensorflow-cpu-darwin-x86_64-2.3.0.tar.gz
export LIB_TENSORFLOW_FILE=libtensorflow-cpu-linux-x86_64-2.3.0.tar.gz
```

#### Standard Install

```bash
tar -C /usr/local -xzf ${LIB_TENSORFLOW_FILE}
ldconfig
```

#### Non Sys Dir Install

```bash
export LIB_TENSORFLOW_INSTALL_DIR=~/mydir
tar -C ${LIB_TENSORFLOW_INSTALL_DIR} -xzf ${LIB_TENSORFLOW_FILE}

# config linker
export LIBRARY_PATH=$LIBRARY_PATH:${LIB_TENSORFLOW_INSTALL_DIR}/lib # For both Linux and macOS X
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:${LIB_TENSORFLOW_INSTALL_DIR}/lib # For Linux only
export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:${LIB_TENSORFLOW_INSTALL_DIR}/lib # For macOS X only
```

#### C Lib Check

```cgo
#include <stdio.h>
#include <tensorflow/c/c_api.h>

int main() {
  printf("Hello from TensorFlow C library version %s\n", TF_Version());
  return 0;
}
```

- `gcc hello_tf.c -ltensorflow -o hello_tf`
- `./hello_tf`

### Go Usage

`go get github.com/xwi88/gp`

### Example

```bash
# register model
RegisterTFModelWithParamName(modelName, exportDir, tags, []string{"param_name_input"}, "param_name_output")

# predict.go
# get model
GetModel(modelName)
# predict with the special model
output, err := Predict(modelName, inputS)
```

### Params Look Up

- `pip install tensorflow`
- `saved_model_cli show --all --dir output/keras`

### Docker Images

>**tf version now will be ok in: macos & ubuntu16.04**

|image repos|target|notes|
|:--|:--|:--|
|v8fg/ubuntu:16.04-go1.16-tf-cpu|build||
|v8fg/ubuntu:16.04-tf-cpu|run||

- more docker images ref: [v8fg/docker-compose-resources](https://github.com/v8fg/docker-compose-resources)
- official images: [TensorFlow Docker](https://www.tensorflow.org/install/docker)

## Tips

- [version incompatible ref](https://github.com/tensorflow/tensorflow/issues/41808)
    - https://github.com/photoprism/photoprism/pull/775
- [hack fixed, only ref](https://github.com/tensorflow/tensorflow/blob/master/tensorflow/go/README.md)
- [saved_model_half_plus_two](https://github.com/tensorflow/serving/blob/master/tensorflow_serving/servables/tensorflow/testdata/saved_model_half_plus_two.py)
    - input, output, tags
