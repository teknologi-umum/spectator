package file

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: deadline_passed
type DeadlinePassed struct {
	SessionId string    `json:"session_id" csv:"session_id"` // tag
	Timestamp time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryDeadlinePassed(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]DeadlinePassed, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "deadline_passed" and r["session_id"] == `+sessionID.String()+`)`,
	)
	if err != nil {
		return []DeadlinePassed{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	var outputDeadlinePassed []DeadlinePassed

	for afterExamSamRows.Next() {
		rows := afterExamSamRows.Record()

		outputDeadlinePassed = append(outputDeadlinePassed, DeadlinePassed{
			SessionId: sessionID.String(),
			Timestamp: rows.Time(),
		})
	}

	return outputDeadlinePassed, nil
}
