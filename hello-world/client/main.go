package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/davidmontoyago/go-grpc-api-lab/pkg/osutil"

	api "github.com/davidmontoyago/go-grpc-api-lab/api/hello-world"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(getAddress(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewHelloServiceClient(conn)
	response, err := c.SayHello(context.Background(), &api.HelloRequest{Greeting: "David"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Reply)
}

func getAddress() string {
	port := osutil.GetenvOrDefault("server_port", "50051")
	return fmt.Sprintf("%s:%s", os.Getenv("server_host"), port)
}
