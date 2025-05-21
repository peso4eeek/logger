package logger

import (
	"context"

	"github.com/peso4eeek/logger/core"

	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

// Обертка для grpc
func GRPCInterceptor(log core.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, level grpclog.Level, msg string, fields ...any) {
		args := make([]interface{}, 0, len(fields)*2)
		for i := 0; i < len(fields); i += 2 {
			if i+1 < len(fields) {
				key, ok := fields[i].(string)
				if !ok {
					continue
				}
				args = append(args, key, fields[i+1])
			}
		}

		switch level {
		case grpclog.LevelDebug:
			log.Debug(msg, args...)
		case grpclog.LevelInfo:
			log.Info(msg, args...)
		case grpclog.LevelWarn:
			log.Warn(msg, args...)
		case grpclog.LevelError:
			log.Error(msg, args...)
		}
	})
}
