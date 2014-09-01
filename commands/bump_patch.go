package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var bumpPatchCommand = &cobra.Command{
	Use:   "patch-release",
	Short: "Bump a patch release",
	Long:  `Increments the patch version and bumps it.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentVersion, err := versionStorage.ReadVersionFile(versionFile)

		if err != nil {
			log.Fatal(err)
		}

		bumpedVersion := *currentVersion

		bumpedVersion.BumpPatch()

		err = versionStorage.WriteVersionFile(bumpedVersion, versionFile)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Bumped patch version from %s to %s", currentVersion.String(), bumpedVersion.String())
	},
}
