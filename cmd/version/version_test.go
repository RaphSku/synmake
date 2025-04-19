package version_test

import (
	"testing"

	"github.com/RaphSku/synmake/cmd/version"
	"github.com/stretchr/testify/assert"
)

func TestIfVersionCmdIsRunning(t *testing.T) {
	t.Parallel()

	versionCommand := version.NewVersionCmd()
	cmd := versionCommand.GetVersionCmd()

	shortDescription := cmd.Short
	longDescription := cmd.Long
	assert.Greater(t, len(longDescription), len(shortDescription))
}
