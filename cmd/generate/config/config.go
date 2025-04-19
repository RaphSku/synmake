package config

import (
	c2 "github.com/RaphSku/synmake/internal/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

type GenerateConfigSubCmd struct {
	logger *zap.Logger
}

func NewGenerateConfigSubCmd(logger *zap.Logger) *GenerateConfigSubCmd {
	return &GenerateConfigSubCmd{
		logger: logger,
	}
}

func (gcsc *GenerateConfigSubCmd) GetGenerateConfigSubCmd() cobra.Command {
	generateConfigCmd := cobra.Command{
		Use:   "config",
		Short: "Generate a basic config.yaml for customization.",
		Long:  `Generate a basic config.yaml, customize it and then apply it.`,
		Run:   gcsc.runGenerateConfigCmd,
	}

	return generateConfigCmd
}

func (gcsc *GenerateConfigSubCmd) runGenerateConfigCmd(cmd *cobra.Command, args []string) {
	file, err := os.OpenFile("config.yaml", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		gcsc.logger.Error("Failed to create a config.yaml", zap.String("func", "runGenerateConfigCmd"), zap.Error(err))
		os.Exit(1)
	}
	err = c2.GenerateExampleYamlConfig(gcsc.logger, file)
	if err != nil {
		os.Remove(file.Name())
		gcsc.logger.Error("Failed to create a valid config.yaml", zap.String("func", "runGenerateConfigCmd"), zap.Error(err))
		os.Exit(1)
	}
}
