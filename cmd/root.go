package cmd

import (
	"fmt"
	"os"

	"github.com/RaphSku/synmake/cmd/generate"
	"github.com/RaphSku/synmake/cmd/generate/config"
	"github.com/RaphSku/synmake/cmd/version"
	c2 "github.com/RaphSku/synmake/internal/config"
	"github.com/RaphSku/synmake/internal/logging"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type CLI struct {
	logger *zap.Logger

	rootCmd *cobra.Command
}

func NewCLI() *CLI {
	return &CLI{}
}

func (cli *CLI) Run() {
	rootCmd := &cobra.Command{
		Use:   "synmake",
		Short: "synmake helps you with the setup of Makefiles!",
		Long:  `Based on a yaml config, synmake will generate a Makefile template for you!`,
		Run:   cli.runSynmakeCommand,
	}

	// --- ROOT CMD
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().StringP("config", "", "", "Specify the filepath to your config yaml.")
	rootCmd.MarkFlagRequired("config")
	rootCmd.Flags().BoolP("debug", "", false, "Enable debug mode to get more helpful log messages.")
	cli.rootCmd = rootCmd

	// --- SUB CMD
	versionCmd := version.NewVersionCmd().GetVersionCmd()
	cli.rootCmd.AddCommand(versionCmd)

	generateCmd := generate.NewGenerateConfigCmd().GetGenerateConfigCmd()
	cli.rootCmd.AddCommand(&generateCmd)

	configCmd := config.NewGenerateConfigSubCmd(cli.logger).GetGenerateConfigSubCmd()
	generateCmd.AddCommand(&configCmd)

	// --- EXECUTE
	if err := cli.rootCmd.Execute(); err != nil {
		fmt.Printf("CLI failed to run due to %v\n", err)
	}
}

func (cli *CLI) runSynmakeCommand(cmd *cobra.Command, args []string) {
	debugMode, _ := cmd.Flags().GetBool("debug")
	cli.logger = logging.SetupZapLogger(debugMode)
	configFilePath, _ := cmd.Flags().GetString("config")

	file, err := os.OpenFile(configFilePath, os.O_RDONLY, 0644)
	if err != nil {
		cli.logger.Error("Could not create config file", zap.String("func", "runSynmakeCommand"), zap.Error(err))
		os.Exit(1)
	}
	configManager, err := c2.NewConfigManager(cli.logger, file)
	if err != nil {
		cli.logger.Error("Failed to initialize a new config manager", zap.String("func", "runSynmakeCommand"), zap.Error(err))
		os.Exit(1)
	}
	defer configManager.Close()

	err = configManager.Parse()
	if err != nil {
		cli.logger.Error("Parsing of the config file failed", zap.String("func", "runSynmakeCommand"), zap.Error(err))
		os.Exit(1)
	}

	file, err = os.OpenFile("Makefile", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		cli.logger.Error("Could not create Makefile", zap.String("func", "runSynmakeCommand"), zap.Error(err))
		os.Exit(1)
	}
	defer file.Close()
	err = configManager.Apply(file)
	if err != nil {
		cli.logger.Error("The config file could not be applied", zap.String("func", "runSynmakeCommand"), zap.Error(err))
		os.Exit(1)
	}
	return
}
