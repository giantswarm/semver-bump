package main

import (
	"github.com/giantswarm/semver-bump/commands"
	"github.com/giantswarm/semver-bump/storage"
)

func main() {
	storageFactory := storage.VersionStorageFactory{StorageType: "file", DefaultVersion: "0.0.2"}
	commands.Execute(storageFactory.GetVersionStorage())
}
