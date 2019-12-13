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
| Stream data down to multiple clients. <br/> Keeps an in-mem broker to subscribe clients and broadcast to all. | `./api/client-streaming/`  	  | `./client-streaming/`  	|
| TLS Server Authentication | `./api/tls-auth/`  	  | `./tls-auth/`  	|
| Mutual TLS Authentication. Uses in-mem copy of the host trustore to append self signed cert. | `./api/tls-auth/`  	  | `./mtls-auth/`  	|
| Tracing and metering client/server interceptors with OpenTelemetry. | `./api/hello-world/` | `./opentelemetry-interceptor/` |
| Server streaming to web app with gRPC-Web and grpcwebproxy. | `./api/client-streaming/` | `./server-streaming-grpc-web/` |

### Compile all .protos

```
make grpc
```

### Run a server

``` bash
# example is the name of the example server to run
make run-server example=example-name

# for instance, run server for ./client-streaming example
make run-server example=client-streaming
```

### Run a client

``` bash
# example is the name of the example client to run
make run-client example=example-name

# for instance, run client for ./client-streaming example
make run-client example=client-streaming
```
