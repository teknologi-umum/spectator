package file_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestListFiles_EmptyList(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	randomSessionID, _ := uuid.NewRandom()
	listFiles, err := deps.ListFiles(ctx, randomSessionID)
	if err != nil {
		t.Errorf("finding list files: %v", err)
	}

	if len(listFiles) != 0 {
		t.Errorf("list files was expected to be empty")
	}
}
