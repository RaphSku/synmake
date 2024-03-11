package config

import (
	"go.uber.org/zap"
)

type content struct {
	logger *zap.Logger
	config Config

	Phony     string
	Constants []string
	Default   string
	Preflight []string
	Targets   string
	Help      string
}

func newContent(logger *zap.Logger, config Config) *content {
	return &content{
		logger: logger,
		config: config,
	}
}

func (cnt *content) concatStringsWithWhiteSpaces(strs ...string) string {
	result := ""
	for i := range strs {
		result += strs[i]
		if i != len(strs)-1 {
			result += " "
		}
	}

	return result
}
