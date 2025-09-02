package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var commitHash string
var log *zap.Logger

// InitLogging initializes the Zap log with commit hash
func InitLogging() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	var err error
	log, err = config.Build(zap.Fields(zap.String("commit", commitHash)))
	if err != nil {
		panic("failed to initialize log: " + err.Error())
	}
}
