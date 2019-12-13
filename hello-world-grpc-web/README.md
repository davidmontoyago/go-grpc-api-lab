# hello-world-grpc-web

For Web interopetability gRPC-Web requires a proxy that can handle HTTP/1 and speak HTTP/2 to the gRPC server. This can be achieved with [envoy](https://github.com/envoyproxy/envoy/issues/6897) or the [Golang grpcwebproxy](https://github.com/improbable-eng/grpc-web/tree/master/go/grpcwebproxy).

### Pre-requisites

```
make pre-reqs
```

### Generate Js client

```
make grpc
```

### Bundle Js app

```
make webpack
```