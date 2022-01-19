package main

import (
	"context"
	"fmt"

	logger "worker/logger_proto"
	pb "worker/worker_proto"

	"github.com/google/uuid"
)

// GenerateFile is the handler for generating file into CSV and JSON based on
// the input data (which only contains the Session ID).
func (d *Dependency) GenerateFiles(ctx context.Context, in *pb.Member) (*pb.EmptyResponse, error) {
	sessionID, err := uuid.Parse(in.GetSessionId())
	if err != nil {
		defer d.Logger.Log(
			err.Error(),
			logger.Level_ERROR.Enum(),
			in.RequestId,
			map[string]string{
				"session_id": in.SessionId,
				"function":   "GenerateFiles",
				"info":       "parsing uuid",
			},
		)
		return &pb.EmptyResponse{}, fmt.Errorf("parsing uuid: %v", err)
	}

	go d.File.CreateFile(in.RequestId, sessionID)

	return &pb.EmptyResponse{}, nil
}
