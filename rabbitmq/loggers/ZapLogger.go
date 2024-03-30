package loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func init() {
	err := initializeLogger()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

var Zap *zap.SugaredLogger

func initializeLogger() error {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("Jan 02 15:04:05.000000000")
	config.EncoderConfig.StacktraceKey = ""

	logger, err := config.Build()
	if err != nil {
		return err
	}

	var errDefers error
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			errDefers = err
		}
	}(logger)

	Zap = logger.Sugar()

	return errDefers
}
