package main

import (
	"context"
	"log"
	"net"

	api "go-grpc-api-lab/api/hello-world"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement api.HelloServiceServer
type server struct {
	api.UnimplementedHelloServiceServer
}

// SayHello implements api.HelloServiceServer
func (s *server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	log.Printf("Received: %v", in.GetGreeting())
	return &api.HelloResponse{Reply: "Hello " + in.Greeting}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
