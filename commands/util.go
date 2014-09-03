package commands

import (
	"fmt"

	"github.com/coreos/go-semver/semver"
	"github.com/giantswarm/semver-bump/storage"
	"github.com/juju/errgo/errors"
)

type versionBumpCallback func(version *semver.Version)

func readModifyWriteVersionFile(versionStorage storage.VersionStorage, bumpCallback versionBumpCallback) error {
	currentVersion, err := versionStorage.ReadVersionFile(versionFile)

	if err != nil {
		return errors.Mask(err)
	}

	bumpedVersion := *currentVersion

	bumpCallback(&bumpedVersion)

	err = versionStorage.WriteVersionFile(versionFile, bumpedVersion)

	if err != nil {
		return errors.Mask(err)
	}

	fmt.Printf("Bumped version from %s to %s\n", currentVersion.String(), bumpedVersion.String())

	return nil
}

func getVersionStorage() (storage.VersionStorage, error) {
	switch versionStorageType {
	case "local":
		return storage.NewVersionStorageLocal(versionStorageLocalDefaultVersion)
	case "file":
		return &storage.VersionStorageFile{}, nil
	default:
		panic("Unknown storage backend: " + versionStorageType)
	}
}
