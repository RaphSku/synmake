package config

import (
	"go.uber.org/zap"
)

type Target struct {
	HelpDescription string   `yaml:"helpDescription"`
	PreTargets      []string `yaml:"preTargets,omitempty"`
	Commands        []string `yaml:"commands"`
	Display         bool     `yaml:"display"`
}

type HelpTemplate struct {
	Enabled   bool   `yaml:"enabled"`
	Delimiter string `yaml:"delimiter"`
}

type VersionTemplate struct {
	Enabled    bool   `yaml:"enabled"`
	Library    string `yaml:"library"`
	MinVersion string `yaml:"minVersion"`
}

type OptionalTemplates struct {
	HelpTargetTemplate     HelpTemplate    `yaml:"helpTemplate"`
	VersionCommandTemplate VersionTemplate `yaml:"versionTemplate"`
}

type Config struct {
	Targets   map[string]Target `yaml:"targets"`
	Templates OptionalTemplates `yaml:"templates"`
}

type ConfigManager struct {
	logger *zap.Logger
	Config Config
	File   FileInterface
}

func NewConfigManager(logger *zap.Logger, file FileInterface) (*ConfigManager, error) {
	return &ConfigManager{
		logger: logger,
		Config: Config{},
		File:   file,
	}, nil
}

func (cm *ConfigManager) Close() error {
	return cm.File.Close()
}
