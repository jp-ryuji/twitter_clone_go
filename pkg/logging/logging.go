package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/jp-ryuji/twitter_clone_go/pkg/config"
)

// NewLogger returns a new logger.
func NewLogger(env config.AppEnv) (*zap.Logger, error) {
	if env == config.EnvDev || env == config.EnvTest {
		return zap.NewDevelopment()
	}

	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.InfoLevel)

	config := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "severity",
			NameKey:        "name",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	return config.Build()
}
