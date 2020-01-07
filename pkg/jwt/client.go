package jwt

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

// AuthClientInterceptor authenticates the client against an "auth-micro-service" and gets a JWT token
func AuthClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// in a Cloud Native environment secrets would be injected as env vars
	signingKey, present := os.LookupEnv("SIGNING_KEY")
	if !present {
		log.Fatalf("JWT signing key not present! stopping now.")
	}

	requestMetadata, _ := metadata.FromOutgoingContext(ctx)
	metadataCopy := requestMetadata.Copy()

	token := authenticateClient(signingKey)
	tokenAsHeader := fmt.Sprintf("Bearer %s", token)
	metadataCopy.Set("authorization", tokenAsHeader)

	ctx = metadata.NewOutgoingContext(ctx, metadataCopy)
	return invoker(ctx, method, req, reply, cc, opts...)
}

// simulates a call to an "auth-micro-service" and returns a JWT token
func authenticateClient(privateKey string) string {
	key := []byte(privateKey)
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: key}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		log.Fatal(err)
	}

	notBefore := time.Now().Add(time.Duration(-10) * time.Second)
	claims := jwt.Claims{
		Subject:   "hello-grpc-micro-client",
		Issuer:    "auth-micro-service",
		NotBefore: jwt.NewNumericDate(notBefore),
		Audience:  jwt.Audience{"hello-grpc-micro-server"},
	}
	token, err := jwt.Signed(sig).Claims(claims).CompactSerialize()
	if err != nil {
		log.Fatal(err)
	}

	return token
}
