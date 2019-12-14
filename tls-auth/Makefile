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

run-server:
	$(GOBUILD) -i -v -o ./bin/server ./server
	./bin/server

run-client:
	$(GOBUILD) -i -v -o ./bin/client ./client
	./bin/client

grpc:
	protoc -I ../api/ -I ${GOPATH}/src --go_out=plugins=grpc:../api ../api/tls-auth/secure-service.proto
