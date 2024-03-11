package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestAssemblePhonyToContent(t *testing.T) {
	t.Parallel()

	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	exp_phony := ".PHONY: default help\n"

	var config Config
	config.Phony = []string{"default", "help"}
	cnt := newContent(logger, config)
	cnt.assemblePhonyToContent()

	assert.Equal(t, exp_phony, cnt.Phony)
}

func TestDefaultToContent(t *testing.T) {
	t.Parallel()

	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	exp_default := "default: help"

	var config Config
	cnt := newContent(logger, config)
	cnt.assembleDefaultToContent()

	assert.Equal(t, exp_default, cnt.Default)
}

func TestAssembleTargetsToContent(t *testing.T) {
	t.Parallel()

	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	exp_target := `targetA: default preflight
	echo "Test this"

`

	var config Config

	targetA := Target{
		HelpDescription: "help test",
		PreTargets:      []string{"default", "preflight"},
		Commands:        []string{"echo \"Test this\""},
		Display:         false,
	}

	config.Targets = map[string]Target{"targetA": targetA}
	cnt := newContent(logger, config)
	cnt.assembleTargetsToContent()

	assert.Equal(t, exp_target, cnt.Targets)
}

func TestAssembleAll(t *testing.T) {
	t.Parallel()

	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	exp_content := `




preflight:
	`

	var config Config
	config.Phony = []string{"help"}
	targetA := Target{
		HelpDescription: "help test",
		PreTargets:      []string{"default", "preflight"},
		Commands:        []string{"echo \"Test this\""},
		Display:         false,
	}

	config.Targets = map[string]Target{"targetA": targetA}
	cnt := newContent(logger, config)
	content := cnt.assembleAll()

	assert.Equal(t, exp_content, content)
}
