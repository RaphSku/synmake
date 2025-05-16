package config

import (
	"io"

	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v3"
)

func (cm *ConfigManager) Parse() error {
	err := cm.parseConfig(cm.File)
	if err != nil {
		cm.logger.Error("config file could not be parsed successfully", zap.String("func", "Parse"), zap.Error(err))
		return err
	}

	return nil
}

func (cm *ConfigManager) parseConfig(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		cm.logger.Error("data could not be read", zap.String("func", "parseConfig"), zap.Error(err))
		return err
	}

	err = yaml.Unmarshal(data, &cm.Config)
	if err != nil {
		cm.logger.Error("data could not be parsed into a valid Config", zap.String("func", "parseConfig"), zap.Error(err))
		return err
	}

	return nil
}
