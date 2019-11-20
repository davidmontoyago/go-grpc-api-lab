package grpctrace

import (
	"context"
	"strings"

	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc/metadata"
)

const (
	Vendor = "ot"
)

var (
	propagator = propagation.HTTPTraceContextPropagator{}
)

type metadataSupplier struct {
	metadata *metadata.MD
}

func (s *metadataSupplier) Get(key string) string {
	values := s.metadata.Get(key)
	return strings.Join(values, ",")
}

func (s *metadataSupplier) Set(key string, value string) {
	s.metadata.Append(key, value)
}

func Inject(ctx context.Context, metadata *metadata.MD) {
	propagator.Inject(ctx, &metadataSupplier{
		metadata: metadata,
	})
}
