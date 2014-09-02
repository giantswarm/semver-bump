package commands

import (
	"fmt"

	"github.com/coreos/go-semver/semver"
	"github.com/giantswarm/semver-bump/storage"
)

type versionBumpCallback func(version *semver.Version)

func readModifyWriteVersionFile(versionStorage storage.VersionStorage, bumpCallback versionBumpCallback) error {
	currentVersion, err := versionStorage.ReadVersionFile(versionFile)

	if err != nil {
		return err
	}

	bumpedVersion := *currentVersion

	bumpCallback(&bumpedVersion)

	err = versionStorage.WriteVersionFile(versionFile, bumpedVersion)

	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Bumped version from %s to %s", currentVersion.String(), bumpedVersion.String()))

	return nil
}

func getVersionStorage() storage.VersionStorage {
	switch versionStorageType {
	case "local":
		return storage.NewVersionStorageLocal(versionStorageLocalDefaultVersion)
	case "file":
		return &storage.VersionStorageFile{}
	default:
		panic("Unknown storage backend: " + versionStorageType)
	}
}
