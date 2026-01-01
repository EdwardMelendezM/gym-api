package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init(env string) {
	cfg := zap.NewProductionConfig()

	if env == "development" {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.Encoding = "json"
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	Log = logger
}

func Sync() {
	_ = Log.Sync()
}
