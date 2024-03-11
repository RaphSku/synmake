package cmd

import (
	"os"

	"github.com/RaphSku/synmake/cmd/generate"
	"github.com/RaphSku/synmake/cmd/generate/config"
	"github.com/RaphSku/synmake/cmd/version"
	c2 "github.com/RaphSku/synmake/internal/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type CLI struct {
	logger *zap.Logger

	rootCmd *cobra.Command
}

func NewCLI(logger *zap.Logger) *CLI {
	var configFilePath string

	rootCmd := &cobra.Command{
		Use:   "synmake",
		Short: "synmake helps you with the setup of Makefiles!",
		Long:  `Based on a yaml config, synmake will generate a Makefile template for you!`,
		Run: func(cmd *cobra.Command, args []string) {
			if configFilePath != "" {
				configManager := c2.NewConfigManager(logger, configFilePath)
				err := configManager.Parse().Apply()
				if err != nil {
					logger.Error("The config file could not be parsed and applied", zap.Error(err))
					os.Exit(1)
				}
			}
		},
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().StringVarP(&configFilePath, "config", "", "", "Specify the filepath to your config yaml.")
	rootCmd.MarkFlagRequired("config")

	return &CLI{
		logger:  logger,
		rootCmd: rootCmd,
	}
}

func (cli *CLI) AddSubCommands() {
	versionCmd := version.GetVersionCmd(cli.logger)
	cli.rootCmd.AddCommand(versionCmd)

	generateCmd := generate.GetGenerateCmd()
	cli.rootCmd.AddCommand(&generateCmd)
	configCmd := config.GetGenerateConfigCmd(cli.logger)
	generateCmd.AddCommand(&configCmd)
}

func (cli *CLI) Execute() {
	if err := cli.rootCmd.Execute(); err != nil {
		cli.logger.Info("CLI failed to run", zap.Error(err))
		os.Exit(1)
	}
}
