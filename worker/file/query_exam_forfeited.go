package file

import (
	"context"
	"fmt"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: exam_forfeited
type ExamForfeited struct {
	Measurement string    `json:"_measurement" csv:"_measurement"` // tag
	SessionId   string    `json:"session_id" csv:"session_id"`     // tag
	Timestamp   time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryExamForfeited(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]ExamForfeited, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementExamForfeited+`" and r["session_id"] == "`+sessionID.String()+`")`,
	)
	if err != nil {
		return []ExamForfeited{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	var outputExamForfeited []ExamForfeited
	for afterExamSamRows.Next() {
		record := afterExamSamRows.Record()

		outputExamForfeited = append(outputExamForfeited, ExamForfeited{
			Measurement: common.MeasurementExamForfeited,
			SessionId:   sessionID.String(),
			Timestamp:   record.Time(),
		})
	}

	return outputExamForfeited, nil
}
