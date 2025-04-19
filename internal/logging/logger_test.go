package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestLoggerWithDebugModeSetup(t *testing.T) {
	t.Parallel()

	logger := SetupZapLogger(true)

	assert.Equal(t, zapcore.InfoLevel, logger.Level())
	assert.Equal(t, "synmake", logger.Name())
}

func TestLoggerWithoutDebugModeSetup(t *testing.T) {
	t.Parallel()

	logger := SetupZapLogger(false)

	assert.Equal(t, zapcore.Level(6), logger.Level())
	assert.Equal(t, "", logger.Name())
}
