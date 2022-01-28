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

// measurement: before_exam_sam_submitted
type BeforeExamSAMSubmitted struct {
	SessionId    string    `json:"session_id" csv:"session_id"` // Tag
	ArousedLevel uint32    `json:"aroused_level" csv:"aroused_level"`
	PleasedLevel uint32    `json:"pleased_level" csv:"pleased_level"`
	Timestamp    time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryBeforeExamSam(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]BeforeExamSAMSubmitted, error) {
	outputBeforeExam := []BeforeExamSAMSubmitted{}
	for _, field := range []string{"aroused_level", "pleased_level"} {
		afterExamSamRows, err := queryAPI.Query(
			ctx,
			influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
				Measurement: "before_exam_sam_submitted",
				SessionID:   sessionID.String(),
				Buckets:     d.BucketSessionEvents,
				Field:       field,
			}),
		)
		if err != nil {
			return []BeforeExamSAMSubmitted{}, fmt.Errorf("failed to query keystrokes: %w", err)
		}

		//var lastTableIndex int = -1
		tempAfterExam := BeforeExamSAMSubmitted{}
		var tablePosition int64
		for afterExamSamRows.Next() {
			rows := afterExamSamRows.Record()
			table, ok := rows.ValueByKey("table").(int64)
			if !ok {
				table = 0
			}

			switch field {
			case "aroused_level":
				v, ok := rows.Value().(int64)
				if !ok {
					v = 0
				}

				tempAfterExam.ArousedLevel = uint32(v)
			case "pleased_level":
				v, ok := rows.Value().(int64)
				if !ok {
					v = 0
				}

				tempAfterExam.PleasedLevel = uint32(v)
			}

			if d.IsDebug() {
				log.Println(rows.String())
				log.Printf("table %d\n", rows.Table())
			}

			if table != 0 && table > tablePosition {
				outputBeforeExam = append(outputBeforeExam, tempAfterExam)
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

		if len(outputBeforeExam) > 0 || tempAfterExam.SessionId != "" {
			outputBeforeExam = append(outputBeforeExam, tempAfterExam)
		}
	}

	return outputBeforeExam, nil

}