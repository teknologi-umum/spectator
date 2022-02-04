package file

import (
	"context"
	"fmt"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: after_exam_sam_submitted
type AfterExamSAMSubmitted struct {
	Measurement  string    `json:"_measurement" csv:"_measurement"`
	SessionId    string    `json:"session_id" csv:"session_id"`
	ArousedLevel uint32    `json:"aroused_level" csv:"aroused_level"`
	PleasedLevel uint32    `json:"pleased_level" csv:"pleased_level"`
	Timestamp    time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryAfterExamSam(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]AfterExamSAMSubmitted, error) {
	afterExamSamRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementAfterExamSAMSubmitted+`" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return []AfterExamSAMSubmitted{}, fmt.Errorf("failed to query after_exam_sam_submitted: %w", err)
	}
	defer afterExamSamRows.Close()

	var outputAfterExam []AfterExamSAMSubmitted

	for afterExamSamRows.Next() {
		record := afterExamSamRows.Record()

		arousedLevel, ok := record.ValueByKey("aroused_level").(int64)
		if !ok {
			arousedLevel = 0
		}

		pleasedLevel, ok := record.ValueByKey("pleased_level").(int64)
		if !ok {
			pleasedLevel = 0
		}

		sessionId, ok := record.ValueByKey("session_id").(string)
		if !ok {
			sessionId = ""
		}

		outputAfterExam = append(
			outputAfterExam,
			AfterExamSAMSubmitted{
				Measurement:  common.MeasurementAfterExamSAMSubmitted,
				SessionId:    sessionId,
				ArousedLevel: uint32(arousedLevel),
				PleasedLevel: uint32(pleasedLevel),
				Timestamp:    record.Time(),
			},
		)
	}

	return outputAfterExam, nil
}
