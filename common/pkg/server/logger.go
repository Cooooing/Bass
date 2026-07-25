package server

import (
	"context"
	"io"
	"log/slog"
	"os"

	"common/pkg/constant"
	"common/proto/gen/common"

	kratoslog "github.com/go-kratos/kratos/v3/log"
	"go.opentelemetry.io/otel/trace"
)

func NewLogger(
	server *common.Server,
	conf *common.Bootstrap_Log,
) *slog.Logger {
	mode := ""
	name := "unknown"
	version := "unknown"
	if server != nil {
		mode = server.GetMode()
		if server.GetName() != "" {
			name = server.GetName()
		}
		if server.GetVersion() != "" {
			version = server.GetVersion()
		}
	}

	level := "info"
	if conf != nil && conf.GetLevel() != "" {
		level = conf.GetLevel()
	}

	format := kratoslog.FormatJSON
	if mode == "dev" {
		format = kratoslog.FormatText
	}

	handler := kratoslog.NewHandler(
		kratoslog.WithWriter(io.Writer(os.Stdout)),
		kratoslog.WithFormat(format),
		kratoslog.WithLevel(kratoslog.ParseLevel(level)),
		kratoslog.WithAddSource(false),
		kratoslog.WithFilter(kratoslog.FilterKey(
			constant.LogFieldPassword,
			constant.LogFieldToken,
			constant.LogFieldSecret,
			constant.LogFieldAuthorization,
			constant.LogFieldCookie,
		)),
		kratoslog.WithExtractor(func(ctx context.Context) []slog.Attr {
			span := trace.SpanContextFromContext(ctx)
			if !span.IsValid() {
				return nil
			}
			attrs := make([]slog.Attr, 0, 2)
			if span.HasTraceID() {
				attrs = append(attrs, slog.String(constant.LogFieldTraceID, span.TraceID().String()))
			}
			if span.HasSpanID() {
				attrs = append(attrs, slog.String(constant.LogFieldSpanID, span.SpanID().String()))
			}
			return attrs
		}),
	)
	handler = sourceHandler{
		next: handler,
	}
	logger := slog.New(handler).With(
		slog.String(constant.LogFieldServiceName, name),
		slog.String(constant.LogFieldServiceVersion, version),
	)
	slog.SetDefault(logger)
	kratoslog.SetDefault(logger)
	return logger
}

type sourceHandler struct {
	next slog.Handler
}

func (h sourceHandler) Enabled(
	ctx context.Context,
	level slog.Level,
) bool {
	return h.next.Enabled(ctx, level)
}

func (h sourceHandler) Handle(
	ctx context.Context,
	record slog.Record,
) error {
	withSource := false
	attrs := make([]slog.Attr, 0, record.NumAttrs())
	record.Attrs(func(attr slog.Attr) bool {
		if attr.Key == constant.LogFieldUnexpected {
			if attr.Value.Kind() == slog.KindBool && attr.Value.Bool() {
				withSource = true
			}
			return true
		}
		attrs = append(attrs, attr)
		return true
	})

	if withSource {
		if source := record.Source(); source != nil {
			attrs = append(attrs, slog.Group(constant.LogFieldSource,
				slog.String("function", source.Function),
				slog.String(constant.LogFieldFile, source.File),
				slog.Int(constant.LogFieldLine, source.Line),
			))
		}
	}

	nextRecord := slog.NewRecord(record.Time, record.Level, record.Message, record.PC)
	nextRecord.AddAttrs(attrs...)
	return h.next.Handle(ctx, nextRecord)
}

func (h sourceHandler) WithAttrs(
	attrs []slog.Attr,
) slog.Handler {
	return sourceHandler{
		next: h.next.WithAttrs(attrs),
	}
}

func (h sourceHandler) WithGroup(
	name string,
) slog.Handler {
	return sourceHandler{
		next: h.next.WithGroup(name),
	}
}
