package file

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: exam_ended
type ExamEnded struct {
	SessionId string    `json:"session_id" csv:"session_id"` // tag
	Timestamp time.Time `json:"timepstamp" csv:"timestamp"`
}

func (d *Dependency) QueryExamEnded(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]ExamEnded, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "exam_ended" and r["session_id"] == `+sessionID.String()+`)`,
	)
	if err != nil {
		return []ExamEnded{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	var outputExamEnded []ExamEnded

	for afterExamSamRows.Next() {
		rows := afterExamSamRows.Record()

		outputExamEnded = append(outputExamEnded, ExamEnded{
			SessionId: sessionID.String(),
			Timestamp: rows.Time(),
		})
	}

	return outputExamEnded, nil
}
