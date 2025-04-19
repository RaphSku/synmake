package generate_test

import (
	"testing"

	"github.com/RaphSku/synmake/cmd/generate"
	"github.com/stretchr/testify/assert"
)

func TestIfGenerateCmdIsRunning(t *testing.T) {
	t.Parallel()

	generateCommand := generate.NewGenerateConfigCmd()
	cmd := generateCommand.GetGenerateConfigCmd()

	err := cmd.Execute()
	assert.NoError(t, err)

	shortDescription := cmd.Short
	longDescription := cmd.Long
	assert.Greater(t, len(longDescription), len(shortDescription))
}
