package version

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestIfVersionCmdIsRunning(t *testing.T) {
	changelog := `v0.1.0`
	file, err := os.Create("CHANGELOG.md")
	assert.NoError(t, err)
	defer file.Close()
	_, err = file.WriteString(changelog)
	assert.NoError(t, err)

	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)

	cmd := GetVersionCmd(logger)

	err = cmd.Execute()
	assert.NoError(t, err)

	shortDescription := cmd.Short
	longDescription := cmd.Long
	assert.Greater(t, len(longDescription), len(shortDescription))

	os.Remove("CHANGELOG.md")
}
