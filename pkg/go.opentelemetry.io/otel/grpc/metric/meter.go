package metric

import (
	"context"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/api/metric"
	"google.golang.org/grpc"
)

var (
	appKey       = key.New("app")
	operationKey = key.New("operation")
)

// UnaryClientInterceptor records custom app metrics
func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
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
