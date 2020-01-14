package paseto

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/o1egl/paseto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthClientInterceptor authenticates the client against an "auth-micro-service" and gets a Paseto encrypted token
func AuthClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	requestMetadata, _ := metadata.FromOutgoingContext(ctx)
	metadataCopy := requestMetadata.Copy()

	token := authenticateClient()
	tokenAsHeader := fmt.Sprintf("Bearer %s", token)
	metadataCopy.Set("authorization", tokenAsHeader)

	ctx = metadata.NewOutgoingContext(ctx, metadataCopy)
	return invoker(ctx, method, req, reply, cc, opts...)
}

// simulates a secure call to an "auth-micro-service" and returns a Paseto token
func authenticateClient() string {
	symmetricKey := getEncryptingKey()

	now := time.Now()
	expiration := now.Add(24 * time.Hour)
	notBefore := now

	jsonToken := paseto.JSONToken{
		Audience:   "hello-grpc-micro-server",
		Issuer:     "auth-micro-service",
		Jti:        "12345",
		Subject:    "hello-grpc-micro-client",
		IssuedAt:   now,
		Expiration: expiration,
		NotBefore:  notBefore,
	}
	// custom claim
	jsonToken.Set("data", "some-sensitive-user-context")

	// a footer allows passing unencrypted data to the receiver. I.e.: a key id
	footer := "my-key-id"

	v2 := paseto.NewV2()
	token, err := v2.Encrypt([]byte(symmetricKey), jsonToken, footer)
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func getEncryptingKey() string {
	privateKey, present := os.LookupEnv("PRIVATE_KEY")
	if !present {
		log.Fatalf("encrypting key not present! stopping now.")
	}
	return privateKey
}
