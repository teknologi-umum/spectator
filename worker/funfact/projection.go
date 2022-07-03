package funfact

import (
	"context"
	"time"

	"worker/common"
	loggerpb "worker/logger_proto"
	"worker/status"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/rs/zerolog/log"
)

func (d *Dependency) CreateProjection(sessionID uuid.UUID, wpm int64, attempts int64, deletionRate float64, requestID string) {
	// Defer func to avoid panic
	defer func() {
		if r, ok := recover().(error); ok {
			log.Error().
				Str("request_id", requestID).
				Str("session_id", sessionID.String()).
				Err(r).
				Msg("recovered from panic on CreateProjection")

			d.Logger.Log(
				r.(error).Error(),
				loggerpb.Level_CRITICAL.Enum(),
				requestID,
				map[string]string{
					"session_id": sessionID.String(),
					"function":   "CreateProjection",
					"info":       "recovering from panic",
				},
			)
		}
	}()

	log.Info().
		Str("request_id", requestID).
		Str("session_id", sessionID.String()).
		Msg("received create projection request")

	// Create a new context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	d.Status.AppendState(ctx, sessionID, "create_projection", status.StatePending)

	// We shall find the student number
	personalInfoRows, err := d.DB.QueryAPI(d.DBOrganization).Query(
		ctx,
		`from(bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementPersonalInfoSubmitted+`" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"])`,
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
		d.Status.AppendState(ctx, sessionID, "create_projection", status.StateFailed)
		return
	}
	defer personalInfoRows.Close()

	var studentNumber string
	for personalInfoRows.Next() {
		value, ok := personalInfoRows.Record().ValueByKey("student_number").(string)
		if !ok {
			value = ""
		}
		studentNumber = value
	}

	point := influxdb2.NewPoint(
		common.MeasurementFunfactProjection,
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
		WriteAPIBlocking(d.DBOrganization, common.BucketInputStatisticEvents).
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
		d.Status.AppendState(ctx, sessionID, "create_projection", status.StateFailed)
		return
	}

	log.Info().
		Str("session_id", sessionID.String()).
		Str("request_id", requestID).
		Msg("successfully created projection")

	d.Status.AppendState(ctx, sessionID, "create_projection", status.StateSuccess)
}
