package version

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	CHANGELOG_FILEPATH = "CHANGELOG.md"
)

func scanFileForVersions(scanner *bufio.Scanner) string {
	var version string
	pattern := `v\d+\.\d+\.\d+`
	re := regexp.MustCompile(pattern)
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllString(line, -1)
		matchLength := len(matches)
		if matchLength != 0 {
			version = matches[0]
			break
		}
	}

	return version
}

func GetVersionCmd(logger *zap.Logger) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of synmake.",
		Long:  `This will show you the version of synmake in the format: {MAJOR}-{MINOR}-{PATCH}.`,
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.Open(CHANGELOG_FILEPATH)
			if err != nil {
				logger.Error("Could not open CHANGELOG", zap.Error(err))
				os.Exit(1)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			version := scanFileForVersions(scanner)

			fmt.Println(version)

			if err := scanner.Err(); err != nil {
				logger.Error("Error scanning file", zap.Error(err))
				os.Exit(1)
			}
		},
	}

	return versionCmd
}
