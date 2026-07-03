package client

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"log/slog"
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
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var (
	serverRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "bass_server_requests_total", Help: "Total number of server requests."},
		[]string{"service", "transport", "operation", "status_code"},
	)
	serverRequestDurationSeconds = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{Name: "bass_server_request_duration_seconds", Help: "Server request duration in seconds.", Buckets: prometheus.DefBuckets},
		[]string{"service", "transport", "operation", "status_code"},
	)
	clientRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "bass_client_requests_total", Help: "Total number of client requests."},
		[]string{"service", "target", "transport", "operation", "status_code"},
	)
	clientRequestDurationSeconds = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{Name: "bass_client_request_duration_seconds", Help: "Client request duration in seconds.", Buckets: prometheus.DefBuckets},
		[]string{"service", "target", "transport", "operation", "status_code"},
	)
	MessageRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "bass_message_requests_total", Help: "Total number of message publish and consume operations."},
		[]string{"service", "direction", "subject", "status"},
	)
	MessageRequestDurationSeconds = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{Name: "bass_message_request_duration_seconds", Help: "Message publish and consume duration in seconds.", Buckets: prometheus.DefBuckets},
		[]string{"service", "direction", "subject", "status"},
	)
	DeadLetterAlertsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "bass_dead_letter_alerts_total", Help: "Total number of deduplicated dead letter alerts."},
		[]string{"service", "source", "event_type", "subject"},
	)
)

func init() {
	prometheus.MustRegister(serverRequestsTotal, serverRequestDurationSeconds, clientRequestsTotal, clientRequestDurationSeconds, MessageRequestsTotal, MessageRequestDurationSeconds, DeadLetterAlertsTotal)
}

type Observer struct {
	service      string
	logger       *slog.Logger
	serverTracer oteltrace.Tracer
	clientTracer oteltrace.Tracer
}

func NewObservability(logger *slog.Logger, server *common.Server) *Observer {
	service := "unknown"
	if server != nil {
		if server.GetName() != "" {
			service = server.GetName()
		}
	}
	return &Observer{
		service:      service,
		logger:       logger,
		serverTracer: otel.Tracer(service + ".server"),
		clientTracer: otel.Tracer(service + ".client"),
	}
}

func (o *Observer) Service() string {
	if o == nil || o.service == "" {
		return "unknown"
	}
	return o.service
}

func SetupTracing(ctx context.Context, serviceName string, version string, traceConf *common.Trace) (func(context.Context) error, error) {
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName(serviceName), semconv.ServiceVersion(version)))
	if err != nil {
		return nil, err
	}

	if traceConf == nil || !traceConf.GetEnableOtel() {
		tp := sdktrace.NewTracerProvider(sdktrace.WithSampler(sdktrace.AlwaysSample()), sdktrace.WithResource(res))
		otel.SetTracerProvider(tp)
		slog.Info("tracing disabled: using local tracer provider")
		return func(context.Context) error { return nil }, nil
	}

	sampler := traceConf.GetSampler()
	if sampler < 0 || sampler > 1 {
		slog.Warn("invalid tracing sampler, fallback to 1", "sampler", sampler)
		sampler = 1
	}

	opts := []otlptracegrpc.Option{otlptracegrpc.WithEndpoint(traceConf.GetEndpoint())}
	if traceConf.GetInsecure() {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	client := otlptracegrpc.NewClient(opts...)
	exp, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(sampler))),
	)
	otel.SetTracerProvider(tp)
	slog.Info("tracing setup complete")
	return tp.Shutdown, nil
}

