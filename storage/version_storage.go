package storage

import "github.com/coreos/go-semver/semver"

type VersionStorage interface {
	ReadVersionFile(file string) (*semver.Version, error)

	WriteVersionFile(version *semver.Version, file string) error
}
