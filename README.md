# nutrient

The repo contains common `protobuf` definitions that are shared across micro-services. 

## Dependencies 

1. `protoc`: protocol buffers compiler.
2. `github.com/golang/protobuf/protoc-gen-go`: protocol buffers golang code generator.
3. `github.com/micro/protoc-gen-micro`: protocol buffers `go-micro` code generator. 

## Build Steps

1. Make sure `protoc` is installed, please refer to [https://developers.google.com/protocol-buffers/docs/gotutorial](Protocol Buffers for Go).
2. Do `go get -u github.com/golang/protobuf/protoc-gen-go` to install go's protobuf plugin. 
3. Do `go get -u github.com/micro/protoc-gen-micro` to install `go-micro`'s protobuf plugin.
4. Do `make build` to build the `.proto` file. 

## TODO

* Add `make` tasks to build all services and docker images at once