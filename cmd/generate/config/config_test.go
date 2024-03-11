package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestIfGenerateConfigCmdIsRunning(t *testing.T) {
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	cmd := GetGenerateConfigCmd(logger)

	err = cmd.Execute()
	assert.NoError(t, err, "Expected no error")
	assert.FileExists(t, "config.yaml", "Expected file to be created")

	shortDescription := cmd.Short
	longDescription := cmd.Long
	assert.Greater(t, longDescription, shortDescription)

	os.Remove("config.yaml")
}
