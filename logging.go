package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var commitHash string
var log *zap.Logger

// InitLogging initializes the Zap log with commit hash
func InitLogging() {
	logConf := zap.NewProductionConfig()
	logConf.OutputPaths = []string{"stdout", "logs/app.log"}
	logConf.EncoderConfig.TimeKey = "timestamp" // Overrides "ts" key for clarity
	logConf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logConf.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	var err error
	log, err = logConf.Build(zap.Fields(zap.String("commit", commitHash)))
	if err != nil {
		panic("failed to initialize log: " + err.Error())
	}
}
