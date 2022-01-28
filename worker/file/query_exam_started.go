package file

import (
	"context"
	"fmt"
	"log"
	"time"
	"worker/influxhelpers"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: exam_started
type ExamStarted struct {
	SessionId       string    // tag
	QuestionNumbers uint32    // field
	Deadline        time.Time // field
	Timestamp       time.Time
}

func (d *Dependency) QueryExamStarted(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]ExamStarted, error) {
	outputExamStarted := []ExamStarted{}
	for _, x := range []string{"question_numbers", "deadline"} {
		afterExamSamRows, err := queryAPI.Query(
			ctx,
			influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
				Measurement: "exam_ide_started",
				SessionID:   sessionID.String(),
				Buckets:     d.BucketSessionEvents,
				Field:       x,
			}),
		)
		if err != nil {
			return []ExamStarted{}, fmt.Errorf("failed to query keystrokes: %w", err)
		}

		//var lastTableIndex int = -1
		tempExamStarted := ExamStarted{}
		var tablePosition int64
		for afterExamSamRows.Next() {
			rows := afterExamSamRows.Record()
			table, ok := rows.ValueByKey("table").(int64)
			if !ok {
				table = 0
			}

			switch x {
			case "question_number":
				tempExamStarted.QuestionNumbers, ok = rows.Value().(uint32)
				if !ok {
					tempExamStarted.QuestionNumbers = 0
				}
			case "deadline":
				tempExamStarted.Deadline, ok = rows.Value().(time.Time)
				if !ok {
					tempExamStarted.Deadline = time.Unix(0, 0)
				}
			}

			if d.IsDebug() {
				log.Println(rows.String())
				log.Printf("table %d\n", rows.Table())
			}

			if table != 0 && table > tablePosition {
				outputExamStarted = append(outputExamStarted, tempExamStarted)
				tablePosition = table
			} else {
				var ok bool

				tempExamStarted.SessionId, ok = rows.ValueByKey("session_id").(string)
				if !ok {
					tempExamStarted.SessionId = ""
				}
				tempExamStarted.Timestamp = rows.Time()
			}
		}

		if len(outputExamStarted) > 0 || tempExamStarted.SessionId != "" {
			outputExamStarted = append(outputExamStarted, tempExamStarted)
		}

	}
	return outputExamStarted, nil

}