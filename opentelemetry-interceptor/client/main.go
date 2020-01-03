package main

import (
	"context"
	"log"
	"time"

	"github.com/davidmontoyago/go-grpc-api-lab/opentelemetry-interceptor/config"

	api "github.com/davidmontoyago/go-grpc-api-lab/api/hello-world"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/davidmontoyago/go-grpc-api-lab/pkg/go.opentelemetry.io/otel/grpc/metric"
	"github.com/davidmontoyago/go-grpc-api-lab/pkg/go.opentelemetry.io/otel/grpc/trace"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

func main() {
	config.InitTracer()
	defer config.InitMeter().Stop()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				trace.UnaryClientInterceptor,
				metric.UnaryClientInterceptor,
			),
		),
	)
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
