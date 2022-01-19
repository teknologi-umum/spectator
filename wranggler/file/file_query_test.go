package file_test

import (
	"testing"
	"worker/file"
)

func TestIsDebug(t *testing.T) {
	// This is not necessary to test, but eh.
	overrideDeps := &file.Dependency{
		Environment: "DEVELOPMENT",
	}

	if overrideDeps.IsDebug() != true {
		t.Errorf("expected true, got false")
	}

	overrideDeps = &file.Dependency{
		Environment: "PRODUCTION",
	}

	if overrideDeps.IsDebug() != false {
		t.Errorf("expected false, got true")
	}
}
