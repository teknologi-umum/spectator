package file

import (
	"context"
	"fmt"
	"time"
	"worker/common"
	"worker/logger_proto"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// measurement: exam_started
type ExamStarted struct {
	Measurement     string    `json:"_measurement" csv:"_measurement"` // tag
	SessionId       string    `json:"session_id" csv:"session_id"`     // tag
	QuestionNumbers string    `json:"question_numbers" csv:"question_numbers"`
	Deadline        time.Time `json:"deadline" csv:"deadline"`
	Timestamp       time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryExamStarted(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (*ExamStarted, error) {
	examStartedRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementExamStarted+`" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return &ExamStarted{}, fmt.Errorf("failed to query exam_started: %w", err)
	}
	defer examStartedRows.Close()

	var outputExamStarted ExamStarted

	for examStartedRows.Next() {
		record := examStartedRows.Record()

		d.Logger.Log(
			record.Time().String(),
			logger_proto.Level_DEBUG.Enum(),
			sessionID.String(),
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "QueryExamStarted",
			},
		)

		questionNumbers, ok := record.ValueByKey("question_numbers").(string)
		if !ok {
			questionNumbers = ""
		}

		deadlineUnix := record.ValueByKey("deadline").(int64)
		if !ok {
			deadlineUnix = 0
		}

		deadline := time.UnixMilli(deadlineUnix)

		outputExamStarted = ExamStarted{
			Measurement:     common.MeasurementExamStarted,
			SessionId:       sessionID.String(),
			QuestionNumbers: questionNumbers,
			Deadline:        deadline,
			Timestamp:       record.Time(),
		}
	}

	return &outputExamStarted, nil
}
