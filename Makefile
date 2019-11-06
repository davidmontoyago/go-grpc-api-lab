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
	$(GOBUILD) -i -v -o ./bin/client ./client

test:
	$(GOTEST) ./

clean:
	$(GOCLEAN)

fmt:
	$(GOCMD) fmt ./...

run-server:
	make build
	./bin/server

run-client:
	make build
	./bin/client

grpc:
	protoc -I api/ -I ${GOPATH}/src --go_out=plugins=grpc:api/ api/hello-service.proto
