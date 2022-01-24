package funfact

import (
	"context"

	"worker/file"

	"github.com/google/uuid"

	loggerpb "worker/logger_proto"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func (d *Dependency) CreateProjection(ctx context.Context, sessionID uuid.UUID, wpm uint32, attempts uint32, deletionRate float32, requestID string) {
	personalInfoh, err := d.DB.QueryAPI(d.BucketSessionEvents).Query(
		ctx,
		file.ReinaldysBuildQuery(file.Queries{
			Level:     "personal_info",
			SessionID: sessionID.String(),
			Buckets:   d.BucketSessionEvents,
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

	p := influxdb2.NewPointWithMeasurement("funfact_projection")
	p.AddTag("session_id", sessionID.String())
	p.AddTag("student_number", studentNumber)
	p.AddField("words_per_minute", wpm)
	p.AddField("deletion_rate", deletionRate)
	p.AddField("submission_attemps", attempts)

	err = d.DB.WriteAPIBlocking(d.DBOrganization, d.BucketResultEvents).WritePoint(ctx, p)
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
	}

	return
}
