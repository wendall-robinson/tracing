package tracer

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/wendall-robinson/tracing/pkg/loggers"
)

// Tracer is a wrapper around the OpenTelemetry Tracer
type Tracer struct {
	tracer trace.Tracer
	span   trace.Span
	logger loggers.Logger
}

// Args is a map of argument names and values
type Args map[string]string

// New will create a new tracer
func New(tracerName string, logger loggers.Logger) *Tracer {
	tracer := otel.Tracer(tracerName)

	return &Tracer{
		tracer: tracer,
		logger: logger,
	}
}

// TraceWithAttributes will create, start and add args to a new span
// This is a convenience function that combines New, Start and AddParams
func TraceWithAttributes(ctx context.Context, logger loggers.Logger, tracerName, operation string, args Args) *Tracer {
	tracer := New(tracerName, logger)
	tracer.StartSpan(ctx, operation)

	if args != nil {
		tracer.AddParams(args)
	}

	return tracer
}

// StartSpan starts a new span
func (t *Tracer) StartSpan(ctx context.Context, operation string) context.Context {
	ctx, t.span = t.tracer.Start(ctx, operation)
	return ctx
}

// AddParams adds argument parameters to the span
func (t *Tracer) AddParams(params Args) {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		t.logger.Errorf("Failed to marshal params to JSON: %v", err)
		return
	}

	t.span.SetAttributes(attribute.String("args", string(jsonParams)))
}

// EndSpan ends the span
func (t *Tracer) EndSpan() {
	t.span.End()
}

// BuildSpanParams builds a map of parameter names and values from the provided arguments
func BuildSpanParams(args ...interface{}) Args {
	params := make(map[string]string)

	for i := 0; i < len(args); i += 2 {
		key, keyOk := args[i].(string)
		if !keyOk {
			continue
		}

		switch value := args[i+1].(type) {
		case string:
			params[key] = value
		case int:
			params[key] = strconv.Itoa(value)
		case float64:
			params[key] = strconv.FormatFloat(value, 'f', -1, 64)
		case bool:
			params[key] = strconv.FormatBool(value)
		case int32:
			params[key] = strconv.FormatInt(int64(value), 10)
		case int64:
			params[key] = strconv.FormatInt(value, 10)
		case float32:
			params[key] = strconv.FormatFloat(float64(value), 'f', -1, 32)
		case uint:
			params[key] = strconv.FormatUint(uint64(value), 10)
		case uint32:
			params[key] = strconv.FormatUint(uint64(value), 10)
		case uint64:
			params[key] = strconv.FormatUint(value, 10)
		default:
			params[key] = fmt.Sprintf("%v", value)
		}
	}

	return params
}
