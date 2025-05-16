package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupZapLogger(debugMode bool) *zap.Logger {
	logLevel := zapcore.InfoLevel

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		logLevel,
	)

	loggerOptions := []zap.Option{
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Fields(zap.String("service", "synmake")),
	}

	if debugMode {
		logger := zap.New(core, loggerOptions...)
		logger = logger.Named("synmake")
		return logger
	}

	return zap.NewNop()
}
