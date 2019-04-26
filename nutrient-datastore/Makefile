GOPATH:=$(shell go env GOPATH)

.PHONY: build
build: 
	go build -o nutrient-datastore main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t suhdev/nutrient-datastore:latest