package file

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: exam_forfeited
type ExamForfeited struct {
	SessionId string // tag
	Timestamp time.Time
}

func (d *Dependency) QueryExamForfeited(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]ExamForfeited, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		ReinaldysBuildQuery(Queries{
			Level:     "after_exam_sam_submitted",
			SessionID: sessionID.String(),
			Buckets:   d.BucketSessionEvents,
		}),
	)
	if err != nil {
		return []ExamForfeited{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	//var lastTableIndex int = -1
	outputExamForfeited := []ExamForfeited{}
	tempExamForfeited := ExamForfeited{}
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
			outputExamForfeited = append(outputExamForfeited, tempExamForfeited)
			tablePosition = table
		} else {
			var ok bool

			tempExamForfeited.SessionId, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempExamForfeited.SessionId = ""
			}
			tempExamForfeited.Timestamp = rows.Time()
		}
	}

	if len(outputExamForfeited) > 0 || tempExamForfeited.SessionId != "" {
		outputExamForfeited = append(outputExamForfeited, tempExamForfeited)
	}

	return outputExamForfeited, nil

}
