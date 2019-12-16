package main

import (
	"context"
	"log"
	"net"

	api "go-grpc-api-lab/api/tls-auth"
	"go-grpc-api-lab/pkg/osutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

// Server is used to implement api.SecureService
type Server struct {
	api.UnimplementedSecureServiceServer
}

// CheckMyCreds implements api.UnimplementedSecureServiceServer
func (s *Server) CheckMyCreds(ctx context.Context, req *api.SecureRequest) (*api.SecureResponse, error) {
	log.Printf("Received: %v", req.Data)
	return &api.SecureResponse{Success: true}, nil
}

func main() {
	certFile := osutil.GetenvOrDefault("CERT_PEM", "./certs/self-signed-cert.pem")
	keyFile := osutil.GetenvOrDefault("KEY_PEM", "./certs/self-signed-key.pem")

	// Configure server TLS creds
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("unable to load tls config: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	api.RegisterSecureServiceServer(s, &Server{})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
