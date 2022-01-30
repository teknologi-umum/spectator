package funfact_test

import (
	"context"
	"testing"
	"time"
)

func TestCalculateDeletionRate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	res := make(chan float32, 1)
	defer close(res)
	err := deps.CalculateDeletionRate(ctx, globalID, res)
	if err != nil {
		t.Fatalf("an error was thrown: %v", err)
	}

	deletionRate := <-res
	if deletionRate != 0.14285715 {
		t.Errorf("expected deletionRate to be 0.14285715, got: %v", deletionRate)
	}
}

func TestCalculateDeletionRate_NoDeletion(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	res := make(chan float32, 1)
	defer close(res)
	err := deps.CalculateDeletionRate(ctx, globalID2, res)
	if err != nil {
		t.Fatalf("an error was thrown: %v", err)
	}

	deletionRate := <-res
	if deletionRate != 0 {
		t.Errorf("expected deletionRate to be 0, got: %v", deletionRate)
	}
}
