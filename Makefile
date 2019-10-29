# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GO111MODULE=off
GOOS?=darwin
GOARCH=amd64

all: test build

build:
	# go mod vendor
	$(GOBUILD) -i -v -o ./bin/server ./server

test:
	$(GOTEST) ./

clean:
	$(GOCLEAN)

fmt:
	$(GOCMD) fmt ./...

run:
	make build
	./go-grpc-api-lab

grpc:
	protoc -I api/ -I ${GOPATH}/src --go_out=plugins=grpc:api/ api/hello-service.proto
