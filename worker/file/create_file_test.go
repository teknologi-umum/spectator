package file_test

import "testing"

func TestCreateFile(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()

	deps.CreateFile("TESTING", globalID)
}
