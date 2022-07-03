package main

import (
	"context"
	pb "video/video_proto"
)

// Ping is the gRPC presentation layer handler for ping commands.
func (d *Dependency) Ping(ctx context.Context, in *pb.EmptyRequest) (*pb.PingResponse, error) {
	health, err := d.DB.Health(ctx)
	if err != nil {
		return &pb.PingResponse{}, err
	}

	return &pb.PingResponse{
		Message: string(health.Status),
	}, nil
}
