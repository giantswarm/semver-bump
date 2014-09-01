package commands

import (
	"github.com/giantswarm/semver-bump/storage"
	"github.com/spf13/cobra"
)

var versionFile string
var versionStorage storage.VersionStorage

var SemverBumpCommand = &cobra.Command{
	Use:   "semver-bump",
	Short: "Semantic Versioning Bumper",
	Long:  `A semantic versioning file bumper built by giantswarm`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute(storage storage.VersionStorage) {
	versionStorage = storage

	AddGlobalFlags()
	AddCommands()

	SemverBumpCommand.Execute()
}

func AddCommands() {
	SemverBumpCommand.AddCommand(bumpMajorCommand)
}

func AddGlobalFlags() {
	SemverBumpCommand.PersistentFlags().StringVarP(&versionFile, "version-file", "f", "VERSION", "Version file to use")
}
