package main

import (
	"log"
	"net"

	api "github.com/davidmontoyago/go-grpc-api-lab/api/server-streaming"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service := &Server{
		broker: NewBroker(),
	}
	// Start consuming events from upstream system
	go service.ConsumeEvents()

	s := grpc.NewServer()
	api.RegisterEventStreamingServiceServer(s, service)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
