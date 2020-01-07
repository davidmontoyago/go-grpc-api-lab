# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GO111MODULE=off
GOOS?=darwin
GOARCH=amd64

clean:
	$(GOCLEAN)

fmt:
	$(GOCMD) fmt ./...

build-server:
	$(GOBUILD) -i -v -o ./bin/server ./server

build-client:
	$(GOBUILD) -i -v -o ./bin/client ./client

run-server:
	make build-server
	SIGNING_KEY=this-is-a-secret ./bin/server

run-client:
	make build-client
	SIGNING_KEY=this-is-a-secret ./bin/client

grpc:
	protoc -I ../api/ -I ${GOPATH}/src --go_out=plugins=grpc:../api ../api/tls-auth/secure-service.proto
