package main

import (
	"context"
	"log"

	mtlsutil "github.com/davidmontoyago/go-grpc-api-lab/pkg/mtlsutil"

	api "github.com/davidmontoyago/go-grpc-api-lab/api/tls-auth"

	"google.golang.org/grpc"
)

// Do not check in certs to your repo! This is only for demo purposes. Inject them as env config
const (
	certFile = "./certs/self-signed-cert.pem"
	keyFile  = "./certs/self-signed-key.pem"
)

func main() {
	creds, err := mtlsutil.NewMutualTLSClientCreds(certFile, keyFile, "localhost")
	if err != nil {
		log.Fatalf("unable to load client mtls config: %v", err)
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(":50051", grpc.WithTransportCredentials(creds))
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
