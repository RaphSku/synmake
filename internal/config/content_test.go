package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestContentInit(t *testing.T) {
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	var config Config
	content := newContent(logger, config)
	assert.Equal(t, true, isEmptyStruct(content.config))
}

func TestConcatenation(t *testing.T) {
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	var config Config
	content := newContent(logger, config)
	act_result := content.concatStringsWithWhiteSpaces([]string{"hey", "this", "is", "a", "test"}...)
	assert.Equal(t, "hey this is a test", act_result)
}
