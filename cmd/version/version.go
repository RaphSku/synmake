package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

type VersionCmd struct{}

func NewVersionCmd() *VersionCmd {
	return &VersionCmd{}
}

func (vc *VersionCmd) GetVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of synmake.",
		Long:  `This will show you the version of synmake in the format: {MAJOR}-{MINOR}-{PATCH}.`,
		Run:   vc.runVersionCmd,
	}

	return versionCmd
}

func (vc *VersionCmd) runVersionCmd(cmd *cobra.Command, args []string) {
	fmt.Println("v0.2.0")
}
