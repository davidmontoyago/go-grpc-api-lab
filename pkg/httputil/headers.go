package httputil

import (
	"errors"
	"strings"
)

// ParseBearerToken gets a bearer token from a given authorization header
func ParseBearerToken(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1], nil
	}

	return "", errors.New("unable to parse authorization header")
}
