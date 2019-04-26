GOPATH:=$(shell go env GOPATH)

.PHONY: build
build: 
	go build -o nutrient-business main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t suhdev/nutrient-business:latest