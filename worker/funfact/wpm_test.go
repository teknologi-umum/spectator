package funfact_test

import (
	"context"
	"testing"
	"time"
)

func TestCalculateWordsPerMinute(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	res := make(chan uint32, 1)
	err := deps.CalculateWordsPerMinute(ctx, globalID, res)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	t.Logf("average wpm: %v", <-res)
}

func TestCalculateWordsPerMinute_Forfeit(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	res := make(chan uint32, 1)
	err := deps.CalculateWordsPerMinute(ctx, globalID2, res)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	t.Logf("average wpm: %v", <-res)
}
