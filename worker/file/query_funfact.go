package file

import (
	"context"
	"fmt"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Funfact struct {
	SessionId          string    `json:"session_id" csv:"session_id"` // tag
	WordsPerMinute     int64     `json:"words_per_minute" csv:"words_per_minute"`
	DeletionRate       int64     `json:"deletion_rate" csv:"deletion_rate"`
	SubmissionAttempts int64     `json:"submission_attempts" csv:"submission_attempts"`
	Timestamp          time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryFunfact(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (Funfact, error) {
	funfactRows, err := queryAPI.Query(
		ctx,
		`from(bucket:"`+common.BucketInputStatisticEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementFunfactProjection+`" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey:["_time"], columnKey:["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return Funfact{}, fmt.Errorf("failed to query funfact: %w", err)
	}
	defer funfactRows.Close()

	var outputFunfact Funfact

	for funfactRows.Next() {
		record := funfactRows.Record()
		wordsPerMinute, ok := record.ValueByKey("words_per_minute").(int64)
		if !ok {
			wordsPerMinute = 0
		}

		deletionRate, ok := record.ValueByKey("deletion_rate").(int64)
		if !ok {
			deletionRate = 0
		}

		submissionAttempts, ok := record.ValueByKey("submission_attempts").(int64)
		if !ok {
			submissionAttempts = 0
		}

		outputFunfact = Funfact{
			SessionId:          sessionID.String(),
			WordsPerMinute:     wordsPerMinute,
			DeletionRate:       deletionRate,
			SubmissionAttempts: submissionAttempts,
			Timestamp:          record.Time(),
		}
	}

	return outputFunfact, nil
}
