package main

import (
	"context"
	"log"
	"net"

	hs "go-grpc-api-lab/api/hello-world"

	"google.golang.org/grpc"
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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hs.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
