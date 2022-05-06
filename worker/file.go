package main

import (
	"context"

	logger "worker/logger_proto"
	pb "worker/worker_proto"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return &pb.EmptyResponse{}, status.Errorf(codes.InvalidArgument, "parsing uuid: %vw", err)
	}

	exists, err := d.File.CheckIfSessionExists(ctx, sessionID)
	if err != nil {
		defer d.Logger.Log(
			err.Error(),
			logger.Level_ERROR.Enum(),
			in.RequestId,
			map[string]string{
				"session_id": in.SessionId,
				"function":   "GenerateFiles",
				"info":       "checking if session exists",
			},
		)
		return &pb.EmptyResponse{}, status.Errorf(codes.Internal, "checking if session exists: %v", err)
	}

	if !exists {
		return &pb.EmptyResponse{}, status.Error(codes.NotFound, "session not found")
	}

	go d.File.CreateFile(in.RequestId, sessionID)

	return &pb.EmptyResponse{}, nil
}
