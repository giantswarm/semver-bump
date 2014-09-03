package commands

import (
	"fmt"
	"log"

	"github.com/coreos/go-semver/semver"
	"github.com/spf13/cobra"
)

var initialVersionString string

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initalize version number",
	Long:  `Initalize the version number for the project either from 0.1.0 or a custom one.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := getVersionStorage()

		if err != nil {
			log.Fatal(err)
		}

		if s.VersionFileExists(versionFile) {
			log.Fatal("version file already exists.")
		}

		initialVersion, err := semver.NewVersion(initialVersionString)

		if err != nil {
			log.Fatal(err)
		}

		s.WriteVersionFile(versionFile, *initialVersion)

		fmt.Printf("Bumped initial version to %s\n", initialVersion.String())
	},
}
