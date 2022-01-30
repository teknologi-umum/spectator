package file

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: exam_forfeited
type ExamForfeited struct {
	SessionId string    `json:"session_id" csv:"session_id"` // tag
	Timestamp time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryExamForfeited(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]ExamForfeited, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "exam_forfeited" and r["session_id"] == "`+sessionID.String()+`")`,
	)
	if err != nil {
		return []ExamForfeited{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	var outputExamForfeited []ExamForfeited
	for afterExamSamRows.Next() {
		rows := afterExamSamRows.Record()

		outputExamForfeited = append(outputExamForfeited, ExamForfeited{
			SessionId: sessionID.String(),
			Timestamp: rows.Time(),
		})
	}

	return outputExamForfeited, nil
}
