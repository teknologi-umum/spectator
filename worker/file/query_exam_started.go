package file

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: exam_started
type ExamStarted struct {
	SessionId       string    // tag
	QuestionNumbers string    // field
	Deadline        time.Time // field
	Timestamp       time.Time
}

func (d *Dependency) QueryExamStarted(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]ExamStarted, error) {
	examStartedRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "exam_started" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return []ExamStarted{}, fmt.Errorf("failed to query exam_started: %w", err)
	}
	defer examStartedRows.Close()

	var outputExamStarted []ExamStarted

	for examStartedRows.Next() {
		rows := examStartedRows.Record()

		questionNumbers, ok := rows.ValueByKey("question_numbers").(string)
		if !ok {
			questionNumbers = ""
		}
		deadline, ok := rows.ValueByKey("deadline").(time.Time)
		if !ok {
			deadline = time.Time{}
		}

		outputExamStarted = append(outputExamStarted, ExamStarted{
			SessionId:       sessionID.String(),
			QuestionNumbers: questionNumbers,
			Deadline:        deadline,
			Timestamp:       time.Unix(0, 0),
		})
	}

	return outputExamStarted, nil
}
