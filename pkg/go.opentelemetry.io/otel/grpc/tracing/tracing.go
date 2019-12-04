package tracing

// Open Telemetry gRPC Tracer integration
// https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/data-rpc.md
import (
	"context"
	"go-grpc-api-lab/pkg/opentelemetry-go/plugin/grpctrace"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/distributedcontext"
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/api/metric"
	"go.opentelemetry.io/otel/api/trace"

	"go.opentelemetry.io/otel/api/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor intercepts and extracts incoming trace data
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	requestMetadata, _ := metadata.FromIncomingContext(ctx)
	metadataCopy := requestMetadata.Copy()

	entries, spanCtx := grpctrace.Extract(ctx, &metadataCopy)
	ctx = distributedcontext.WithMap(ctx, distributedcontext.NewMap(distributedcontext.MapUpdate{
		MultiKV: entries,
	}))

	grpcServerKey := key.New("grpc.server")
	serverSpanAttrs := []core.KeyValue{
		grpcServerKey.String("hello-world-server"),
	}

	tr := global.TraceProvider().Tracer("example/grpc")
	ctx, span := tr.Start(
		ctx,
		"hello-api-op",
		trace.WithAttributes(serverSpanAttrs...),
		trace.ChildOf(spanCtx),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	return handler(ctx, req)
}

// UnaryClientInterceptor intercepts and injects outgoing trace
func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	requestMetadata, _ := metadata.FromOutgoingContext(ctx)
	metadataCopy := requestMetadata.Copy()

	tracer := global.TraceProvider().Tracer("example/grpc")

	err := tracer.WithSpan(ctx, "hello-api-op",
		func(ctx context.Context) error {
			grpctrace.Inject(ctx, &metadataCopy)
			ctx = metadata.NewOutgoingContext(ctx, metadataCopy)

			err := invoker(ctx, method, req, reply, cc, opts...)
			setTraceStatus(ctx, err)
			return err
		})
	return err
}

var (
	appKey       = key.New("app")
	operationKey = key.New("operation")
)

// UnaryClientMeteringInterceptor records custom app metrics
func UnaryClientMeteringInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	meter := global.MeterProvider().Meter("example/grpc")
	labels := meter.Labels(appKey.String("example/grpc/client"), operationKey.String("hello-api-op"))

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		failedCallsMetric := meter.NewInt64Counter("example.grpc.client.calls.fail", metric.WithKeys(appKey, operationKey))
		counter := failedCallsMetric.AcquireHandle(labels)
		defer counter.Release()
		counter.Add(ctx, 1)
	} else {
		successCallsMetric := meter.NewInt64Counter("example.grpc.client.calls.success", metric.WithKeys(appKey, operationKey))
		counter := successCallsMetric.AcquireHandle(labels)
		defer counter.Release()
		counter.Add(ctx, 1)
	}
	return err
}

func setTraceStatus(ctx context.Context, err error) {
	if err != nil {
		status, _ := status.FromError(err)
		trace.CurrentSpan(ctx).SetStatus(status.Code())
	} else {
		trace.CurrentSpan(ctx).SetStatus(codes.OK)
	}
}
