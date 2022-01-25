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

// measurement: deadline_passed
type DeadlinePassed struct {
	SessionId string    `json:"session_id" csv:"session_id"` // tag
	Timestamp time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryDeadlinePassed(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]DeadlinePassed, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
			Measurement: "deadline_passed",
			SessionID:   sessionID.String(),
			Buckets:     d.BucketSessionEvents,
		}),
	)
	if err != nil {
		return []DeadlinePassed{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	//var lastTableIndex int = -1
	outputDeadlinePassed := []DeadlinePassed{}
	tempDeadlinePassed := DeadlinePassed{}
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
			outputDeadlinePassed = append(outputDeadlinePassed, tempDeadlinePassed)
			tablePosition = table
		} else {
			var ok bool

			tempDeadlinePassed.SessionId, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempDeadlinePassed.SessionId = ""
			}
			tempDeadlinePassed.Timestamp = rows.Time()
		}
	}

	if len(outputDeadlinePassed) > 0 || tempDeadlinePassed.SessionId != "" {
		outputDeadlinePassed = append(outputDeadlinePassed, tempDeadlinePassed)
	}

	return outputDeadlinePassed, nil

}
