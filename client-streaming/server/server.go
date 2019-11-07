package main

import (
	"fmt"
	api "go-grpc-api-lab/api/client-streaming"
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(35))

// Server is used to implement api.EventStreamingServiceServer
// Keeps an in-mem broker to subscribe clients and broadcast Events to all
type Server struct {
	api.UnimplementedEventStreamingServiceServer
	broker *Broker
}

// Simulates consuming events
func (s *Server) ConsumeEvents() {
	go s.broker.Start()
	for {
		// Randomly generate events
		randomValue := random.Int()
		switch randomValue % 3 {
		case 0:
			// new event
			event := &api.Event{
				Id:                fmt.Sprintf("create-event-%d", randomValue),
				Type:              api.EventType_CREATE,
				Description:       fmt.Sprintf("sample create event %d for client", randomValue),
				SourceApplication: "upstream-event-api-us-east-1",
			}
			s.broker.Publish(event)
		case 1:
			// new event
			event := &api.Event{
				Id:                fmt.Sprintf("update-event-%d", randomValue),
				Type:              api.EventType_UPDATE,
				Description:       fmt.Sprintf("sample update event %d for client", randomValue),
				SourceApplication: "upstream-event-api-us-east-2",
			}
			s.broker.Publish(event)
		default:
			// no event
			time.Sleep(5 * time.Second)
		}
	}
}

// Subscribes the client to the broker and streams down broadcasted Events
func (s *Server) GetEventStream(req *api.EventRequest, stream api.EventStreamingService_GetEventStreamServer) error {
	msgCh := s.broker.Subscribe()
	for msg := range msgCh {
		event := msg.(*api.Event)
		if event.Type == req.Type {
			if err := stream.Send(event); err != nil {
				return err
			}
		}
	}
	return nil
}
