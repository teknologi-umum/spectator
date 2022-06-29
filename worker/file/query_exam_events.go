package file

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"worker/common"
	"worker/logger_proto"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type ExamEvent struct {
	Measurement string    `json:"_measurement" csv:"_measurement"`
	SessionId   string    `json:"session_id" csv:"session_id"`
	Timestamp   time.Time `json:"timepstamp" csv:"timestamp"`
}

func (d *Dependency) QueryExamEnded(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (*ExamEvent, error) {
	events, err := d.queryExamEvents(ctx, queryAPI, sessionID, common.MeasurementExamEnded)
	if err != nil {
		return &ExamEvent{}, err
	}

	if len(*events) == 0 {
		return &ExamEvent{}, nil
	}

	return &(*events)[0], nil
}

func (d *Dependency) QueryExamForfeited(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (*ExamEvent, error) {
	events, err := d.queryExamEvents(ctx, queryAPI, sessionID, common.MeasurementExamForfeited)
	if err != nil {
		return &ExamEvent{}, err
	}

	if len(*events) == 0 {
		return &ExamEvent{}, nil
	}

	return &(*events)[0], nil
}

func (d *Dependency) QueryExamIDEReloaded(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (*[]ExamEvent, error) {
	return d.queryExamEvents(ctx, queryAPI, sessionID, common.MeasurementExamIDEReloaded)
}

func (d *Dependency) queryExamEvents(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID, measurement string) (*[]ExamEvent, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+measurement+`" and r["session_id"] == "`+sessionID.String()+`")`,
	)
	if err != nil {
		return &[]ExamEvent{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	var outputExamEvents []ExamEvent

	for afterExamSamRows.Next() {
		record := afterExamSamRows.Record()

		if record.Time().Year() != 2022 {
			d.Logger.Log(
				record.Time().String(),
				logger_proto.Level_DEBUG.Enum(),
				sessionID.String(),
				map[string]string{
					"session_id": sessionID.String(),
					"function":   "queryExamEvents",
				},
			)
			log.Printf("current time from record.Time() is not 2022, it's " + strconv.Itoa(record.Time().Year()))
		}

		outputExamEvents = append(outputExamEvents, ExamEvent{
			Measurement: measurement,
			SessionId:   sessionID.String(),
			Timestamp:   record.Time(),
		})
	}

	return &outputExamEvents, nil
}
