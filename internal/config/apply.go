package config

import (
	"fmt"
	"io"

	"github.com/RaphSku/synmake/internal/functions"
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

	var requireToolFunctionContent string
	if cm.Config.Functions.RequireToolFunction.Enabled {
		requireToolFunctionContent = functions.GetRequireToolDefineFunction()
	}

	var requireEnvFunctionContent string
	if cm.Config.Functions.RequireEnvFunction.Enabled {
		requireEnvFunctionContent = functions.GetRequireEnvFunction()
	}

	var requireFileFunctionContent string
	if cm.Config.Functions.RequireFileFunction.Enabled {
		requireFileFunctionContent = functions.GetRequireFileFunction()
	}

	var requireDirFunctionContent string
	if cm.Config.Functions.RequireDirFunction.Enabled {
		requireDirFunctionContent = functions.GetRequireDirFunction()
	}

	var requireConfirmFunctionContent string
	if cm.Config.Functions.RequireConfirmationFunction.Enabled {
		requireConfirmFunctionContent = functions.GetRequireConfirmFunction()
	}

	// --- VARIABLES
	cm.logger.Info("Variables will be added", zap.String("func", "applyConfig"))
	variableContent, err := assembleVariables(cm.Config.Variables)
	if err != nil {
		cm.logger.Error("variable content could not be assembled, check variables spec", zap.String("func", "applyConfig"), zap.Error(err))
		return fmt.Errorf("variable content could not be assembled, check variables spec: %w", err)
	}
	if !variableContent.isEmpty() {
		content = content.addSuggar(variableContent).lineBreak()
	}
	cm.logger.Info("Variables added", zap.String("func", "applyConfig"))

	// --- HELP TEMPLATE DEFAULT TARGET
	if cm.Config.Templates.HelpTargetTemplate.Enabled {
		cm.logger.Info("Help template default target will be added", zap.String("func", "applyConfig"))
		content = content.addSuggar(assembleDefaultTarget()).lineBreak().lineBreak()
		cm.logger.Info("Help template default target added", zap.String("func", "applyConfig"))
	}

	// --- FUNCTIONS
	if cm.Config.Functions.RequireToolFunction.Enabled {
		cm.logger.Info("Require tool function will be added", zap.String("func", "applyConfig"))
		content = content.addSuggar(assembleRequireToolFunction(requireToolFunctionContent)).lineBreak()
		cm.logger.Info("Require tool function will be added", zap.String("func", "applyConfig"))
	}

	if cm.Config.Functions.RequireEnvFunction.Enabled {
		cm.logger.Info("Require env function will be added", zap.String("func", "applyConfig"))
		content = content.addSuggar(assembleRequireEnvFunction(requireEnvFunctionContent)).lineBreak()
		cm.logger.Info("Require env function will be added", zap.String("func", "applyConfig"))
	}

	if cm.Config.Functions.RequireFileFunction.Enabled {
		cm.logger.Info("Require file function will be added", zap.String("func", "applyConfig"))
		content = content.addSuggar(assembleRequireFileFunction(requireFileFunctionContent)).lineBreak()
		cm.logger.Info("Require file function will be added", zap.String("func", "applyConfig"))
	}

	if cm.Config.Functions.RequireDirFunction.Enabled {
		cm.logger.Info("Require directory function will be added", zap.String("func", "applyConfig"))
		content = content.addSuggar(assembleRequireDirFunction(requireDirFunctionContent)).lineBreak()
		cm.logger.Info("Require directory function will be added", zap.String("func", "applyConfig"))
	}

	if cm.Config.Functions.RequireConfirmationFunction.Enabled {
		cm.logger.Info("Require confirmation function will be added", zap.String("func", "applyConfig"))
		content = content.addSuggar(assembleRequireConfirmationFunction(requireConfirmFunctionContent)).lineBreak()
		cm.logger.Info("Require confirmation function will be added", zap.String("func", "applyConfig"))
	}

	// --- PREFLIGHT
	cm.logger.Info("Preflight target will be added", zap.String("func", "applyConfig"))
	commands := []string{}
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
	_, err = w.WriteString(content.getString())
	if err != nil {
		cm.logger.Error("Failed to write to Makefile", zap.String("func", "applyConfig"), zap.Error(err))
		return fmt.Errorf("failed to write to Makefile: %w", err)
	}
	cm.logger.Info("Makefile has been created successfully!", zap.String("func", "applyConfig"))

	return nil
}
