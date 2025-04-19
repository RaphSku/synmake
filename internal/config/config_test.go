package config_test

import (
	"testing"

	"github.com/RaphSku/synmake/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestConfigManagerInitNoError(t *testing.T) {
	t.Parallel()

	// --- Logger Setup
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	// --- Checking whether the config manager initialization throws no error
	fileBuffer := NewFileBuffer("config.yaml")
	_, err = config.NewConfigManager(logger, fileBuffer)
	assert.NoError(t, err)
}
