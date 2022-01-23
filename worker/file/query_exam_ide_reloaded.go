package file

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: exam_ide_reloaded
type ExamIDEReloaded struct {
	SessionId string // tag
	Timestamp time.Time
}

func (d *Dependency) QueryExamIDEReloaded(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]ExamIDEReloaded, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		ReinaldysBuildQuery(Queries{
			Level:     "after_exam_sam_submitted",
			SessionID: sessionID.String(),
			Buckets:   d.BucketSessionEvents,
		}),
	)
	if err != nil {
		return []ExamIDEReloaded{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	//var lastTableIndex int = -1
	outputExamIDEReloaded := []ExamIDEReloaded{}
	tempExamIDEReloaded := ExamIDEReloaded{}
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
			outputExamIDEReloaded = append(outputExamIDEReloaded, tempExamIDEReloaded)
			tablePosition = table
		} else {
			var ok bool

			tempExamIDEReloaded.SessionId, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempExamIDEReloaded.SessionId = ""
			}
			tempExamIDEReloaded.Timestamp = rows.Time()
		}
	}

	if len(outputExamIDEReloaded) > 0 || tempExamIDEReloaded.SessionId != "" {
		outputExamIDEReloaded = append(outputExamIDEReloaded, tempExamIDEReloaded)
	}

	return outputExamIDEReloaded, nil

}
