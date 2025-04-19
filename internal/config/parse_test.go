package config_test

import (
	"testing"

	"github.com/RaphSku/synmake/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestParsing(t *testing.T) {
	t.Parallel()

	// --- Setup Logger
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	// --- Generating example configuration
	fileBuffer := NewFileBuffer("config.yaml")
	err = config.GenerateExampleYamlConfig(logger, fileBuffer)
	assert.NoError(t, err)

	// Checking whether parsing works and fields are set correctly
	configManager, err := config.NewConfigManager(logger, fileBuffer)
	assert.NoError(t, err)
	err = configManager.Parse()
	assert.NoError(t, err)
	assert.Equal(t, false, isEmptyStruct(configManager.Config))
	assert.Equal(t, true, configManager.Config.Templates.HelpTargetTemplate.Enabled)
	assert.Equal(t, "example", configManager.Config.Templates.VersionCommandTemplate.Library)
}
