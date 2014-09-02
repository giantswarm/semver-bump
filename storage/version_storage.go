package storage

import "github.com/coreos/go-semver/semver"

type VersionStorage interface {
	ReadVersionFile(file string) (*semver.Version, error)

	WriteVersionFile(file string, version semver.Version) error

	VersionFileExists(file string) bool
}
