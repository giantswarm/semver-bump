package storage

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/coreos/go-semver/semver"
)

type VersionStorageFile struct{}

func (s VersionStorageFile) ReadVersionFile(file string) (*semver.Version, error) {
	versionBuffer, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, fmt.Errorf("File '%s' could not be openend", file)
	}

	versionBuffer = bytes.TrimSpace(versionBuffer)

	version, err := semver.NewVersion(string(versionBuffer))

	if err != nil {
		return nil, fmt.Errorf("Version string in file '%s' could not pe processed with error: %s", file, err)
	}

	return version, nil
}

func (s VersionStorageFile) WriteVersionFile(file string, version semver.Version) error {
	return ioutil.WriteFile(file, []byte(version.String()), os.ModeExclusive)
}
