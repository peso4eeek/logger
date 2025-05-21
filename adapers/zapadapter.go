package adapers

import (
	"github.com/peso4eeek/logger/core"

	"go.uber.org/zap"
)

type ZapAdapter struct {
	*zap.Logger
}

// Обертка для zap
func NewZapAdapter(log *zap.Logger) core.Logger {
	return &ZapAdapter{Logger: log}
}

func (l *ZapAdapter) With(args ...interface{}) core.Logger {
	fields := make([]zap.Field, 0, len(args)/2)
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			fields = append(fields, zap.Any(args[i].(string), args[i+1]))
		}
	}
	return &ZapAdapter{Logger: l.Logger.With(fields...)}
}

func (l *ZapAdapter) Info(msg string, args ...interface{}) {
	l.Logger.Sugar().Infof(msg, args...)
}

func (l *ZapAdapter) Error(msg string, args ...interface{}) {
	l.Logger.Sugar().Errorf(msg, args...)
}

func (l *ZapAdapter) Debug(msg string, args ...interface{}) {
	l.Logger.Sugar().Debugf(msg, args...)
}

func (l *ZapAdapter) Warn(msg string, args ...interface{}) {
	l.Logger.Sugar().Warnf(msg, args...)
}

func (l *ZapAdapter) Fatal(msg string, args ...interface{}) {
	l.Logger.Sugar().Fatal(msg, args)
}
