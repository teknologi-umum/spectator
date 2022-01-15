package main

import (
	"context"
	"fmt"
	pb "logger/proto"
)

func (d *Dependency) Ping(ctx context.Context, _ *pb.EmptyRequest) (*pb.Healthcheck, error) {
	health, err := d.DB.Health(ctx)
	if err != nil {
		return &pb.Healthcheck{}, fmt.Errorf("health check call: %v", err)
	}

	return &pb.Healthcheck{
		Status: string(health.Status),
	}, nil
}
