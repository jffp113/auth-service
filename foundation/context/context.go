package context

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type ctxKey int

const key ctxKey = 1

type Values struct {
	TraceId string
	Tracer  trace.Tracer
}

func Wrap(ctx context.Context, traceId string, tracer trace.Tracer) context.Context {
	return context.WithValue(ctx, key, &Values{
		TraceId: traceId,
		Tracer:  tracer,
	})
}

func GetValues(ctx context.Context) *Values {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return &Values{
			TraceId: "00000000-0000-0000-0000-000000000000",
			Tracer:  trace.NewNoopTracerProvider().Tracer(""),
		}
	}
	return v
}

func GetTraceID(ctx context.Context) string {
	values := GetValues(ctx)
	return values.TraceId
}

// AddSpan adds a OpenTelemetry span to the trace and context.
func AddSpan(ctx context.Context, spanName string, keyValues ...attribute.KeyValue) (context.Context, trace.Span) {
	v, ok := ctx.Value(key).(*Values)
	if !ok || v.Tracer == nil {
		return ctx, trace.SpanFromContext(ctx)
	}

	ctx, span := v.Tracer.Start(ctx, spanName)
	for _, kv := range keyValues {
		span.SetAttributes(kv)
	}

	return ctx, span
}
