FROM golang:1.13.5-alpine3.10

# server or client mode
ENV mode server
ENV server_host ""
ENV server_port "50051"

WORKDIR /hello-world

COPY ./bin/server /hello-world
COPY ./bin/client /hello-world

CMD ["/bin/ash", "-c", "./${mode}"]