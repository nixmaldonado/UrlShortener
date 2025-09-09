package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.TimeKey = "timestamp" // Overrides "ts" key for clarity
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	var err error
	log, err = config.Build(zap.Fields(zap.String("commit", commitHash)))
	if err != nil {
		panic("failed to initialize log: " + err.Error())
	}

	InitConfig()

	// Run tests
	os.Exit(m.Run())
}
