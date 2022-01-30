package file_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestListFiles(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id := globalID

	result, err := deps.ListFiles(ctx, id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	pathJSON, err := filepath.Glob("./results/*_*.json")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	pathCSV, err := filepath.Glob("./results/*_*.csv")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(result) != 250 {
		t.Errorf("Expected 250 file, got %d", len(result))
	}

	for _, i := range append(pathJSON, pathCSV...) {
		err = os.Remove(i)
		if err != nil {
			t.Errorf("removing a file: %v", err)
			return
		}
	}
}
