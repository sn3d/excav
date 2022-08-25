package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var suggar *zap.SugaredLogger

func init() {
	suggar = zap.NewNop().Sugar()
}

func EnableDebug() {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := cfg.Build()
	suggar = logger.Sugar()
}

func Error(msg string, err error) {
	suggar.Errorf("%s: %v", msg, err)
}

func Errorw(msg string, keysAndVals ...interface{}) {
	suggar.Errorw(msg, keysAndVals...)
}

func Debug(msg string, keysAndVals ...interface{}) {
	suggar.Debugw(msg, keysAndVals...)
}
