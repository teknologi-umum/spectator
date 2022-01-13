package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"log"
	"strings"
	"time"

	pb "logger/proto"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func (d *Dependency) ValidatePayload(p *pb.LogRequest) error {
	if p.GetAccessToken() != d.AccessToken {
		return fmt.Errorf("access token must be provided")
	}

	var missing []string
	if p.GetData().GetRequestId() == "" || strings.Contains(p.GetData().GetRequestId(), ",") {
		missing = append(missing, "request_id")
	}

	if p.GetData().GetApplication() == "" || strings.Contains(p.GetData().GetApplication(), ",") {
		missing = append(missing, "application")
	}

	if p.GetData().GetMessage() == "" {
		missing = append(missing, "message")
	}

	if len(missing) == 0 {
		return nil
	}

	if len(missing) > 0 {
		return fmt.Errorf("proper %s must be provided", strings.Join(missing, ", "))
	}

	return nil
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

func (d *Dependency) writeIntoLog(ctx context.Context, payload LogData) error {
	// recovering here in case of any error
	defer func() {
		r := recover()
		if r != nil {
			log.Printf("panic: %v", r.(error))
		}
	}()

	writeAPI := d.DB.WriteAPIBlocking(d.Org, "log")

	// write defaults first
	if payload.Environment == "" {
		payload.Environment = "UNSET"
	}

	if payload.Level == "" {
		payload.Level = "DEBUG"
	}

	if payload.Timestamp.IsZero() {
		payload.Timestamp = time.Now()
	}

	// convert body into json as influxdb doesnt accept map
	bodyBytes, err := json.Marshal(payload.Body)
	if err != nil {
		return fmt.Errorf("m3arshalling json: %v", err)
	}

	point := influxdb2.NewPoint(
		payload.Level,
		map[string]string{
			"request_id":  payload.RequestID,
			"application": payload.Application,
			"environment": payload.Environment,
		},
		map[string]interface{}{
			"language": payload.Language,
			"message":  hex.EncodeToString([]byte(payload.Message)),
			"body":     hex.EncodeToString(bodyBytes),
		},
		payload.Timestamp,
	)

	err = writeAPI.WritePoint(ctx, point)
	if err != nil {
		return fmt.Errorf("writing point: %v", err)
	}
	return nil
}
