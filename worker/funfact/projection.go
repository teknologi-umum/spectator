package funfact

import (
	"context"
	"time"

	"worker/influxhelpers"
	loggerpb "worker/logger_proto"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func (d *Dependency) CreateProjection(ctx context.Context, sessionID uuid.UUID, wpm uint32, attempts uint32, deletionRate float32, requestID string) {
	personalInfoh, err := d.DB.QueryAPI(d.BucketSessionEvents).Query(
		ctx,
		influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
			Measurement: "personal_info",
			SessionID:   sessionID.String(),
			Buckets:     d.BucketSessionEvents,
		}),
	)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "Create Projection",
				"info":       "cannot proceed student number",
			},
		)
		return
	}

	var studentNumber string
	for personalInfoh.Next() {
		rows := personalInfoh.Record()
		switch rows.Field() {
		case "student_number":
			var ok bool
			studentNumber, ok = rows.Value().(string)
			if !ok {
				studentNumber = ""
			}
		}
	}

	point := influxdb2.NewPoint(
		"funfact_projection",
		map[string]string{
			"session_id":     sessionID.String(),
			"student_number": studentNumber,
		},
		map[string]interface{}{
			"words_per_minute":    wpm,
			"deletion_rate":       deletionRate,
			"submission_attempts": attempts,
		},
		time.Now(),
	)

	err = d.DB.
		WriteAPIBlocking(d.DBOrganization, d.BucketExamResult).
		WritePoint(ctx, point)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "Create Projection",
				"info":       "cannot storing results",
			},
		)
		return
	}
}
