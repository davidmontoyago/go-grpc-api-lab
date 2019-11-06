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
	$(GOBUILD) -i -v -o ./bin/hello-world/server ./hello-world/server
	$(GOBUILD) -i -v -o ./bin/hello-world/client ./hello-world/client

test:
	$(GOTEST) ./

clean:
	$(GOCLEAN)

fmt:
	$(GOCMD) fmt ./...

run-server:
	make build
	./bin/hello-world/server

run-client:
	make build
	./bin/hello-world/client

grpc:
	protoc -I api/ -I ${GOPATH}/src --go_out=plugins=grpc:api api/hello-world/hello-service.proto
