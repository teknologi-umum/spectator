package file

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: exam_ide_reloaded
type ExamIDEReloaded struct {
	SessionId string    `json:"session_id" csv:"session_id"` // tag
	Timestamp time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryExamIDEReloaded(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]ExamIDEReloaded, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "exam_ide_reloaded" and r["session_id"] == "`+sessionID.String()+`")`,
	)
	if err != nil {
		return []ExamIDEReloaded{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	var outputExamIDEReloaded []ExamIDEReloaded

	for afterExamSamRows.Next() {
		record := afterExamSamRows.Record()

		outputExamIDEReloaded = append(outputExamIDEReloaded, ExamIDEReloaded{
			SessionId: sessionID.String(),
			Timestamp: record.Time(),
		})
	}

	return outputExamIDEReloaded, nil
}
