package config

import (
	"os"

	c2 "github.com/RaphSku/synmake/internal/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func GetGenerateConfigCmd(logger *zap.Logger) cobra.Command {
	generateConfigCmd := cobra.Command{
		Use:   "config",
		Short: "Generate a basic config.yaml for customization.",
		Long:  `Generate a basic config.yaml, customize it and then apply it.`,
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.Create("config.yaml")
			if err != nil {
				logger.Error("Failed to create a config.yaml", zap.Error(err))
				os.Exit(1)
			}

			content := "---\n"
			defaultConfig := c2.GetConfigAsYamlString(logger)
			content += defaultConfig

			_, err = file.WriteString(content)
			if err != nil {
				logger.Error("Failed to write to config.yaml", zap.Error(err))
				os.Exit(1)
			}
		},
	}

	return generateConfigCmd
}
