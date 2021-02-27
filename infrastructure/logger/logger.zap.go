package logger

import (
	"context"

	"github.com/hthl85/aws-vanguard-ca-etf-scraper/utils/corid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger struct
type ZapLogger struct {
	log   *zap.Logger
	sugar *zap.SugaredLogger
}

// NewZapLogger create new application logger
func NewZapLogger() (*ZapLogger, error) {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	l, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	s := l.Sugar()
	return &ZapLogger{
		log:   l,
		sugar: s,
	}, nil
}

// Close flush any buffered log entries
func (z *ZapLogger) Close() {
	z.sugar.Infow("flush log entries")
	z.log.Sync()
}

///////////////////////////////////////////////////////////
// Implement app logger interface
///////////////////////////////////////////////////////////

// Info logs info
func (z *ZapLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	if corID, ok := corid.FromContext(ctx); ok {
		keysAndValues = append(keysAndValues, "correlation-id", corID.String())
	}

	z.sugar.Infow(msg, keysAndValues...)
}

// Warn logs info
func (z *ZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	if corID, ok := corid.FromContext(ctx); ok {
		keysAndValues = append(keysAndValues, "correlation-id", corID.String())
	}

	z.sugar.Warnw(msg, keysAndValues...)
}

// Error logs info
func (z *ZapLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	if corID, ok := corid.FromContext(ctx); ok {
		keysAndValues = append(keysAndValues, "correlation-id", corID.String())
	}

	z.sugar.Errorw(msg, keysAndValues...)
}
