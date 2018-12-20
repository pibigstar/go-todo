package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 日志对象
type Logger struct {
	category string
	ctx      context.Context
	logger   *zap.SugaredLogger
}

// New 创建一个新的日志对象
func New(category string) *Logger {
	return &Logger{
		category: category,
		logger:   logger.Sugar(),
	}
}

// WithContext 加入上下文
func (log *Logger) WithContext(ctx context.Context) *Logger {
	log.ctx = ctx
	log.logger.Error()
	return log
}

type logFunc func(msg string, keysAndValues ...interface{})

// Error 错误信息打印
func (log Logger) Error(msg string, keysAndValues ...interface{}) {
	log.log(log.logger.Errorw, msg, keysAndValues...)
}

// CtxError 带上下文的错误打印
func (log Logger) CtxError(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.WithContext(ctx).Error(msg, keysAndValues...)
}

// Info 提示信息打印
func (log Logger) Info(msg string, keysAndValues ...interface{}) {
	log.log(log.logger.Infow, msg, keysAndValues...)
}

// CtxInfo 带上下文的提示信息打印
func (log Logger) CtxInfo(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.WithContext(ctx).Info(msg, keysAndValues...)
}

func (log Logger) log(fc logFunc, msg string, keysAndValues ...interface{}) {
	args := log.buildArgs(keysAndValues...)
	fc(msg, args...)
}

func (log *Logger) buildArgs(keysAndValues ...interface{}) []interface{} {
	args := []interface{}{
		"category", log.category,
	}

	extra := map[string]interface{}{}
	if len(keysAndValues) > 0 {
		for i := 0; i < len(keysAndValues); {
			if i == len(keysAndValues)-1 {
				break
			}
			key, value := keysAndValues[i], keysAndValues[i+1]
			if keyStr, ok := key.(string); ok {
				extra[keyStr] = value
			}
			i = i + 2
		}
	}
	if len(extra) > 0 {
		args = append(args, "context", extra)
	}

	log.ctx = nil

	return args
}

var logger = newLogger()

// newLogger 常见一个zap对象的log
func newLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	config.EncoderConfig = buildLogConfig(config.EncoderConfig)
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}

	log, _ := config.Build()

	return log
}

// buildLogConfig 设置配置参数
func buildLogConfig(config zapcore.EncoderConfig) zapcore.EncoderConfig {
	config.MessageKey = "message"
	config.TimeKey = "time"
	config.LevelKey = "level"
	config.CallerKey = "caller"
	config.StacktraceKey = "backtrace"
	config.EncodeLevel = zapcore.LowercaseLevelEncoder
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.LineEnding = zapcore.DefaultLineEnding
	config.EncodeDuration = zapcore.StringDurationEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder

	return config
}
