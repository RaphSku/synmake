package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssembleTargetsToContent(t *testing.T) {
	t.Parallel()

	// --- Checking whether targets are assembled correctly
	exp_target := ".PHONY: targetA\n help test\ntargetA: default preflight\n\techo \"Test this\"\n\n"

	var config Config
	targetA := Target{
		HelpDescription: "help test",
		PreTargets:      []string{"default", "preflight"},
		Commands:        []string{"echo \"Test this\""},
		Display:         false,
	}

	config.Targets = map[string]Target{"targetA": targetA}
	ss := assembleTargets(config.Targets, "")

	assert.Equal(t, exp_target, ss.getString())
}
