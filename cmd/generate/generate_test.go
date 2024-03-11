package generate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfGenerateCmdIsRunning(t *testing.T) {
	cmd := GetGenerateCmd()

	err := cmd.Execute()
	assert.NoError(t, err)

	shortDescription := cmd.Short
	longDescription := cmd.Long
	assert.Greater(t, len(longDescription), len(shortDescription))
}
