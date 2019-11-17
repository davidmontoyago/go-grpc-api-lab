package main

import (
	"context"
	"log"
	"net"

	hs "go-grpc-api-lab/api/hello-world"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":50054"
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

// Telemetry Server side interceptor
func TracingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	metadata, _ := metadata.FromIncomingContext(ctx)
	log.Printf("tracing request with metadata: %v", metadata)
	// TODO emit metric pre-handler with metadata
	response, err := handler(ctx, req)
	// TODO emit metric post-handler with metadata
	return response, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(TracingInterceptor))

	hs.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
