package main

import (
	"context"
	logger "worker/logger_proto"
	pb "worker/worker_proto"
)

func (d *Dependency) Ping(ctx context.Context, in *pb.EmptyRequest) (*pb.Health, error) {
	health, err := d.DB.Health(ctx)
	if err != nil {
		defer d.Logger.Log(err.Error(), logger.Level_ERROR.Enum(), "", map[string]string{})
		return &pb.Health{}, err
	}

	return &pb.Health{
		Status: string(health.Status),
	}, nil
}
