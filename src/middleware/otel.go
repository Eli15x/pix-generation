package middleware

import (
  "context"
  "log"
  "go.opentelemetry.io/otel"
  "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
  sdktrace "go.opentelemetry.io/otel/sdk/trace"
  "go.opentelemetry.io/otel/sdk/resource"
  semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func setupTracing(ctx context.Context) func(context.Context) error {
  exp, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpointURL(
    getenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://otel-collector:4318"),
  ))
  if err != nil { log.Fatal(err) }

  res, _ := resource.New(ctx,
    resource.WithAttributes(
      semconv.ServiceNameKey.String(getenv("OTEL_SERVICE_NAME","pix-generation")),
    ),
  )

  tp := sdktrace.NewTracerProvider(
    sdktrace.WithResource(res),
    sdktrace.WithBatcher(exp), // envia em batch!
  )
  otel.SetTracerProvider(tp)
  return tp.Shutdown
}
