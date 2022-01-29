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

	// We map the files returned by the ListFiles function into Protobuf format
	var files []*pb.File
	for _, file := range result {
		files = append(files, &pb.File{
			SessionId: file.SessionId,
			StudentNumber: file.StudentNumber,
			FileUrlJson: file.JSONFile,
			FileUrlCsv: file.CSVFile,
		})
	}

	return &pb.FilesList{Files: files}, nil
}
