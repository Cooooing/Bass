package pkg

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/log"
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
		EncodeCaller:   zapcore.FullCallerEncoder,
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

	return NewZapLogger(
		mode,
		file,
		encoder,
		zap.NewAtomicLevelAt(l),
		zap.AddStacktrace(zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.Development(),
	)
}

// 日志自动切割，采用 lumberjack 实现的
func getLogWriter(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 指定日志存储位置
		MaxSize:    1024, // 日志的最大大小（M）
		MaxBackups: 5,    // 日志的最大保存数量
		MaxAge:     30,   // 日志文件存储最大天数
		Compress:   true, // 是否执行压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

// NewZapLogger return a zap logger.
func NewZapLogger(mode string, file string, encoder zapcore.EncoderConfig, level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {
	writeSyncer := getLogWriter(file) // 日志切割
	var core zapcore.Core
	if mode == "dev" {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoder),                                                    // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writeSyncer)), // 打印到控制台
			level, // 日志级别
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoder), // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writeSyncer)), // 打印到控制台和文件
			level, // 日志级别
		)
	}
	zapLogger := zap.New(core, opts...)
	return &ZapLogger{
		log:    zapLogger,
		msgKey: log.DefaultMessageKey,
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
