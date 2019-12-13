# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GO111MODULE=off
GOOS?=darwin
GOARCH=amd64

pre-reqs:
	curl -L -O "https://github.com/grpc/grpc-web/releases/download/1.0.7/protoc-gen-grpc-web-1.0.7-darwin-x86_64"
	mv -f protoc-gen-grpc-web-1.0.7-darwin-x86_64 ~/bin/protoc-gen-grpc-web
	chmod 750 ~/bin/protoc-gen-grpc-web

	cd client && npm install --save-dev webpack && cd -

grpc:
	#
	# generates:
	# api/hello-world/hello-service_grpc_web_pb.js with the service stub for the web client
	# api/hello-world/hello-service_pb.js with the js protobuf messages
	#
	protoc -I=./../api hello-world/hello-service.proto --js_out=import_style=commonjs:./../hello-world-grpc-web/client/api --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./../hello-world-grpc-web/client/api

webpack:
	cd ./client && npm install && npx webpack --mode development client.js && cd -

run-server:
	$(GOBUILD) -i -v -o ./bin/server ./server
	./bin/server

run-client:
	cd proxy && docker build -t grpc-web/envoy . && cd -
	
	docker ps -f name=grpc-web-proxy -q | xargs docker rm -f
	docker run --name grpc-web-proxy -d -p 8080:8080 -p 50052:50052 --network="bridge" grpc-web/envoy
	
	cd client && python3 -m http.server 8081