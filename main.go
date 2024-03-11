package main

import (
	"github.com/RaphSku/synmake/cmd"
	"github.com/RaphSku/synmake/internal/logging"
)

func main() {
	logger := logging.SetupZapLogger()

	cli := cmd.NewCLI(logger)
	cli.AddSubCommands()
	cli.Execute()
}
