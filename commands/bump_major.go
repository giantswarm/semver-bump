package commands

import (
	"log"

	"github.com/coreos/go-semver/semver"
	"github.com/spf13/cobra"
)

var bumpMajorCommand = &cobra.Command{
	Use:   "major-release",
	Short: "Bump a major release",
	Long:  `Increments the major version and bumps it.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := getVersionStorage()

		if err != nil {
			log.Fatal(err)
		}

		err = readModifyWriteVersionFile(s, func(version *semver.Version) {
			version.BumpMajor()
		})

		if err != nil {
			log.Fatal(err)
		}

	},
}
