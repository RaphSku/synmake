package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestIfGenerateConfigCmdIsRunning(t *testing.T) {
	defer os.Remove("config.yaml")

	// --- Setup Logger
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	// --- Checking whether the config.yaml is being created
	generateConfigCommand := NewGenerateConfigSubCmd(logger)
	cmd := generateConfigCommand.GetGenerateConfigSubCmd()

	err = cmd.Execute()
	assert.NoError(t, err)
	assert.FileExists(t, "config.yaml")

	shortDescription := cmd.Short
	longDescription := cmd.Long
	assert.Greater(t, longDescription, shortDescription)
}
