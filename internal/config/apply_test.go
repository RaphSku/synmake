package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestApply(t *testing.T) {
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	configString := GetConfigAsYamlString(logger)

	file, err := os.Create("config.yaml")
	assert.NoError(t, err)
	defer file.Close()
	_, err = file.WriteString(configString)
	assert.NoError(t, err)

	cm := NewConfigManager(logger, "config.yaml")
	cm.Apply()
	assert.FileExists(t, "Makefile")

	os.Remove("Makefile")
	os.Remove("config.yaml")
}
