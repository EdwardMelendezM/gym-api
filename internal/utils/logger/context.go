package logger

import (
	"context"

	"go.uber.org/zap"
)

type ctxKey string

const loggerKey ctxKey = "logger"

func WithContext(ctx context.Context, fields ...zap.Field) context.Context {
	l := Log.With(fields...)
	return context.WithValue(ctx, loggerKey, l)
}

func FromContext(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return l
	}
	return Log
}
