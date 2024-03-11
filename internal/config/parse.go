package config

import (
	"os"

	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v3"
)

func (c *ConfigManager) Parse() *ConfigManager {
	err := c.parseConfig()
	if err != nil {
		c.logger.Error("config file could not be parsed successfully", zap.Error(err))
		os.Exit(1)
	}

	return c
}

func (c *ConfigManager) parseConfig() error {
	data, err := os.ReadFile(c.filePath)
	if err != nil {
		c.logger.Error("data could not be read", zap.Error(err))
		return err
	}

	err = yaml.Unmarshal(data, &c.config)
	if err != nil {
		c.logger.Error("data could not be parsed into a valid Config", zap.Error(err))
		return err
	}

	return nil
}
