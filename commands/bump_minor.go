package commands

import (
	"log"

	"github.com/coreos/go-semver/semver"
	"github.com/spf13/cobra"
)

var bumpMinorCommand = &cobra.Command{
	Use:   "minor-release",
	Short: "Bump a minor release",
	Long:  `Increments the minor version and bumps it.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := getVersionStorage()

		if err != nil {
			log.Fatal(err)
		}

		err = readModifyWriteVersionFile(s, func(version *semver.Version) {
			version.BumpMinor()
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}
