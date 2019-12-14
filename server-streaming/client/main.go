package main

import (
	"context"
	api "go-grpc-api-lab/api/server-streaming"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewEventStreamingServiceClient(conn)
	stream, err := c.GetEventStream(context.Background(), &api.EventRequest{Type: api.EventType_CREATE})
	if err != nil {
		log.Fatalf("Error when calling GetEventStream: %s", err)
	}

	for {
		event, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error receiving data: %s", err)
		}
		log.Printf("Got an Event from server: %v", event)
	}
}
