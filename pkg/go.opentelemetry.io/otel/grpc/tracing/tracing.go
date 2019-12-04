package tracing

// Open Telemetry gRPC Tracer integration
// https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/data-rpc.md
import (
	"context"
	"go-grpc-api-lab/pkg/opentelemetry-go/plugin/grpctrace"
	"log"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/distributedcontext"
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/exporter/trace/stdout"
	"go.opentelemetry.io/otel/global"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func init() {
	initOtelTracer()
}

func initOtelTracer() {
	exporter, err := stdout.NewExporter(stdout.Options{PrettyPrint: true})
	if err != nil {
		log.Fatal(err)
	}
	tp, err := sdktrace.NewProvider(
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithSyncer(exporter),
	)
	if err != nil {
		log.Fatal(err)
	}
	global.SetTraceProvider(tp)
}

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

	tr := global.TraceProvider().Tracer("example/grpc")
	err := tr.WithSpan(ctx, "hello-api-op",
		func(ctx context.Context) error {
			grpctrace.Inject(ctx, &metadataCopy)
			ctx = metadata.NewOutgoingContext(ctx, metadataCopy)

			err := invoker(ctx, method, req, reply, cc, opts...)
			setTraceStatus(ctx, err)
			return err
		})
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
