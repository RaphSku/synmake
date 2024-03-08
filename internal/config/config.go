package config

import (
	"bytes"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Target struct {
	HelpDescription string   `yaml:"helpDescription"`
	PreTargets      []string `yaml:"preTargets,omitempty"`
	Commands        []string `yaml:"commands"`
	Display         bool     `yaml:"display"`
}

type Config struct {
	Phony        []string          `yaml:"phony"`
	Targets      map[string]Target `yaml:"targets"`
	HelpTemplate struct {
		Enabled   bool   `yaml:"enabled"`
		Delimiter string `yaml:"delimiter"`
	} `yaml:"helpTemplate"`
	VersionTemplate struct {
		Enabled    bool   `yaml:"enabled"`
		Library    string `yaml:"library"`
		MinVersion string `yaml:"minVersion"`
	} `yaml:"versionTemplate"`
}

type ConfigManager struct {
	logger   *zap.Logger
	config   Config
	filePath string
}

func NewConfigManager(logger *zap.Logger, filePath string) *ConfigManager {
	return &ConfigManager{
		logger:   logger,
		config:   Config{},
		filePath: filePath,
	}
}

func GetConfigAsYamlString(logger *zap.Logger) string {
	tA := Target{
		HelpDescription: "targetA just prints an output",
		Commands:        []string{"echo \"Hello World\"", "echo \"This is how you specify commands!\""},
		Display:         false,
	}

	tB := Target{
		HelpDescription: "targetB just prints an output",
		PreTargets:      []string{"targetA"},
		Commands:        []string{"echo \"This is targetB!\"", "echo \"How are you doing?\""},
		Display:         true,
	}

	config := Config{
		Phony: []string{"default", "preflight", "targetA", "targetB", "help"},
		Targets: map[string]Target{
			"targetA": tA,
			"targetB": tB,
		},
		HelpTemplate: struct {
			Enabled   bool   "yaml:\"enabled\""
			Delimiter string "yaml:\"delimiter\""
		}{
			Enabled:   true,
			Delimiter: "##",
		},
		VersionTemplate: struct {
			Enabled    bool   "yaml:\"enabled\""
			Library    string "yaml:\"library\""
			MinVersion string "yaml:\"minVersion\""
		}{
			Enabled:    true,
			Library:    "example",
			MinVersion: "0.1.0",
		},
	}

	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	err := encoder.Encode(config)
	if err != nil {
		logger.Error("Config could not be marshalled", zap.Error(err))
		os.Exit(1)
	}

	return buf.String()
}
