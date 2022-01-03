package main

import (
	"context"
	pb "worker/proto"
)

func (d *Dependency) Ping(ctx context.Context, in *pb.EmptyRequest) (*pb.Health, error) {
	health, err := d.DB.Health(ctx)
	if err != nil {
		return &pb.Health{}, err
	}

	return &pb.Health{
		Status: string(health.Status),
	}, nil
}
