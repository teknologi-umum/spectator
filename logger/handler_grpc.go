package main

import (
	"context"
	"fmt"
	pb "logger/proto"
	"strings"
	"time"
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

func (d *Dependency) ValidatePayload(p *pb.LogRequest) error {
	if p.GetAccessToken() != d.AccessToken {
		return fmt.Errorf("access token must be provided")
	}

	var missing []string
	for _, field := range p.GetData() {
		if field.RequestId == "" || strings.Contains(field.RequestId, ",") {
			missing = append(missing, "request_id")
		}

		if field.Application == "" || strings.Contains(field.Application, ",") {
			missing = append(missing, "application")
		}

		if field.Message == "" {
			missing = append(missing, "message")
		}
	}

	if len(missing) == 0 {
		return nil
	}

	return fmt.Errorf("proper %s must be provided", strings.Join(missing, ", "))
}

func (*Dependency) convertIntoLogData(l []*pb.LogData) []LogData {
	var data []LogData
	for _, d := range l {
		data = append(
			data,
			LogData{
				RequestID:   d.GetRequestId(),
				Application: d.GetApplication(),
				Message:     d.GetMessage(),
				Body:        d.GetBody(),
				Level:       d.GetLevel().String(),
				Environment: d.GetEnvironment().String(),
				Language:    d.GetLanguage(),
				Timestamp:   time.UnixMilli(d.GetTimestamp()),
			},
		)
	}

	return data
}

func (*Dependency) convertIntoProtoData(l []LogData) []*pb.LogData {
	var data []*pb.LogData
	for _, d := range l {
		var level pb.Level
		var environment pb.Environment
		var timestamp = d.Timestamp.UnixMilli()

		switch d.Level {
		case "INFO":
			level = pb.Level_INFO
		case "WARNING":
			level = pb.Level_WARNING
		case "ERROR":
			level = pb.Level_ERROR
		case "CRITICAL":
			level = pb.Level_CRITICAL
		default:
			level = pb.Level_DEBUG
		}

		switch d.Environment {
		case "PRODUCTION":
			environment = pb.Environment_PRODUCTION
		case "STAGING":
			environment = pb.Environment_STAGING
		case "DEVELOPMENT":
			environment = pb.Environment_DEVELOPMENT
		default:
			environment = pb.Environment_UNSET
		}

		data = append(data, &pb.LogData{
			RequestId:   d.RequestID,
			Application: d.Application,
			Message:     d.Message,
			Body:        d.Body,
			Level:       &level,
			Environment: &environment,
			Language:    &d.Language,
			Timestamp:   &timestamp,
		})
	}

	return data
}

func (d *Dependency) CreateLog(ctx context.Context, r *pb.LogRequest) (*pb.EmptyResponse, error) {
	err := d.ValidatePayload(r)
	if err != nil {
		return &pb.EmptyResponse{}, err
	}

	err = d.writeIntoLog(ctx, d.convertIntoLogData(r.GetData()))
	if err != nil {
		return &pb.EmptyResponse{}, fmt.Errorf("writing log: %v", err)
	}

	return &pb.EmptyResponse{}, nil
}

type queries struct {
	Level       string
	RequestID   string
	Application string
	TimeFrom    time.Time
	TimeTo      time.Time
}

func (d *Dependency) ReadLog(ctx context.Context, r *pb.ReadLogRequest) (*pb.ReadLogResponse, error) {
	var query = queries{
		Level:       r.GetLevel().String(),
		RequestID:   r.GetRequestId(),
		Application: r.GetApplication(),
		TimeFrom:    time.UnixMilli(r.GetTimestampFrom()),
		TimeTo:      time.UnixMilli(r.GetTimestampTo()),
	}

	logs, err := d.fetchLog(ctx, query)
	if err != nil {
		return &pb.ReadLogResponse{}, fmt.Errorf("reading log: %v", err)
	}

	return &pb.ReadLogResponse{
		Data: d.convertIntoProtoData(logs),
	}, nil
}
