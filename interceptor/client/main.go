package main

import (
	"context"
	"log"
	"time"

	api "go-grpc-api-lab/api/hello-world"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"go-grpc-api-lab/pkg/go.opentelemetry.io/otel/grpc/tracing"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50054", grpc.WithInsecure(), grpc.WithUnaryInterceptor(tracing.UnaryClientInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewHelloServiceClient(conn)

	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "web-api-client-us-east-1",
		"user-id", "some-test-user-id",
	)
	context := metadata.NewOutgoingContext(context.Background(), md)
	response, err := c.SayHello(context, &api.HelloRequest{Greeting: "World"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Reply)
}
