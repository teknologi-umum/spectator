package file_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestListFiles(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id, err := uuid.NewUUID()
	if err != nil {
		t.Errorf("failed to generate uuid: %v", err)
	}

	result, err := deps.ListFiles(ctx, id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	pathJSON, err := filepath.Glob("./*_*.json")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	pathCSV, err := filepath.Glob("./*_*.csv")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(result) != 50 {
		t.Errorf("Expected 50 file, got %d", len(result))
	}

	for _, i := range append(pathJSON, pathCSV...) {
		err = os.Remove(i)
		if err != nil {
			t.Errorf("removing a file: %v", err)
			return
		}
	}
}
