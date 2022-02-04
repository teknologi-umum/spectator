package file

import (
	"context"
	"fmt"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type ExamEvent struct {
	Measurement string    `json:"_measurement" csv:"_measurement"`
	SessionId   string    `json:"session_id" csv:"session_id"`
	Timestamp   time.Time `json:"timepstamp" csv:"timestamp"`
}

func (d *Dependency) QueryExamEnded(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]ExamEvent, error) {
	return d.queryExamEvents(ctx, queryAPI, sessionID, common.MeasurementExamEnded)
}

func (d *Dependency) QueryExamForfeited(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]ExamEvent, error) {
	return d.queryExamEvents(ctx, queryAPI, sessionID, common.MeasurementExamForfeited)
}

func (d *Dependency) QueryExamIDEReloaded(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]ExamEvent, error) {
	return d.queryExamEvents(ctx, queryAPI, sessionID, common.MeasurementExamIDEReloaded)
}

func (d *Dependency) queryExamEvents(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID, measurement string) ([]ExamEvent, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementExamEnded+`" and r["session_id"] == "`+sessionID.String()+`")`,
	)
	if err != nil {
		return []ExamEvent{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	var outputExamEvents []ExamEvent

	for afterExamSamRows.Next() {
		record := afterExamSamRows.Record()

		outputExamEvents = append(outputExamEvents, ExamEvent{
			Measurement: common.MeasurementExamEnded,
			SessionId:   sessionID.String(),
			Timestamp:   record.Time(),
		})
	}

	return outputExamEvents, nil
}
