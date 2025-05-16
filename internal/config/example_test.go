package config_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"io"
	"testing"

	"github.com/RaphSku/synmake/internal/config"
	"go.uber.org/zap"
)

func TestDefaultConfigAsYAML(t *testing.T) {
	t.Parallel()

	// --- Logger Setup
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	// --- Checking whether the generated example config can be unmarshalled
	// --- and is not empty
	buffer := &bytes.Buffer{}

	config.GenerateExampleYamlConfig(logger, buffer)

	data, err := io.ReadAll(buffer)
	assert.NoError(t, err)

	var c config.Config
	err = yaml.Unmarshal(data, &c)
	assert.NoError(t, err)
	assert.Equal(t, false, isEmptyStruct(c))
}
