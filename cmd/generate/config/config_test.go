package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGenerateConfigCmd(t *testing.T) {
	logger := zap.NewExample()

	cmd := GetGenerateConfigCmd(logger)
	args := []string{}

	fmt.Println(cmd.Runnable())
	if cmd.Runnable() {
		err := cmd.RunE(cmd, args)
		assert.NoError(t, err)
	}
}
