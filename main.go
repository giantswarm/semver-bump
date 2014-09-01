package main

import (
	"github.com/giantswarm/semver-bump/commands"
	"github.com/giantswarm/semver-bump/storage"
)

var projectVersion string = "dev"

func main() {
	storageFactory := storage.VersionStorageFactory{StorageType: "file", DefaultVersion: "0.0.2"}
	commands.Execute(projectVersion, storageFactory.GetVersionStorage())
}
