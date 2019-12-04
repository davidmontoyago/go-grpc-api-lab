package config

import (
	"log"
	"time"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/exporter/trace/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	metricExporter "go.opentelemetry.io/otel/exporter/metric/stdout"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/batcher/defaultkeys"
	"go.opentelemetry.io/otel/sdk/metric/controller/push"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

// InitTracer configures otel tracer
func InitTracer() {
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

// InitMeter configures otel meter
func InitMeter() *push.Controller {
	exporter, err := metricExporter.New(metricExporter.Options{PrettyPrint: true})
	if err != nil {
		log.Fatal(err)
	}
	selector := simple.NewWithExactMeasure()
	batcher := defaultkeys.New(selector, sdkmetric.NewDefaultLabelEncoder(), false)
	pusher := push.New(batcher, exporter, 5*time.Second)
	pusher.Start()

	global.SetMeterProvider(pusher)
	return pusher
}
