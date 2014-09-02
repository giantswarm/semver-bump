package commands

import (
	"fmt"

	"github.com/coreos/go-semver/semver"
	"github.com/giantswarm/semver-bump/storage"
)

type versionBumpCallback func(version *semver.Version)

func readModifyWriteVersionFile(storage storage.VersionStorage, bumpCallback versionBumpCallback) error {
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
