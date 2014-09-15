package bump

import (
	"testing"

	"github.com/giantswarm/semver-bump/storage"
)

func NewTestSemverBumper(t *testing.T, initialVersion string) *SemverBumper {
	s, err := storage.NewVersionStorageLocal(initialVersion)

	if err != nil {
		t.Fatalf("NewVersionStorageLocal: %s", err)
	}

	return NewSemverBumper(s, "testfile")
}

func TestBumpMajorVersion(t *testing.T) {
	sb := NewTestSemverBumper(t, "1.0.0")

	v, err := sb.BumpMajorVersion("", "")

	if err != nil {
		t.Fatalf("BumpMajorVersion: %s", err)
	}

	expectedVersion := "2.0.0"

	if expectedVersion != v.String() {
		t.Fatalf("BumpMajorVersion: Expected bumping of major version would result in %s but got %s", expectedVersion, v.String())
	}
}

func TestBumpMinorVersion(t *testing.T) {
	sb := NewTestSemverBumper(t, "1.0.0")

	v, err := sb.BumpMinorVersion("", "")

	if err != nil {
		t.Fatalf("BumpMinorVersion: %s", err)
	}

	expectedVersion := "1.1.0"

	if expectedVersion != v.String() {
		t.Fatalf("BumpMinorVersion: Expected bumping of minor version would result in %s but got %s", expectedVersion, v.String())
	}
}

func TestBumpPatchVersion(t *testing.T) {
	sb := NewTestSemverBumper(t, "1.0.0")

	v, err := sb.BumpPatchVersion("", "")

	if err != nil {
		t.Fatalf("BumpPatchVersion: %s", err)
	}

	expectedVersion := "1.0.1"

	if expectedVersion != v.String() {
		t.Fatalf("BumpPatchVersion: Expected bumping of patch version would result in %s but got %s", expectedVersion, v.String())
	}
}

func TestGetCurrentVersion(t *testing.T) {
	expectedVersion := "2.13.4"
	sb := NewTestSemverBumper(t, expectedVersion)

	v, err := sb.GetCurrentVersion()

	if err != nil {
		t.Fatalf("GetCurrentVersion: %s", err)
	}

	if expectedVersion != v.String() {
		t.Fatalf("GetCurrentVersion: Epexcted to receive version %s but got %s", expectedVersion, v.String())
	}
}
