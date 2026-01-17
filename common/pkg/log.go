package pkg

import (
	"common/pkg/constant"
	"common/pkg/cutil/base/logger"
	"context"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _ log.Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	log    *zap.Logger
	msgKey string
	Sync   func() error
}

// TraceIDValuer 返回一个能从 context 中获取 trace_id 的 Valuer
func TraceIDValuer() log.Valuer {
	return func(ctx context.Context) interface{} {
		if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
			return span.TraceID().String()
		}
		return ""
	}
}

// SpanIDValuer 返回一个能从 context 中获取 span_id 的 Valuer
func SpanIDValuer() log.Valuer {
	return func(ctx context.Context) interface{} {
		if span := trace.SpanContextFromContext(ctx); span.HasSpanID() {
			return span.SpanID().String()
		}
		return ""
	}
}

func NewLogger(name string, version string, mode string, level string, file string) log.Logger {
	// 创建 ZapLogger
	zapLogger := Logger(mode, level, file)

	// 添加全局字段（With 包装）
	l := log.With(zapLogger,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.name", name,
		"service.version", version,
		"trace_id", TraceIDValuer(), // 自动注入 trace_id
		"span_id", SpanIDValuer(), // 自动注入 span_id
	)
	log.SetLogger(l)
	return log.GetLogger()
}

// Logger 配置zap日志,将zap日志库引入
func Logger(mode string, level string, file string) log.Logger {
	// 配置zap日志库的编码器
	encoder := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stack",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	levelMap := map[string]zapcore.Level{
		"debug":  zap.DebugLevel,
		"info":   zap.InfoLevel,
		"warn":   zap.WarnLevel,
		"error":  zap.ErrorLevel,
		"dpanic": zap.DPanicLevel,
		"panic":  zap.PanicLevel,
		"fatal":  zap.FatalLevel,
	}
	l, ok := levelMap[level]
	if !ok {
		l = zap.InfoLevel
	}

	opts := []zap.Option{
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.AddCaller(),
		zap.AddCallerSkip(2), // 调整为 1，因为 Kratos log 层已包装
	}

	if mode == constant.Dev {
		opts = append(opts, zap.Development()) // dev 模式添加开发友好字段
	}

	zapLogger := NewZapLogger(mode, file, encoder, zap.NewAtomicLevelAt(l), opts...)
	// 将工具库的 logger 转换为 zap 的 Logger
	adapter := NewZapAdapter(zapLogger.log)
	logger.SetLogger(adapter)

	return zapLogger
}

// NewZapLogger 创建 ZapLogger 实例
func NewZapLogger(mode string, file string, encoder zapcore.EncoderConfig, level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {
	writers := []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}
	var core zapcore.Core
	if file != "" {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   file,
			MaxSize:    1024, // MB
			MaxBackups: 5,
			MaxAge:     30, // days
			Compress:   true,
		}
		writers = append(writers, zapcore.AddSync(lumberJackLogger))
	}

	if mode == constant.Dev {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoder),
			zapcore.NewMultiWriteSyncer(writers...), // 总是输出到 stdout，可选文件
			level,
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoder),
			zapcore.NewMultiWriteSyncer(writers...), // 总是输出到 stdout，可选文件
			level,
		)
	}

	zapLogger := zap.New(core, opts...)
	return &ZapLogger{
		log:    zapLogger,
		msgKey: "msg", // 统一 msgKey
		Sync:   zapLogger.Sync,
	}
}

// Log 实现log接口
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if zapcore.Level(level) < zapcore.DPanicLevel && !l.log.Core().Enabled(zapcore.Level(level)) {
		return nil
	}
	var (
		msg    = ""
		keylen = len(keyvals)
	)
	if keylen == 0 || keylen%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	data := make([]zap.Field, 0, (keylen/2)+1)
	for i := 0; i < keylen; i += 2 {
		if keyvals[i].(string) == l.msgKey {
			msg, _ = keyvals[i+1].(string)
			continue
		}
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	switch level {
	case log.LevelDebug:
		l.log.Debug(msg, data...)
	case log.LevelInfo:
		l.log.Info(msg, data...)
	case log.LevelWarn:
		l.log.Warn(msg, data...)
	case log.LevelError:
		l.log.Error(msg, data...)
	case log.LevelFatal:
		l.log.Fatal(msg, data...)
	}
	return nil
}

func SetupTracing(ctx context.Context, serviceName, version, endpoint string, enableOtel bool, insecure bool, sampler float64) (func(context.Context) error, error) {
	if !enableOtel {
		// 创建仅本地 traceID、不上报的 TracerProvider
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
		log.Info("Tracing disabled: using local tracer (traceID preserved, no export)")
		return func(context.Context) error { return nil }, nil
	}

	// --- enableOtel == true 的正常上报逻辑 ---
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

	// 资源属性
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(version),
		),
	)
	if err != nil {
		return nil, err
	}

	// TracerProvider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(sampler))),
	)
	otel.SetTracerProvider(tp)

	log.Info("Tracing setup complete (OTEL enabled)")
	return tp.Shutdown, nil
}

// cutil 工具类日志实现

type ZapAdapter struct {
	zapLogger *zap.Logger
}

// NewZapAdapter 构造函数，接受一个 *zap.Logger，并返回一个适配器实例
func NewZapAdapter(zapLogger *zap.Logger) *ZapAdapter {
	return &ZapAdapter{zapLogger: zapLogger}
}

// Debug 转发日志到 zap 的 Debug 方法
func (a *ZapAdapter) Debug(format string, v ...interface{}) {
	a.zapLogger.Debug(fmt.Sprintf(format, v...))
}

// Info 转发日志到 zap 的 Info 方法
func (a *ZapAdapter) Info(format string, v ...interface{}) {
	a.zapLogger.Info(fmt.Sprintf(format, v...))
}

// Warn 转发日志到 zap 的 Warn 方法
func (a *ZapAdapter) Warn(format string, v ...interface{}) {
	a.zapLogger.Warn(fmt.Sprintf(format, v...))
}

// Error 转发日志到 zap 的 Error 方法
func (a *ZapAdapter) Error(format string, v ...interface{}) {
	a.zapLogger.Error(fmt.Sprintf(format, v...))
}
