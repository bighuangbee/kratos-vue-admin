package kitGrpc

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type TracerProviderOption struct {
	Endpoint    string
	ServiceName string
	Sampler     tracesdk.Sampler
	Attrubtes   []attribute.KeyValue
}

func NewTracerProvider(opt TracerProviderOption) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(opt.Endpoint)))
	if err != nil {
		return nil, err
	}

	opt.Attrubtes = append([]attribute.KeyValue{semconv.ServiceNameKey.String(opt.ServiceName)}, opt.Attrubtes...)
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(opt.Sampler),
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewSchemaless(
			opt.Attrubtes...,
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp, nil
}

func GetTracerProvider() trace.TracerProvider {
	return otel.GetTracerProvider()
}
