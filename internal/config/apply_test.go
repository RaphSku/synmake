package config_test

import (
	"bytes"
	"testing"

	"github.com/RaphSku/synmake/internal/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestApply(t *testing.T) {
	t.Parallel()

	// --- Setup Logger
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	// --- Generate example configuration for testing
	configBuffer := &bytes.Buffer{}
	config.GenerateExampleYamlConfig(logger, configBuffer)

	// --- Initialise new config manager
	configFileBuffer := NewFileBuffer("config.yaml")
	configFileBuffer.Write(configBuffer.Bytes())
	cm, err := config.NewConfigManager(logger, configFileBuffer)
	assert.NoError(t, err)
	err = cm.Parse()
	assert.NoError(t, err)

	// --- Test generation of Makefile
	actualMakefile := NewFileBuffer("Makefile")
	err = cm.Apply(actualMakefile)
	assert.NoError(t, err)

	expectedMakefileByteLength := 1646
	assert.Equal(t, expectedMakefileByteLength, len(actualMakefile.buffer.Bytes()))
}
