# gp

go predict with tensorflow, pytorch and others 

## Import

```bash
go get -u github.com/xwi88/gp
```

## TF

```bash
# register model
RegisterTFModelWithParamName(modelName, exportDir, tags, "x", "y")
# get model
GetModel(modelName)
# predict with the special model
output, err := Predict(modelName, inputS)
```
