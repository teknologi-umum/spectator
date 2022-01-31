package main

import (
	"context"
	"fmt"
	logger "worker/logger_proto"
	pb "worker/worker_proto"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// FunFact is the handler for generating fun fact about the user
// after they had done their coding test.
func (d *Dependency) FunFact(ctx context.Context, in *pb.Member) (*pb.FunFactResponse, error) {
	// Parse UUID
	sessionID, err := uuid.Parse(in.GetSessionId())
	if err != nil {
		defer d.Logger.Log(
			err.Error(),
			logger.Level_ERROR.Enum(),
			in.RequestId,
			map[string]string{
				"session_id": in.GetSessionId(),
				"function":   "funfact",
				"info":       "parsing uuid",
			},
		)
		return &pb.FunFactResponse{}, fmt.Errorf("parsing uuid: %v", err)
	}

	// Read about buffered channel vs non-buffered channels
	wpm := make(chan uint32, 1)
	deletionRate := make(chan float32, 1)
	attempt := make(chan uint32, 1)

	// Run all the calculate function concurently
	errs, gctx := errgroup.WithContext(ctx)
	errs.Go(func() error {
		return d.Funfact.CalculateWordsPerMinute(gctx, sessionID, wpm)
	})
	errs.Go(func() error {
		return d.Funfact.CalculateDeletionRate(gctx, sessionID, deletionRate)
	})
	errs.Go(func() error {
		return d.Funfact.CalculateSubmissionAttempts(gctx, sessionID, attempt)
	})

	err = errs.Wait()
	if err != nil {
		defer d.Logger.Log(
			err.Error(),
			logger.Level_ERROR.Enum(),
			in.RequestId,
			map[string]string{
				"session_id": in.GetSessionId(),
				"function":   "funfact",
				"info":       "calculating fun fact",
			},
		)
		return &pb.FunFactResponse{}, fmt.Errorf("calculating fun fact: %v", err)
	}

	var result = struct {
		Wpm          uint32  `json:"wpm"`
		DeletionRate float32 `json:"deletion_rate"`
		Attempt      uint32  `json:"attempt"`
	}{
		<-wpm,
		<-deletionRate,
		<-attempt,
	}

	defer d.Funfact.CreateProjection(sessionID, result.Wpm, result.Attempt, result.DeletionRate, in.RequestId)

	return &pb.FunFactResponse{
		WordsPerMinute:     result.Wpm,
		DeletionRate:       result.DeletionRate,
		SubmissionAttempts: result.Attempt,
	}, nil
}
