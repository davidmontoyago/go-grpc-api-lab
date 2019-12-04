package grpctrace

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/propagators"
)

const (
	// Vendor is the integration provider
	Vendor = "ot"
)

var (
	propagator = propagators.TraceContext{}
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

// Inject injects the gRPC call metadata into the Span
func Inject(ctx context.Context, metadata *metadata.MD) {
	propagator.Inject(ctx, &metadataSupplier{
		metadata: metadata,
	})
}

// Extract returns the Context Entries and SpanContext that were encoded by Inject.
func Extract(ctx context.Context, metadata *metadata.MD) ([]core.KeyValue, core.SpanContext) {
	spanContext, correlationCtx := propagator.Extract(ctx, &metadataSupplier{
		metadata: metadata,
	})

	var correlationCtxKVs []core.KeyValue
	correlationCtx.Foreach(func(kv core.KeyValue) bool {
		correlationCtxKVs = append(correlationCtxKVs, kv)
		return true
	})

	return correlationCtxKVs, spanContext
}
