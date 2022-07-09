package main_test

import (
	"context"
	"errors"
	"testing"
	"time"
	main "video"
)

func TestFfmpeg(t *testing.T) {
	f, err := main.NewFfmpeg()
	if err != nil {
		t.Fatal("ffmpeg was not installed on this device")
	}

	t.Run("Convert", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		out, err := f.Convert(ctx, "samples/ocean_rocks.webm", "samples/ocean_rocks.mp4")
		if err != nil {
			t.Errorf("an error was thrown: %v", err)
		}

		t.Log(string(out))
	})

	t.Run("Convert/NotExists", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		_, err := f.Convert(ctx, "", "")
		if err == nil {
			t.Errorf("was expecting an error, got nil")
		}

		if !errors.Is(err, main.ErrFfmpegError) {
			t.Errorf("expecting an error of %v, got %v instead", main.ErrFfmpegError, err)
		}
	})

	t.Run("Concat", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		out, err := f.Concat(ctx, "samples/concat.txt", "samples/concat.webm")
		if err != nil {
			t.Errorf("an error was thrown: %v", err)
		}

		t.Log(string(out))
	})

	t.Run("Concat/NotExists", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		_, err := f.Concat(ctx, "", "")
		if err == nil {
			t.Errorf("was expecting an error, got nil")
		}

		if !errors.Is(err, main.ErrFfmpegError) {
			t.Errorf("expecting an error of %v, got %v instead", main.ErrFfmpegError, err)
		}
	})
}
