# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GO111MODULE=off
GOOS?=darwin
GOARCH=amd64

build:
	# go mod vendor
	$(GOBUILD) -i -v -o ./bin/$(example)/server ./$(example)/server
	$(GOBUILD) -i -v -o ./bin/$(example)/client ./$(example)/client

clean:
	$(GOCLEAN)

fmt:
	$(GOCMD) fmt ./...

run-server:
	make build example=$(example)
	./bin/$(example)/server

run-client:
	make build example=$(example)
	./bin/$(example)/client

grpc:
	protoc -I api/ -I ${GOPATH}/src --go_out=plugins=grpc:api api/hello-world/hello-service.proto
	protoc -I api/ -I ${GOPATH}/src --go_out=plugins=grpc:api api/client-streaming/streaming-service.proto
