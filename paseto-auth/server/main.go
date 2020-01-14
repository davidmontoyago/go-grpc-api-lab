package main

import (
	"context"
	"log"
	"net"
	"os"

	api "github.com/davidmontoyago/go-grpc-api-lab/api/tls-auth"
	"github.com/davidmontoyago/go-grpc-api-lab/pkg/osutil"
	"github.com/davidmontoyago/go-grpc-api-lab/pkg/paseto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("Metadata: %v", md)

	log.Printf("Received: %v", req.Data)
	return &api.SecureResponse{Success: true}, nil
}

func main() {
	certFile := osutil.GetenvOrDefault("CERT_PEM", "./certs/self-signed-cert.pem")
	keyFile := osutil.GetenvOrDefault("KEY_PEM", "./certs/self-signed-key.pem")

	// in a Cloud Native environment secrets would be injected as env vars
	_, present := os.LookupEnv("PRIVATE_KEY")
	if !present {
		log.Fatalf("encryption key not present! stopping now.")
	}

	// Configure server TLS creds
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("unable to load tls config: %v", err)
	}

	s := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(paseto.AuthServerInterceptor),
	)
	api.RegisterSecureServiceServer(s, &Server{})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
