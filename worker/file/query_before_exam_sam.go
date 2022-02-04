package file

import (
	"context"
	"fmt"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: before_exam_sam_submitted
type BeforeExamSAMSubmitted struct {
	Measurement  string    `json:"_measurement" csv:"_measurement"` // tag
	SessionId    string    `json:"session_id" csv:"session_id"`     // Tag
	ArousedLevel uint32    `json:"aroused_level" csv:"aroused_level"`
	PleasedLevel uint32    `json:"pleased_level" csv:"pleased_level"`
	Timestamp    time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryBeforeExamSam(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]BeforeExamSAMSubmitted, error) {
	beforeExamSamRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementBeforeExamSAMSubmitted+`" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return []BeforeExamSAMSubmitted{}, fmt.Errorf("failed to query before_exam_sam_submitted: %w", err)
	}
	defer beforeExamSamRows.Close()

	var outputBeforeExam []BeforeExamSAMSubmitted

	for beforeExamSamRows.Next() {
		record := beforeExamSamRows.Record()

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

		outputBeforeExam = append(
			outputBeforeExam,
			BeforeExamSAMSubmitted{
				Measurement:  common.MeasurementBeforeExamSAMSubmitted,
				SessionId:    sessionId,
				ArousedLevel: uint32(arousedLevel),
				PleasedLevel: uint32(pleasedLevel),
				Timestamp:    record.Time(),
			},
		)
	}

	return outputBeforeExam, nil
}
