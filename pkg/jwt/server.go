package jwt

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/square/go-jose.v2/jwt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	errorz "github.com/pkg/errors"
)

// AuthServerInterceptor validates JWT token and claims
func AuthServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	requestMetadata, _ := metadata.FromIncomingContext(ctx)
	metadataCopy := requestMetadata.Copy()

	token, err := getBearerToken(metadataCopy)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	err = authenticateToken(token, os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// else Success!

	return handler(ctx, req)
}

func getBearerToken(md metadata.MD) (string, error) {
	bearerToken := md.Get("authorization")[0]

	parts := strings.Split(bearerToken, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1], nil
	}

	return "", errors.New("unable to parse authorization header")
}

// deserialize and validate token claims
func authenticateToken(token string, privateKey string) error {
	jwtToken, err := jwt.ParseEncrypted(token)
	if err != nil {
		return errorz.Wrap(err, "failed parsing token")
	}

	// decrypt claims
	claims := jwt.Claims{}
	if err := jwtToken.Claims([]byte(privateKey), &claims); err != nil {
		return errorz.Wrap(err, "failed to deserialize token")
	}

	// validate claims
	err = claims.Validate(jwt.Expected{
		Issuer: "auth-micro-service",
		Time:   time.Now(),
	})
	if err != nil {
		return errorz.Wrap(err, "invalid claims")
	}

	log.Printf("auth success! iss: %s, sub: %s\n", claims.Issuer, claims.Subject)
	return nil
}
