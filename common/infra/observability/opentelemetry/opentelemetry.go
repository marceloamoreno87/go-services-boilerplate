package opentelemetry

import (
	"context"
	"os"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type OpenTelemetry struct{}

func (o OpenTelemetry) Start() (optl trace.Tracer, err error) {
	ctx := context.Background()

	otlpEndpoint := os.Getenv("OTLP_ENDPOINT")
	nameProject := os.Getenv("NAME_PROJECT")

	exp, err := o.newOTLPExporter(ctx, otlpEndpoint)
	if err != nil {
		return
	}

	traceProvider, err := o.newTraceProvider(exp, nameProject)
	if err != nil {
		return
	}

	return traceProvider.Tracer(nameProject), nil
}

func (o OpenTelemetry) newOTLPExporter(ctx context.Context, otlpEndpoint string) (sdkTrace.SpanExporter, error) {
	insecureOpt := otlptracehttp.WithInsecure()
	endpointOpt := otlptracehttp.WithEndpoint(otlpEndpoint)

	return otlptracehttp.New(ctx, insecureOpt, endpointOpt)
}

func (o OpenTelemetry) newTraceProvider(exp sdkTrace.SpanExporter, projectName string) (*sdkTrace.TracerProvider, error) {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(projectName),
		),
	)
	if err != nil {
		return nil, err
	}

	return sdkTrace.NewTracerProvider(
		sdkTrace.WithBatcher(exp),
		sdkTrace.WithResource(r),
	), nil
}
