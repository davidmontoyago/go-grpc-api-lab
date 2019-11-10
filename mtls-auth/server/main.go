package main

import (
	"context"
	"log"
	"net"

	api "go-grpc-api-lab/api/tls-auth"
	mtlsutil "go-grpc-api-lab/pkg/mtlsutil"

	"google.golang.org/grpc"
)

const (
	port = ":50053"
)

// Do not check in certs to your repo! This is only for demo purposes. Inject them as env config
const (
	certFile = "./mtls-auth/certs/self-signed-cert.pem"
	keyFile  = "./mtls-auth/certs/self-signed-key.pem"
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
	creds, err := mtlsutil.NewMutualTLSServerCreds(certFile, keyFile)
	if err != nil {
		log.Fatalf("unable to load server mtls config: %v", err)
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
