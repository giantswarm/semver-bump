package commands

import (
	"github.com/coreos/go-semver/semver"
	"github.com/spf13/cobra"
)

var bumpMajorCommand = &cobra.Command{
	Use:   "major-release",
	Short: "Bump a major release",
	Long:  `Increments the major version and bumps it.`,
	Run: func(cmd *cobra.Command, args []string) {
		readModifyWriteVersionFile(versionStorage, func(version *semver.Version) {
			version.BumpMajor()
		})
	},
}
