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
	# install grpc-web protoc plugin
	curl -LO "https://github.com/grpc/grpc-web/releases/download/1.0.7/protoc-gen-grpc-web-1.0.7-darwin-x86_64"
	mv -f protoc-gen-grpc-web-1.0.7-darwin-x86_64 ~/bin/protoc-gen-grpc-web
	chmod 750 ~/bin/protoc-gen-grpc-web

	# install grpc-web proxy
	curl -LO https://github.com/improbable-eng/grpc-web/releases/download/v0.12.0/grpcwebproxy-v0.12.0-osx-x86_64.zip
	unzip -jo grpcwebproxy-v0.12.0-osx-x86_64.zip -d ~/bin/ && ln -sf ~/bin/grpcwebproxy-v0.12.0-osx-x86_64 ~/bin/grpcwebproxy
	chmod 750 ~/bin/grpcwebproxy
	rm grpcwebproxy-v0.12.0-osx-x86_64.zip

	# install js app deps
	cd client && npm install --save-dev webpack && cd -

grpc:
	#
	# generates:
	# client/api/server-streaming/streaming-service_grpc_web_pb.js with the service stub for the web client
	# client/api/server-streaming/streaming-service_pb.js with the js protobuf messages
	#
	protoc -I=./../api server-streaming/streaming-service.proto --js_out=import_style=commonjs:./../server-streaming-grpc-web/client/api --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./../server-streaming-grpc-web/client/api

webpack:
	cd ./client && npm install && npx webpack --mode development client.js && cd -

run-server:
	$(GOBUILD) -i -v -o ./bin/server ./../server-streaming/server
	./bin/server

run-client:
	pkill grpcwebproxy || true
	grpcwebproxy --backend_addr=localhost:50051 --run_tls_server=false --allow_all_origins &

	cd client && python3 -m http.server 8081