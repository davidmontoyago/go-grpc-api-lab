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

// AuthClientInterceptor authenticates the client against an "auth-micro-service" and gets a JWT encrypted (JWE) token
func AuthClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	requestMetadata, _ := metadata.FromOutgoingContext(ctx)
	metadataCopy := requestMetadata.Copy()

	token := authenticateClient()
	tokenAsHeader := fmt.Sprintf("Bearer %s", token)
	metadataCopy.Set("authorization", tokenAsHeader)

	ctx = metadata.NewOutgoingContext(ctx, metadataCopy)
	return invoker(ctx, method, req, reply, cc, opts...)
}

// simulates a secure call to an "auth-micro-service" and returns a JWE token
func authenticateClient() string {
	privateKey := getEncryptingKey()
	enc := newSymetricEncrypter(privateKey)

	// claims - should get encrypted
	notBefore := time.Now().Add(time.Duration(-10) * time.Second)
	claims := jwt.Claims{
		Subject:   "hello-grpc-micro-client",
		Issuer:    "auth-micro-service",
		NotBefore: jwt.NewNumericDate(notBefore),
		Audience:  jwt.Audience{"hello-grpc-micro-server"},
	}

	token, err := jwt.Encrypted(enc).Claims(claims).CompactSerialize()
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func newSymetricEncrypter(privateKey string) jose.Encrypter {
	key := []byte(privateKey)
	enc, err := jose.NewEncrypter(
		jose.A128GCM,
		jose.Recipient{
			Algorithm: jose.DIRECT,
			Key:       key,
		},
		(&jose.EncrypterOptions{}).WithType("JWT"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return enc
}

func getEncryptingKey() string {
	privateKey, present := os.LookupEnv("PRIVATE_KEY")
	if !present {
		log.Fatalf("JWT encrypting key not present! stopping now.")
	}
	return privateKey
}
