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
| Hello World with health check and graceful termination  	| `./api/hello-world/`  	| `./hello-world/`  	|
| Stream events down to multiple clients. <br/> Keeps an in-mem broker to subscribe clients and broadcast to all. | `./api/server-streaming/`  	  | `./server-streaming/`  	|
| TLS Server Auth | `./api/tls-auth/`  	  | `./tls-auth/`  	|
| Mutual TLS Auth. Uses in-mem copy of the host trustore to append self signed cert. | `./api/tls-auth/`  	  | `./mtls-auth/`  	|
| Tracing and metering client/server interceptors with OpenTelemetry. | `./api/hello-world/` | `./opentelemetry-interceptor/` |
| Server streaming to web app with gRPC-Web and grpcwebproxy [Failing]. | `./api/server-streaming/` | `./server-streaming-grpc-web/` |
| JWT (as JWE) Auth with [square/go-jose](https://github.com/square/go-jose) | `./api/tls-auth/` | `./jwt-auth/` |
| Paseto JSON token Auth with [o1egl/paseto](https://github.com/o1egl/paseto) | `./api/tls-auth/` | `./paseto-auth/` |


### Run example

``` bash
go mod vendor

go install ./...

cd ./example

# compile protos if changed
make grpc

make run-server
make run-client
```