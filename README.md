# gp

go predict with tensorflow, pytorch and others 

## Tensorflow

### Ref

- [tensorflow](https://github.com/tensorflow/tensorflow)
- [tensorflow-serving](https://github.com/tensorflow/serving)

### Go API

>https://www.tensorflow.org/install/lang_go

### C lib install

- [download](https://www.tensorflow.org/install/lang_c#download)
- [current support:libtensorflow-cpu-darwin-x86_64-2.3.0](https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-darwin-x86_64-2.3.0.tar.gz)

```bash
# Standard Install
sudo tar -C /usr/local -xzf (downloaded file)

# Non sys dir install
tar -C ~/mydir -xzf libtensorflow-cpu-darwin-x86_64-2.3.0.tar.gz

#ldconfig
export LIBRARY_PATH=$LIBRARY_PATH:~/mydir/lib
export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:~/mydir/lib
```

### Usage

`go get -u github.com/xwi88/gp`

### Example

```bash
# register model
RegisterTFModelWithParamName(modelName, exportDir, tags, "param_name_input", "param_name_output")
# get model
GetModel(modelName)
# predict with the special model
output, err := Predict(modelName, inputS)
```

#### Params Look Up

- `pip install tensorflow`
- `saved_model_cli show --all --dir output/keras`

## Tips

- [version incompatible ref](https://github.com/tensorflow/tensorflow/issues/41808)
    - https://github.com/photoprism/photoprism/pull/775
- [hack fixed, only ref](https://github.com/tensorflow/tensorflow/blob/master/tensorflow/go/README.md)
- [saved_model_half_plus_two](https://github.com/tensorflow/serving/blob/master/tensorflow_serving/servables/tensorflow/testdata/saved_model_half_plus_two.py)
    - input, output, tags
