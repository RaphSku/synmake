package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestLoggerSetup(t *testing.T) {
	logger := SetupZapLogger()

	assert.Equal(t, zapcore.InfoLevel, logger.Level())
	assert.Equal(t, "", logger.Name())
}
