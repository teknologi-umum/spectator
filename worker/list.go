package main

import (
	"context"
	pb "worker/worker_proto"
)

func (d *Dependency) ListFiles(context.Context, *pb.Member) (*pb.FilesList, error) {
	// TODO: airavata to the rescue
	return &pb.FilesList{}, nil
}
