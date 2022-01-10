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

	if len(p.GetData()) == 0 {
		return fmt.Errorf("proper data must be provided")
	}

	if len(p.GetData()) > 5 {
		return fmt.Errorf("log data is more than five, maximum of five")
	}

	var missing []string
	for _, field := range p.GetData() {
		if field.GetRequestId() == "" || strings.Contains(field.GetRequestId(), ",") {
			missing = append(missing, "request_id")
		}

		if field.GetApplication() == "" || strings.Contains(field.GetApplication(), ",") {
			missing = append(missing, "application")
		}

		if field.GetMessage() == "" {
			missing = append(missing, "message")
		}
	}

	if len(missing) == 0 {
		return nil
	}

	return fmt.Errorf("proper %s must be provided", strings.Join(missing, ", "))
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

func (d *Dependency) writeIntoLog(ctx context.Context, p []LogData) error {
	// recovering here in case of any error
	defer func() {
		r := recover()
		if r != nil {
			log.Printf("panic: %v", r.(error))
		}
	}()

	writeAPI := d.DB.WriteAPI(d.Org, "log")

	for index, payload := range p {
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
			return fmt.Errorf("payload index: %d, marshalling json: %v", index, err)
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

		writeAPI.WritePoint(point)
	}

	writeAPI.Flush()
	return nil
}
