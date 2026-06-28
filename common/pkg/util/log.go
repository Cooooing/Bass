package util

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	kratoslog "github.com/go-kratos/kratos/v3/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogHelper struct {
	logger *slog.Logger
	ctx    context.Context
}

func NewLogger(name string, version string, mode string, level string, file string) *slog.Logger {
	logger := Logger(mode, level, file).With(
		slog.String("service.name", name),
		slog.String("service.version", version),
	)
	slog.SetDefault(logger)
	kratoslog.SetDefault(logger)
	return logger
}

func Logger(mode string, level string, file string) *slog.Logger {
	var levelVar slog.LevelVar
	levelVar.Set(parseSlogLevel(level))

	writers := []io.Writer{os.Stdout}
	if file != "" {
		writers = append(writers, &lumberjack.Logger{
			Filename:   file,
			MaxSize:    1024,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   true,
		})
	}

	opts := &slog.HandlerOptions{
		AddSource: mode == "dev",
		Level:     &levelVar,
	}
	writer := io.MultiWriter(writers...)
	var handler slog.Handler
	if mode == "dev" {
		handler = slog.NewTextHandler(writer, opts)
	} else {
		handler = slog.NewJSONHandler(writer, opts)
	}
	return slog.New(traceHandler{Handler: handler})
}

func NewLogHelper(logger *slog.Logger) *LogHelper {
	if logger == nil {
		logger = slog.Default()
	}
	return &LogHelper{logger: logger, ctx: context.Background()}
}

func (h *LogHelper) WithContext(ctx context.Context) *LogHelper {
	if ctx == nil {
		ctx = context.Background()
	}
	return &LogHelper{logger: h.logger, ctx: ctx}
}

func (h *LogHelper) Debug(args ...any) {
	h.logger.DebugContext(h.ctx, fmt.Sprint(args...))
}

func (h *LogHelper) Info(args ...any) {
	h.logger.InfoContext(h.ctx, fmt.Sprint(args...))
}

func (h *LogHelper) Warn(args ...any) {
	h.logger.WarnContext(h.ctx, fmt.Sprint(args...))
}

func (h *LogHelper) Error(args ...any) {
	h.logger.ErrorContext(h.ctx, fmt.Sprint(args...))
}

func (h *LogHelper) Fatal(args ...any) {
	h.logger.ErrorContext(h.ctx, fmt.Sprint(args...))
	os.Exit(1)
}

func (h *LogHelper) Debugf(format string, args ...any) {
	h.logger.DebugContext(h.ctx, fmt.Sprintf(format, args...))
}

func (h *LogHelper) Infof(format string, args ...any) {
	h.logger.InfoContext(h.ctx, fmt.Sprintf(format, args...))
}

func (h *LogHelper) Warnf(format string, args ...any) {
	h.logger.WarnContext(h.ctx, fmt.Sprintf(format, args...))
}

func (h *LogHelper) Errorf(format string, args ...any) {
	h.logger.ErrorContext(h.ctx, fmt.Sprintf(format, args...))
}

func (h *LogHelper) Debugw(keyvals ...any) {
	h.logKeyvals(slog.LevelDebug, keyvals...)
}

func (h *LogHelper) Warnw(keyvals ...any) {
	h.logKeyvals(slog.LevelWarn, keyvals...)
}

func (h *LogHelper) logKeyvals(level slog.Level, keyvals ...any) {
	msg := ""
	attrs := make([]any, 0, len(keyvals))
	for i := 0; i < len(keyvals); i += 2 {
		if i+1 >= len(keyvals) {
			attrs = append(attrs, "bad_keyvals", fmt.Sprint(keyvals[i]))
			break
		}
		key := fmt.Sprint(keyvals[i])
		value := keyvals[i+1]
		if key == "msg" {
			msg = fmt.Sprint(value)
			continue
		}
		attrs = append(attrs, key, value)
	}
	h.logger.Log(h.ctx, level, msg, attrs...)
}

type traceHandler struct {
	slog.Handler
}

func (h traceHandler) Handle(ctx context.Context, record slog.Record) error {
	if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
		record.AddAttrs(slog.String("trace_id", span.TraceID().String()))
	}
	if span := trace.SpanContextFromContext(ctx); span.HasSpanID() {
		record.AddAttrs(slog.String("span_id", span.SpanID().String()))
	}
	return h.Handler.Handle(ctx, record)
}

func (h traceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return traceHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h traceHandler) WithGroup(name string) slog.Handler {
	return traceHandler{Handler: h.Handler.WithGroup(name)}
}

func parseSlogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func SetupTracing(ctx context.Context, serviceName, version, endpoint string, enableOtel bool, insecure bool, sampler float64) (func(context.Context) error, error) {
	if !enableOtel {
		res, err := resource.New(ctx,
			resource.WithAttributes(
				semconv.ServiceName(serviceName),
				semconv.ServiceVersion(version),
			),
		)
		if err != nil {
			return nil, err
		}

		tp := sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(res),
		)
		otel.SetTracerProvider(tp)
		slog.Info("Tracing disabled: using local tracer (traceID preserved, no export)")
		return func(context.Context) error { return nil }, nil
	}

	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(endpoint),
	}
	if insecure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	client := otlptracegrpc.NewClient(opts...)
	exp, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(version),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(sampler))),
	)
	otel.SetTracerProvider(tp)

	slog.Info("Tracing setup complete (OTEL enabled)")
	return tp.Shutdown, nil
}
