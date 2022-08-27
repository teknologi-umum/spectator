package file_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCheckIfSessionExists(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	exists, err := deps.CheckIfSessionExists(ctx, globalID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !exists {
		t.Logf("expected session to exist")
	}

	randomId, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	exists, err = deps.CheckIfSessionExists(ctx, randomId)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if exists {
		t.Logf("expected session to not exist")
	}
}
