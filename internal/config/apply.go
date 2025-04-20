package config

import (
	"io"

	"github.com/RaphSku/synmake/internal/templates"
	"go.uber.org/zap"
)

func (cm *ConfigManager) Apply(file FileInterface) error {
	err := cm.applyConfig(file)
	if err != nil {
		cm.logger.Error("Config could not be applied", zap.String("func", "Apply"), zap.Error(err))
		return err
	}

	return nil
}

func (cm *ConfigManager) applyConfig(w io.StringWriter) error {
	content := &SuggarString{}

	var versionContent string
	versionConstants := []string{}
	if cm.Config.Templates.VersionCommandTemplate.Enabled {
		versionContent, versionConstants = templates.GetVersionTemplate(
			cm.Config.Templates.VersionCommandTemplate.Library,
			cm.Config.Templates.VersionCommandTemplate.MinVersion,
		)
	}

	// --- VARIABLES
	cm.logger.Info("Variables will be added", zap.String("func", "applyConfig"))
	variables := []string{}
	if cm.Config.Templates.VersionCommandTemplate.Enabled {
		for _, constant := range versionConstants {
			variables = append(variables, constant)
		}
	}
	content = content.addSuggar(assembleVariables(variables)).lineBreak()
	cm.logger.Info("Variables added", zap.String("func", "applyConfig"))

	// --- HELP TEMPLATE DEFAULT TARGET
	if cm.Config.Templates.HelpTargetTemplate.Enabled {
		cm.logger.Info("Help template default target will be added", zap.String("func", "applyConfig"))
		content = content.addSuggar(assembleDefaultTarget()).lineBreak().lineBreak()
		cm.logger.Info("Help template default target added", zap.String("func", "applyConfig"))
	}

	// --- PREFLIGHT
	cm.logger.Info("Preflight target will be added", zap.String("func", "applyConfig"))
	commands := []string{}
	if cm.Config.Templates.VersionCommandTemplate.Enabled {
		commands = append(
			commands,
			versionContent,
		)
	}
	content = content.addSuggar(assemblePreflightTarget(commands))
	cm.logger.Info("Preflight target added", zap.String("func", "applyConfig"))

	// --- TARGETS ADDED
	cm.logger.Info("Targets will be added", zap.String("func", "applyConfig"))
	delimiter := "#"
	if cm.Config.Templates.HelpTargetTemplate.Enabled {
		delimiter = cm.Config.Templates.HelpTargetTemplate.Delimiter
	}
	content = content.addSuggar(assembleTargets(cm.Config.Targets, delimiter))
	cm.logger.Info("Targets added", zap.String("func", "applyConfig"))

	// --- HELP TARGET ADDED
	if cm.Config.Templates.HelpTargetTemplate.Enabled {
		cm.logger.Info("Help target will be added", zap.String("func", "applyConfig"))
		content = content.appendString(templates.GetHelpTemplate())
		cm.logger.Info("Help target added", zap.String("func", "applyConfig"))
	}

	// --- WRITE MAKEFILE
	cm.logger.Info("Writing to Makefile", zap.String("func", "applyConfig"))
	_, err := w.WriteString(content.getString())
	if err != nil {
		cm.logger.Error("Failed to write to Makefile", zap.String("func", "applyConfig"), zap.Error(err))
		return err
	}
	cm.logger.Info("Makefile has been created successfully!", zap.String("func", "applyConfig"))

	return nil
}
