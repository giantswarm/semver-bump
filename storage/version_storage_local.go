package storage

import (
	"log"

	"github.com/coreos/go-semver/semver"
)

type VersionStorageLocal struct {
	version *semver.Version
}

func (s VersionStorageLocal) ReadVersionFile(file string) (*semver.Version, error) {
	return s.version, nil
}

func (s VersionStorageLocal) WriteVersionFile(file string, version semver.Version) error {
	s.version = &version

	return nil
}

func NewVersionStorageLocal(versionString string) *VersionStorageLocal {
	version, err := semver.NewVersion(versionString)

	if err != nil {
		log.Fatalf("Version string '%s' could not be processed.", versionString)
	}

	return &VersionStorageLocal{version: version}
}
