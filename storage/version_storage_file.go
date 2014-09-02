package storage

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/coreos/go-semver/semver"
	"github.com/juju/errgo/errors"
)

type VersionStorageFile struct{}

func (s VersionStorageFile) ReadVersionFile(file string) (*semver.Version, error) {
	versionBuffer, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, errors.Mask(err)
	}

	versionBuffer = bytes.TrimSpace(versionBuffer)

	version, err := semver.NewVersion(string(versionBuffer))

	if err != nil {
		return nil, errors.Mask(err)
	}

	return version, nil
}

func (s VersionStorageFile) WriteVersionFile(file string, version semver.Version) error {
	return errors.Mask(ioutil.WriteFile(file, []byte(version.String()), os.ModeExclusive))
}
