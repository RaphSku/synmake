package config

import (
	"os"

	c2 "github.com/RaphSku/synmake/internal/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	scopedLogger *zap.Logger
)

func generateConfigResolver(cmd *cobra.Command, args []string) {
	file, err := os.Create("config.yaml")
	if err != nil {
		scopedLogger.Error("Failed to create a config.yaml", zap.Error(err))
		os.Exit(1)
	}

	content := c2.GetConfigAsYamlString(scopedLogger)

	_, err = file.WriteString(content)
	if err != nil {
		scopedLogger.Error("Failed to write to config.yaml", zap.Error(err))
		os.Exit(1)
	}
}

func GetGenerateConfigCmd(logger *zap.Logger) *cobra.Command {
	scopedLogger = logger

	generateConfigCmd := &cobra.Command{
		Use:   "config",
		Short: "Generate a basic config.yaml for customization.",
		Long:  `Generate a basic config.yaml, customize it and then apply it.`,
		Run:   generateConfigResolver,
	}

	return generateConfigCmd
}
