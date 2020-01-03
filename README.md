# gRPC

[Some gRPC talk on youtube](https://www.youtube.com/watch?v=J-NTfvYL_OE):
> ...gRPC is the Linux pipe for boring distributed components...

### Workflow for creating gRPC services

1. Define service contract using protocol buffers Interface Definition Language (IDL)
2. Compile IDL (.protos) into service interfaces
3. Implement service interfaces

## Examples

| Example  	| .proto (IDL) 	| client/server impl  	|
|---	|---	|---	|
| Simple Hello World  	| `./api/hello-world/`  	| `./hello-world/`  	|
| Stream events down to multiple clients. <br/> Keeps an in-mem broker to subscribe clients and broadcast to all. | `./api/server-streaming/`  	  | `./server-streaming/`  	|
| TLS Server Auth | `./api/tls-auth/`  	  | `./tls-auth/`  	|
| Mutual TLS Auth. Uses in-mem copy of the host trustore to append self signed cert. | `./api/tls-auth/`  	  | `./mtls-auth/`  	|
| Tracing and metering client/server interceptors with OpenTelemetry. | `./api/hello-world/` | `./opentelemetry-interceptor/` |
| Server streaming to web app with gRPC-Web and grpcwebproxy. | `./api/server-streaming/` | `./server-streaming-grpc-web/` |


### Run example

``` bash
# compile protos
make grpc

go mod vendor

go install ./...

cd ./example
make run-server
make run-client
```