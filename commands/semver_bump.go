package commands

import (
	"fmt"
	"log"

	"github.com/giantswarm/semver-bump/storage"
	"github.com/spf13/cobra"
)

var projectVersion string
var versionFile string
var versionStorage storage.VersionStorage

var SemverBumpCommand = &cobra.Command{
	Use:   "semver-bump",
	Short: "Semantic Versioning Bumper",
	Long:  `A semantic versioning file bumper built by giantswarm`,
	Run: func(cmd *cobra.Command, args []string) {
		currentVersion, err := versionStorage.ReadVersionFile(versionFile)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Current bumped version is: %s", currentVersion)
	},
}

func Execute(v string, storage storage.VersionStorage) {
	versionStorage = storage
	projectVersion = v

	AddGlobalFlags()
	AddCommands()

	SemverBumpCommand.Execute()
}

func AddCommands() {
	SemverBumpCommand.AddCommand(bumpMajorCommand)
	SemverBumpCommand.AddCommand(bumpMinorCommand)
	SemverBumpCommand.AddCommand(bumpPatchCommand)
	SemverBumpCommand.AddCommand(versionCommand)
}

func AddGlobalFlags() {
	SemverBumpCommand.PersistentFlags().StringVarP(&versionFile, "version-file", "f", "VERSION", "Version file to use")
}
