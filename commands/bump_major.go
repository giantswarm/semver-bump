package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var bumpMajorCommand = &cobra.Command{
	Use:   "major-release",
	Short: "Bump a major release",
	Long:  `Increments the major version and bumps it.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentVersion, err := versionStorage.ReadVersionFile(versionFile)

		if err != nil {
			log.Fatal(err)
		}

		bumpedVersion := *currentVersion

		bumpedVersion.BumpMajor()

		err = versionStorage.WriteVersionFile(bumpedVersion, versionFile)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Bumped major version from %s to %s", currentVersion.String(), bumpedVersion.String())
	},
}
