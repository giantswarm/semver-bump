package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var bumpMinorCommand = &cobra.Command{
	Use:   "minor-release",
	Short: "Bump a minor release",
	Long:  `Increments the minor version and bumps it.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentVersion, err := versionStorage.ReadVersionFile(versionFile)

		if err != nil {
			log.Fatal(err)
		}

		bumpedVersion := *currentVersion

		bumpedVersion.BumpMinor()

		err = versionStorage.WriteVersionFile(bumpedVersion, versionFile)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Bumped minor version from %s to %s", currentVersion.String(), bumpedVersion.String())
	},
}
