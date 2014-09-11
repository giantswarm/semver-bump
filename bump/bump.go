package bump

import (
	"github.com/coreos/go-semver/semver"
	"github.com/giantswarm/semver-bump/storage"
	"github.com/juju/errgo/errors"
)

type versionBumpCallback func(version *semver.Version)

type SemverBumper struct {
	storage     storage.VersionStorage
	versionFile string
}

func (sb SemverBumper) BumpMajorVersion() (*semver.Version, error) {
	v, err := sb.updateVersionFile(func(version *semver.Version) {
		version.BumpMajor()
	})

	if err != nil {
		return nil, errors.Mask(err)
	}

	return v, nil
}

func (sb SemverBumper) BumpMinorVersion() (*semver.Version, error) {
	v, err := sb.updateVersionFile(func(version *semver.Version) {
		version.BumpMinor()
	})

	if err != nil {
		return nil, errors.Mask(err)
	}

	return v, nil
}

func (sb SemverBumper) BumpPatchVersion() (*semver.Version, error) {
	v, err := sb.updateVersionFile(func(version *semver.Version) {
		version.BumpPatch()
	})

	if err != nil {
		return nil, errors.Mask(err)
	}

	return v, nil
}

func (sb SemverBumper) GetCurrentBumpedVersion() (*semver.Version, error) {
	currentVersion, err := sb.storage.ReadVersionFile(sb.versionFile)

	if err != nil {
		return nil, errors.Mask(err)
	}

	return currentVersion, nil
}

func (sb SemverBumper) InitVersion(initialVersion *semver.Version) error {
	if sb.storage.VersionFileExists(sb.versionFile) {
		return errors.Newf("Version file exists. Looks like this project is already initialized.")
	}

	err := sb.storage.WriteVersionFile(sb.versionFile, *initialVersion)

	if err != nil {
		errors.Mask(err)
	}

	return nil
}

func (sb SemverBumper) updateVersionFile(bumpCallback versionBumpCallback) (*semver.Version, error) {
	currentVersion, err := sb.GetCurrentBumpedVersion()

	if err != nil {
		return nil, errors.Mask(err)
	}

	bumpedVersion := *currentVersion

	bumpCallback(&bumpedVersion)

	err = sb.storage.WriteVersionFile(sb.versionFile, bumpedVersion)

	if err != nil {
		return nil, errors.Mask(err)
	}

	return &bumpedVersion, nil
}

func NewSemverBumper(vs storage.VersionStorage, versionFile string) *SemverBumper {
	return &SemverBumper{vs, versionFile}
}
