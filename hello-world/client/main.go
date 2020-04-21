package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/davidmontoyago/go-grpc-api-lab/pkg/osutil"

	api "github.com/davidmontoyago/go-grpc-api-lab/api/hello-world"

	"google.golang.org/grpc"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	checkHealth()
	callAPI()
}

func getAddr(port string) string {
	return fmt.Sprintf("%s:%s", os.Getenv("server_host"), port)
}

func checkHealth() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	var err error
	var hconn *grpc.ClientConn

	apiAddr := getAddr(osutil.GetenvOrDefault("server_port", "50051"))
	hconn, err = grpc.Dial(apiAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect to health endpoint: %s", err)
	}
	defer hconn.Close()

	healthClient := healthgrpc.NewHealthClient(hconn)
	healthReq := &healthgrpc.HealthCheckRequest{Service: "api.HelloService"}
	healthResp, err := healthClient.Check(ctx, healthReq)
	if err != nil {
		log.Fatal("health check failed", err)
	}
	log.Println("service health", healthResp)
}

func callAPI() {
	ctx := context.Background()

	var err error
	var conn *grpc.ClientConn

	apiAddr := getAddr(osutil.GetenvOrDefault("server_port", "50051"))
	conn, err = grpc.Dial(apiAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect to api: %s", err)
	}
	defer conn.Close()

	apiClient := api.NewHelloServiceClient(conn)
	response, err := apiClient.SayHello(ctx, &api.HelloRequest{Greeting: "David"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Reply)
}
