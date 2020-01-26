package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	api "github.com/davidmontoyago/go-grpc-api-lab/api/hello-world"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	port       = ":50051"
	healthPort = ":50052"
)

// server is used to implement api.HelloServiceServer
type apiServer struct {
	api.UnimplementedHelloServiceServer
}

// SayHello implements api.HelloServiceServer
func (s *apiServer) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	log.Printf("Received: %v", in.GetGreeting())
	return &api.HelloResponse{Reply: "Hello " + in.Greeting}, nil
}

func main() {
	hServer := grpc.NewServer()
	go serveAPI(hServer)

	apiServer := grpc.NewServer()
	go serveHealth(apiServer)

	// graceful shutdown
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown
	log.Println("shutting down now...")

	hServer.GracefulStop()
	apiServer.GracefulStop()
	os.Exit(0)
}

func serveHealth(grpcServer *grpc.Server) {
	lis, err := net.Listen("tcp", healthPort)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", port, err)
	}

	hServer := health.NewServer()
	hServer.SetServingStatus("api.HelloService", healthgrpc.HealthCheckResponse_SERVING)

	healthgrpc.RegisterHealthServer(grpcServer, hServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func serveAPI(grpcServer *grpc.Server) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", port, err)
	}

	api.RegisterHelloServiceServer(grpcServer, &apiServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
