package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuggarInit(t *testing.T) {
	t.Parallel()

	var ss SuggarString
	assert.Equal(t, ss.value, "")

	exp_value := "Test"
	ss.value = exp_value
	assert.Equal(t, exp_value, ss.value)
}

func TestLineBreakAndTabs(t *testing.T) {
	t.Parallel()

	var ss SuggarString
	ss.lineBreak().tab()

	assert.Equal(t, "\n\t", ss.value)
}

func TestAppendString(t *testing.T) {
	t.Parallel()

	exp_value := "Testing This"
	var ss SuggarString
	ss.appendString(exp_value)

	assert.Equal(t, exp_value, ss.value)
}

func TestChaining(t *testing.T) {
	t.Parallel()

	value := "Test this"
	var ss SuggarString
	ss.lineBreak().appendString(value).tab()

	exp_value := "\nTest this\t"
	assert.Equal(t, exp_value, ss.value)
}
