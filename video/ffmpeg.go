package main

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Ffmpeg struct {
	valid bool
}

var ErrFfmpegInvalid = errors.New("invalid validation for ffmpeg")
var ErrFfmpegError = errors.New("ffmpeg command error")

func NewFfmpeg() (*Ffmpeg, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	f := &Ffmpeg{}
	return f, f.validateVersion(ctx)
}

func (f *Ffmpeg) validateVersion(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "ffmpeg", "-version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("running version check: %w", err)
	}

	if strings.Contains(string(out), "ffmpeg version") {
		f.valid = true
		return nil
	}

	return ErrFfmpegInvalid
}

func (f *Ffmpeg) Concat(ctx context.Context, src string, dst string) ([]byte, error) {
	if !f.valid {
		return []byte{}, ErrFfmpegInvalid
	}

	cmd := exec.CommandContext(ctx, "ffmpeg", "-y", "-f", "concat", "-i", src, "-c", "copy", dst)
	stdout, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return []byte{}, fmt.Errorf("%w: %s", ErrFfmpegError, string(exitErr.Stderr))
		}

		return []byte{}, fmt.Errorf("executing concat: %w", err)
	}

	return stdout, nil
}

func (f *Ffmpeg) Convert(ctx context.Context, src string, dst string) ([]byte, error) {
	if !f.valid {
		return []byte{}, ErrFfmpegInvalid
	}

	cmd := exec.CommandContext(ctx, "ffmpeg", "-y", "-i", src, "-crf", "3", "-c", "copy", dst)
	stdout, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return []byte{}, fmt.Errorf("%w: %s", ErrFfmpegError, string(exitErr.Stderr))
		}

		return []byte{}, fmt.Errorf("executing concat: %w", err)
	}

	return stdout, nil
}
