package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

func LoadZapLogger(output_path string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config.OutputPaths = []string{output_path}
	config.ErrorOutputPaths = []string{"zap.log"}

	return config.Build()
}

type ZapGormLogger struct {
	zapLogger *zap.Logger
	logLevel  logger.LogLevel
}

func NewZapGormLogger(zapLogger *zap.Logger) *ZapGormLogger {
	return &ZapGormLogger{
		zapLogger: zapLogger,
		logLevel:  logger.Info, // 默认日志级别
	}
}

// LogMode 实现 logger.Interface 的 LogMode 方法
func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.logLevel = level
	return &newLogger
}

// Info 实现 logger.Interface 的 Info 方法
func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...any) {
	if l.logLevel >= logger.Info {
		l.zapLogger.Sugar().Infof(msg, data...)
	}
}

// Warn 实现 logger.Interface 的 Warn 方法
func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...any) {
	if l.logLevel >= logger.Warn {
		l.zapLogger.Sugar().Warnf(msg, data...)
	}
}

// Error 实现 logger.Interface 的 Error 方法
func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...any) {
	if l.logLevel >= logger.Error {
		l.zapLogger.Sugar().Errorf(msg, data...)
	}
}

// Trace 实现 logger.Interface 的 Trace 方法
func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.logLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := []zap.Field{
		zap.Duration("elapsed", elapsed),
		zap.String("sql", sql),
		zap.Int64("rows", rows),
	}

	if err != nil {
		l.zapLogger.Error("gorm trace error", append(fields, zap.Error(err))...)
	} else {
		l.zapLogger.Debug("gorm trace", fields...)
	}
}
