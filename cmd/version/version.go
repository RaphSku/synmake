package version

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func GetVersionCmd(logger *zap.Logger) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of synmake.",
		Long:  `This will show you the version of synmake in the format: {MAJOR}-{MINOR}-{PATCH}.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v0.1.0")
		},
	}

	return versionCmd
}
