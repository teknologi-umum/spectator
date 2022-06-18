package funfact_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCalculateWordsPerMinute(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	res := make(chan int64, 1)
	defer close(res)
	err := deps.CalculateWordsPerMinute(ctx, globalID, res)
	if err != nil {
		t.Fatalf("an error was thrown: %v", err)
	}

	out := <-res
	t.Logf("average wpm: %v", out)
}

func TestCalculateWordsPerMinute_Forfeit(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	res := make(chan int64, 1)
	defer close(res)
	err := deps.CalculateWordsPerMinute(ctx, globalID, res)
	if err != nil {
		t.Fatalf("an error was thrown: %v", err)
	}

	out := <-res
	t.Logf("average wpm: %v", out)
}

func TestCalculateWordsPerMinute_NotFound(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res := make(chan int64, 1)
	defer close(res)

	err := deps.CalculateWordsPerMinute(ctx, uuid.New(), res)
	if err != nil {
		t.Fatalf("an error was thrown: %v", err)
	}

	out := <-res
	if out != 0 {
		t.Fatalf("expected 0, got %v", out)
	}
}