func (o *Observer) ServerMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			start := time.Now()
			transportKind := "unknown"
			operation := "unknown"
			if tr, ok := transport.FromServerContext(ctx); ok {
				if header := tr.RequestHeader(); header != nil {
					carrier := propagation.MapCarrier{}
					for _, key := range header.Keys() {
						if value := header.Get(key); value != "" {
							carrier.Set(key, value)
						}
					}
					ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
				}
				if value := tr.Kind().String(); value != "" {
					transportKind = value
				}
				if value := tr.Operation(); value != "" {
					operation = value
				}
			}

			spanName := operation
			if spanName == "unknown" {
				spanName = transportKind
			}
			ctx, span := o.serverTracer.Start(ctx, spanName, oteltrace.WithSpanKind(oteltrace.SpanKindServer), oteltrace.WithAttributes(
				attribute.String("rpc.system", transportKind),
				attribute.String("rpc.service", o.service),
				attribute.String("rpc.method", operation),
			))
			defer span.End()

			reply, err := handler(ctx, req)
			statusCode := http.StatusOK
			if err != nil {
				if code, ok := apperror.BusinessCode(err); ok {
					statusCode = apperror.StatusCode(code)
				} else if se := kratoserrors.FromError(err); se != nil && se.Code > 0 {
					statusCode = int(se.Code)
				} else {
					statusCode = apperror.StatusCode(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INTERNAL)
				}
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			statusLabel := strconv.Itoa(statusCode)
			latency := time.Since(start)

			serverRequestsTotal.WithLabelValues(o.service, transportKind, operation, statusLabel).Inc()
			serverRequestDurationSeconds.WithLabelValues(o.service, transportKind, operation, statusLabel).Observe(latency.Seconds())

			attrs := []slog.Attr{
				slog.String(constant.LogFieldKind, constant.LogKindServer),
				slog.String(constant.LogFieldTransport, transportKind),
				slog.String(constant.LogFieldOperation, operation),
				slog.Int(constant.LogFieldStatusCode, statusCode),
				slog.Int64(constant.LogFieldLatencyMS, latency.Milliseconds()),
			}
			if err != nil {
				attrs = append(attrs, slog.Any(constant.LogFieldErr, err))
			}
			level := slog.LevelInfo
			switch {
			case statusCode >= http.StatusInternalServerError:
				level = slog.LevelError
			case statusCode >= http.StatusBadRequest:
				level = slog.LevelWarn
			}
			o.logger.LogAttrs(ctx, level, "server request", attrs...)
			return reply, err
		}
	}
}

func (o *Observer) ClientMiddleware(target string) middleware.Middleware {
	if target == "" {
		target = "unknown"
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			start := time.Now()
			transportKind := "unknown"
			operation := "unknown"
			if tr, ok := transport.FromClientContext(ctx); ok {
				if value := tr.Kind().String(); value != "" {
					transportKind = value
				}
				if value := tr.Operation(); value != "" {
					operation = value
				}
			}

			spanName := operation
			if spanName == "unknown" {
				spanName = transportKind
			}
			ctx, span := o.clientTracer.Start(ctx, spanName, oteltrace.WithSpanKind(oteltrace.SpanKindClient), oteltrace.WithAttributes(
				attribute.String("rpc.system", transportKind),
				attribute.String("rpc.service", target),
				attribute.String("rpc.method", operation),
			))
			defer span.End()

			if tr, ok := transport.FromClientContext(ctx); ok {
				if header := tr.RequestHeader(); header != nil {
					carrier := propagation.MapCarrier{}
					otel.GetTextMapPropagator().Inject(ctx, carrier)
					for _, key := range carrier.Keys() {
						header.Set(key, carrier.Get(key))
					}
				}
			}

			reply, err := handler(ctx, req)
			statusCode := http.StatusOK
			if err != nil {
				if se := kratoserrors.FromError(err); se != nil && se.Code > 0 {
					statusCode = int(se.Code)
				} else {
					statusCode = http.StatusInternalServerError
				}
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			statusLabel := strconv.Itoa(statusCode)
			latency := time.Since(start)

			clientRequestsTotal.WithLabelValues(o.service, target, transportKind, operation, statusLabel).Inc()
			clientRequestDurationSeconds.WithLabelValues(o.service, target, transportKind, operation, statusLabel).Observe(latency.Seconds())

			attrs := []slog.Attr{
				slog.String(constant.LogFieldKind, constant.LogKindClient),
				slog.String(constant.LogFieldTarget, target),
				slog.String(constant.LogFieldTransport, transportKind),
				slog.String(constant.LogFieldOperation, operation),
				slog.Int(constant.LogFieldStatusCode, statusCode),
				slog.Int64(constant.LogFieldLatencyMS, latency.Milliseconds()),
			}
			if err != nil {
				attrs = append(attrs, slog.Any(constant.LogFieldErr, err))
			}
			level := slog.LevelInfo
			switch {
			case statusCode >= http.StatusInternalServerError:
				level = slog.LevelError
			case statusCode >= http.StatusBadRequest:
				level = slog.LevelWarn
			}
			o.logger.LogAttrs(ctx, level, "client request", attrs...)
			return reply, err
		}
	}
}
