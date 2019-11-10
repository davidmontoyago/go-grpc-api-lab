package main

import (
	"context"
	"log"

	api "go-grpc-api-lab/api/tls-auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Do not check in certs to your repo! This is only for demo purposes. Inject them as env config
const (
	certFile = "./tls-auth/certs/self-signed-cert.pem"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile(certFile, "localhost")
	if err != nil {
		log.Fatalf("unable to load tls config: %v", err)
	}

	var conn *grpc.ClientConn
	// conn, err := grpc.Dial(":50053", grpc.WithInsecure())
	conn, err = grpc.Dial(":50053", grpc.WithTransportCredentials(creds))
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
