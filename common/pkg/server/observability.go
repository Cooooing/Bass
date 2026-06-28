package server

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"net/http"
	"strconv"
	"time"

	kratoserrors "github.com/go-kratos/kratos/v3/errors"
	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/transport"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var (
	serverRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "kratos_server_requests_total",
			Help: "Total number of Kratos server requests.",
		},
		[]string{"service", "kind", "operation", "status"},
	)
	serverRequestDurationSeconds = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "kratos_server_request_duration_seconds",
			Help:    "Kratos server request duration in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"service", "kind", "operation", "status"},
	)
)

func init() {
	prometheus.MustRegister(serverRequestTotal, serverRequestDurationSeconds)
}

func MetricsMiddleware(service string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			start := time.Now()
			reply, err := handler(ctx, req)
			kind, operation := transportInfo(ctx)
			status := statusLabel(err)

			serverRequestTotal.WithLabelValues(service, kind, operation, status).Inc()
			serverRequestDurationSeconds.WithLabelValues(service, kind, operation, status).Observe(time.Since(start).Seconds())

			return reply, err
		}
	}
}

func TracingMiddleware(service string) middleware.Middleware {
	tracer := otel.Tracer(service)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			kind, operation := transportInfo(ctx)
			spanName := operation
			if spanName == "unknown" {
				spanName = kind
			}
			ctx, span := tracer.Start(
				ctx,
				spanName,
				oteltrace.WithSpanKind(oteltrace.SpanKindServer),
				oteltrace.WithAttributes(
					attribute.String("rpc.system", kind),
					attribute.String("rpc.service", service),
					attribute.String("rpc.method", operation),
				),
			)
			defer span.End()

			reply, err := handler(ctx, req)
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			return reply, err
		}
	}
}

func transportInfo(ctx context.Context) (string, string) {
	kind := "unknown"
	operation := "unknown"
	if tr, ok := transport.FromServerContext(ctx); ok {
		if value := tr.Kind().String(); value != "" {
			kind = value
		}
		if value := tr.Operation(); value != "" {
			operation = value
		}
	}
	return kind, operation
}

func statusLabel(err error) string {
	if err == nil {
		return strconv.Itoa(http.StatusOK)
	}
	if code, ok := apperror.BusinessCode(err); ok {
		return strconv.Itoa(apperror.StatusCode(code))
	}
	if se := kratoserrors.FromError(err); se != nil && se.Code > 0 {
		return strconv.Itoa(int(se.Code))
	}
	return strconv.Itoa(apperror.StatusCode(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INTERNAL))
}
