package config

import (
	"os"

	"github.com/RaphSku/synmake/internal/templates"
	"go.uber.org/zap"
)

func (c *ConfigManager) Apply() error {
	err := c.applyConfig()
	if err != nil {
		c.logger.Error("config could not be applied", zap.Error(err))
		os.Exit(1)
	}

	return nil
}

func (c *ConfigManager) applyConfig() error {
	file, err := os.Create("Makefile")
	if err != nil {
		c.logger.Error("error creating file", zap.Error(err))
		return err
	}
	defer file.Close()

	cnt := newContent(c.logger, c.config)
	if c.config.HelpTemplate.Enabled {
		cnt.assembleDefaultToContent()
	}
	cnt.assembleTargetsToContent()
	if c.config.HelpTemplate.Enabled {
		helpTemplate := templates.GetHelpTemplate()
		cnt.Help = helpTemplate
	}
	if c.config.VersionTemplate.Enabled {
		versionTemplate, versionConstants := templates.GetVersionTemplate(c.config.VersionTemplate.Library, c.config.VersionTemplate.MinVersion)
		cnt.Constants = append(cnt.Constants, versionConstants...)
		cnt.Preflight = append(cnt.Preflight, versionTemplate)
	}
	cnt.assemblePhonyToContent()

	content := cnt.assembleAll()
	_, err = file.WriteString(content)
	if err != nil {
		c.logger.Error("error writing to file", zap.Error(err))
		return err
	}

	c.logger.Info("file has been created successfully")

	return nil
}
