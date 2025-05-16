package generate

import (
	"github.com/spf13/cobra"
)

type GenerateConfigCmd struct{}

func NewGenerateConfigCmd() *GenerateConfigCmd {
	return &GenerateConfigCmd{}
}

func (gcc *GenerateConfigCmd) GetGenerateConfigCmd() cobra.Command {
	generateCmd := cobra.Command{
		Use:   "generate",
		Short: "The command generate can be used to generate configuration files.",
		Long:  `The command generate can be used to generate configuration files for customization.`,
	}

	return generateCmd
}
