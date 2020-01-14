package paseto

import (
	"context"
	"log"
	"os"

	"github.com/davidmontoyago/go-grpc-api-lab/pkg/httputil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/o1egl/paseto"
	errorz "github.com/pkg/errors"
)

// AuthServerInterceptor validates JWT token and claims
func AuthServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	requestMetadata, _ := metadata.FromIncomingContext(ctx)
	metadataCopy := requestMetadata.Copy()

	authHeader := metadataCopy.Get("authorization")[0]
	token, err := httputil.ParseBearerToken(authHeader)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	err = authenticateToken(token, os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return handler(ctx, req)
}

func authenticateToken(token string, symmetricKey string) error {
	var newJSONToken paseto.JSONToken
	var newFooter string
	v2 := paseto.NewV2()

	err := v2.Decrypt(token, []byte(symmetricKey), &newJSONToken, &newFooter)
	if err != nil {
		return errorz.Wrap(err, "failed decrypting token")
	}
	return nil
}
