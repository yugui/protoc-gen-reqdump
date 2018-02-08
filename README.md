# protoc-gen-reqdump
A protoc plugin which simply emits the given plugin request as the output.

# How to use
```shell
$ protoc -I. --plugin=path/to/protoc-gen-reqdump --reqdump_out=. target.proto
```
