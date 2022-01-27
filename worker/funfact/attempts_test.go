package funfact_test

import (
	"context"
	"testing"
	"time"
)

func TestCalculateSubmissionAttempts(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	res := make(chan uint32, 1)
	defer close(res)
	err := deps.CalculateSubmissionAttempts(ctx, globalID, res)
	if err != nil {
		t.Fatalf("an error was thrown: %v", err)
	}

	attempts := <-res
	if attempts != 25 {
		t.Errorf("expected attempts to be 25, got: %v", attempts)
	}
}
