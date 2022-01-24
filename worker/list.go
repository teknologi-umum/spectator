package main

import (
	"context"
	"fmt"
	logger "worker/logger_proto"
	pb "worker/worker_proto"

	"github.com/google/uuid"
)

func (d *Dependency) ListFiles(ctx context.Context, in *pb.Member) (*pb.FilesList, error) {
	sessionID, err := uuid.Parse(in.GetSessionId())
	if err != nil {
		defer d.Logger.Log(
			err.Error(),
			logger.Level_ERROR.Enum(),
			in.RequestId,
			map[string]string{
				"session_id": in.GetSessionId(),
				"function":   "list file",
				"info":       "parsing uuid",
			},
		)
		return &pb.FilesList{}, fmt.Errorf("parsing uuid: %v", err)
	}

	result, err := d.File.ListFiles(ctx, sessionID)
	if err != nil {
		defer d.Logger.Log(
			err.Error(),
			logger.Level_ERROR.Enum(),
			in.RequestId,
			map[string]string{
				"session_id": in.GetSessionId(),
				"function":   "list file",
				"info":       "listing file",
			},
		)
		return &pb.FilesList{}, fmt.Errorf("listing file: %v", err)
	}

	return &pb.FilesList{Files: result}, nil
}
