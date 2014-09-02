package commands

import (
	"github.com/coreos/go-semver/semver"
	"github.com/spf13/cobra"
)

var bumpPatchCommand = &cobra.Command{
	Use:   "patch-release",
	Short: "Bump a patch release",
	Long:  `Increments the patch version and bumps it.`,
	Run: func(cmd *cobra.Command, args []string) {
		readModifyWriteVersionFile(versionStorage, func(version *semver.Version) {
			version.BumpPatch()
		})

	},
}
