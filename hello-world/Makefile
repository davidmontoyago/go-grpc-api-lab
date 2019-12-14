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
	./bin/server

run-client:
	make build-client
	./bin/client

grpc:
	protoc -I ../api/ -I ${GOPATH}/src --go_out=plugins=grpc:../api ../api/hello-world/hello-service.proto

package:
	make build-server GOOS=linux
	make build-client GOOS=linux
	docker build -t grpc-hello-world:latest .

	# run in server mode
	# docker run -it --rm --name=server -e mode=server -p 50051:50051 grpc-hello-world:latest

	# run in client mode
	# docker run -it --rm -e mode=client -e server_host=server --link=server grpc-hello-world:latest