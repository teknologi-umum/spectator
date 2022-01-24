package file

import (
	"context"
	"fmt"
	"log"
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
		ReinaldysBuildQuery(Queries{
			Level:     "after_exam_sam_submitted",
			SessionID: sessionID.String(),
			Buckets:   d.BucketSessionEvents,
		}),
	)
	if err != nil {
		return []ExamEnded{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	//var lastTableIndex int = -1
	outputExamEnded := []ExamEnded{}
	tempExamEnded := ExamEnded{}
	var tablePosition int64
	for afterExamSamRows.Next() {
		rows := afterExamSamRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		if d.IsDebug() {
			log.Println(rows.String())
			log.Printf("table %d\n", rows.Table())
		}

		if table != 0 && table > tablePosition {
			outputExamEnded = append(outputExamEnded, tempExamEnded)
			tablePosition = table
		} else {
			var ok bool

			tempExamEnded.SessionId, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempExamEnded.SessionId = ""
			}
			tempExamEnded.Timestamp = rows.Time()
		}
	}

	if len(outputExamEnded) > 0 || tempExamEnded.SessionId != "" {
		outputExamEnded = append(outputExamEnded, tempExamEnded)
	}

	return outputExamEnded, nil

}
