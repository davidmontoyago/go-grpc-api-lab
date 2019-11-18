package tracing

// Open Telemetry gRPC Tracer integration
import (
	"context"
	"log"

	"go.opentelemetry.io/otel/exporter/trace/stdout"
	"go.opentelemetry.io/otel/global"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func init() {
	initTracer()
}

func initTracer() {
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

// Server side tracing interceptor
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	metadata, _ := metadata.FromIncomingContext(ctx)
	log.Printf("tracing request with metadata: %v", metadata)
	response, err := handler(ctx, req)
	// TODO trace here
	return response, err
}

// Client side tracing interceptor
func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	// TODO trace here
	return err
}
