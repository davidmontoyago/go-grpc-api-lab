package main

import (
	"context"
	"log"
	"net"

	"github.com/davidmontoyago/go-grpc-api-lab/opentelemetry-interceptor/config"

	hs "github.com/davidmontoyago/go-grpc-api-lab/api/hello-world"

	"google.golang.org/grpc"

	"github.com/davidmontoyago/go-grpc-api-lab/pkg/go.opentelemetry.io/otel/grpc/trace"
)

const (
	port = ":50051"
)

// server is used to implement api.HelloServiceServer
type server struct {
	hs.UnimplementedHelloServiceServer
}

// SayHello implements api.HelloServiceServer
func (s *server) SayHello(ctx context.Context, in *hs.HelloRequest) (*hs.HelloResponse, error) {
	log.Printf("Received: %v", in.GetGreeting())
	return &hs.HelloResponse{Reply: "Hello " + in.Greeting}, nil
}

func main() {
	config.InitTracer()
	defer config.InitMeter().Stop()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(trace.UnaryServerInterceptor))

	hs.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
