FROM golang:1.13.5-alpine3.10

# server or client mode
ENV mode server
ENV server_host ""
ENV server_port "50051"

WORKDIR /tls-grpc-app

COPY ./bin/server /tls-grpc-app
COPY ./bin/client /tls-grpc-app
COPY ./certs /tls-grpc-app/certs

CMD ["/bin/ash", "-c", "./${mode}"]