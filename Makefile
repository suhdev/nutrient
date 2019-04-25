define COMPILE_CMD
protoc -I proton proton/data.proto --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go --plugin=protoc-gen-micro=${GOPATH}/bin/protoc-gen-micro --micro_out=proton --go_out=proton;
endef

.PHONY: all
all: deps build

.PHONY: deps
deps:
	@glide install
	@$(COMPILE_CMD)

.PHONY: build
build:
	@$(COMPILE_CMD)

.PHONY: clean
clean:
	go clean
	go clean -cache
	rm -rf ~/.cache/go-build
	rm -rf ${GOPATH}/pkg/*/github.com/suhdev/nutrient

.PHONY: protobuf
protobuf:
	@$(COMPILE_CMD)

.PHONY: protobuf-reflex
protobuf-reflex: protobuf
	@echo "waiting for proto file changes..."
	@reflex -d none -r "\.proto$$" -- zsh -c "$(COMPILE_CMD) ; repeat 100 printf '#'"

