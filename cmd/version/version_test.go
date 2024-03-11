package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestIfVersionCmdIsRunning(t *testing.T) {
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	cmd := GetVersionCmd(logger)

	err = cmd.Execute()
	assert.NoError(t, err)

	shortDescription := cmd.Short
	longDescription := cmd.Long
	assert.Greater(t, len(longDescription), len(shortDescription))
}
