package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "github.com/davidmontoyago/go-grpc-api-lab/api/tls-auth"
	"github.com/davidmontoyago/go-grpc-api-lab/pkg/osutil"
	"github.com/davidmontoyago/go-grpc-api-lab/pkg/paseto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	certFile := osutil.GetenvOrDefault("CERT_PEM", "./certs/self-signed-cert.pem")
	creds, err := credentials.NewClientTLSFromFile(certFile, "localhost")
	if err != nil {
		log.Fatalf("unable to load tls config: %v", err)
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(
		getAddress(),
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(paseto.AuthClientInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := api.NewSecureServiceClient(conn)
	response, err := client.CheckMyCreds(context.Background(), &api.SecureRequest{Data: "this is a secured message..."})
	if err != nil {
		log.Fatalf("Error when calling CheckMyCreds: %s", err)
	}
	log.Printf("Response from server: %v", response.Success)
}

func getAddress() string {
	return fmt.Sprintf("%s:%s",
		os.Getenv("server_host"),
		osutil.GetenvOrDefault("server_port", "50051"),
	)
}
