package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestParsing(t *testing.T) {
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	configString := GetConfigAsYamlString(logger)

	file, err := os.Create("config.yaml")
	assert.NoError(t, err)
	defer file.Close()
	_, err = file.WriteString(configString)
	assert.NoError(t, err)

	cm := NewConfigManager(logger, "./config.yaml")
	cm.Parse()
	assert.Equal(t, false, isEmptyStruct(cm.config))
	assert.Equal(t, 5, len(cm.config.Phony))
	assert.Equal(t, true, cm.config.HelpTemplate.Enabled)
	assert.Equal(t, "example", cm.config.VersionTemplate.Library)

	os.Remove("config.yaml")
}
