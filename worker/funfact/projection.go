package funfact

import (
	"context"

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

	p := influxdb2.NewPointWithMeasurement("funfact_projection")
	p.AddTag("session_id", sessionID.String())
	p.AddTag("student_number", studentNumber)
	p.AddField("words_per_minute", wpm)
	p.AddField("deletion_rate", deletionRate)
	p.AddField("submission_attemps", attempts)

	// FIXME: the bucket name should be input_statistics
	// TODO: check if the bucket exists first, then create if not exists
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
		return
	}
}
