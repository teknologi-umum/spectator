package file

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: after_exam_sam_submitted
type AfterExamSAMSubmitted struct {
	SessionId    string // Tag
	ArousedLevel uint32
	PleasedLevel uint32
	Timestamp    time.Time
}

func (d *Dependency) QueryAfterExamSam(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]AfterExamSAMSubmitted, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		ReinaldysBuildQuery(Queries{
			Level:     "after_exam_sam_submitted",
			SessionID: sessionID.String(),
			Buckets:   d.BucketSessionEvents,
		}),
	)
	if err != nil {
		return []AfterExamSAMSubmitted{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	//var lastTableIndex int = -1
	outputAfterExam := []AfterExamSAMSubmitted{}
	tempAfterExam := AfterExamSAMSubmitted{}
	var tablePosition int64
	for afterExamSamRows.Next() {
		rows := afterExamSamRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		switch rows.Field() {
		case "aroused_level":
			tempAfterExam.ArousedLevel, ok = rows.Value().(uint32)
			if !ok {
				tempAfterExam.ArousedLevel = 0
			}
		case "pleased_level":
			tempAfterExam.PleasedLevel, ok = rows.Value().(uint32)
			if !ok {
				tempAfterExam.PleasedLevel = 0
			}
		}

		if d.IsDebug() {
			log.Println(rows.String())
			log.Printf("table %d\n", rows.Table())
		}

		if table != 0 && table > tablePosition {
			outputAfterExam = append(outputAfterExam, tempAfterExam)
			tablePosition = table
		} else {
			var ok bool

			tempAfterExam.SessionId, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempAfterExam.SessionId = ""
			}
			tempAfterExam.Timestamp = rows.Time()
		}
	}

	if len(outputAfterExam) > 0 || tempAfterExam.SessionId != "" {
		outputAfterExam = append(outputAfterExam, tempAfterExam)
	}

	return outputAfterExam, nil

}
