package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "github.com/davidmontoyago/go-grpc-api-lab/api/tls-auth"
	"github.com/davidmontoyago/go-grpc-api-lab/pkg/osutil"

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
	conn, err = grpc.Dial(getAddress(), grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewSecureServiceClient(conn)
	response, err := c.CheckMyCreds(context.Background(), &api.SecureRequest{Data: "this is a secured message..."})
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
