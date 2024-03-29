package file

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/rs/zerolog/log"
)

type TestEvent struct {
	Measurement          string    `json:"_measurement" csv:"_measurement"`
	SessionID            string    `json:"session_id" csv:"session_id"` // tag
	Language             string    `json:"language" csv:"language"`
	Solution             string    `json:"solution" csv:"solution"`
	ScratchPad           string    `json:"scratch_pad" csv:"scratch_pad"`
	SerializedTestResult string    `json:"serialized_test_result" csv:"serialized_test_result"`
	QuestionNumber       int64     `json:"question_number" csv:"question_number"`
	Timestamp            time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryTestAccepted(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]TestEvent, error) {
	return d.queryTest(ctx, queryAPI, sessionID, common.MeasurementTestAccepted)
}

func (d *Dependency) QueryTestRejected(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]TestEvent, error) {
	return d.queryTest(ctx, queryAPI, sessionID, common.MeasurementTestRejected)
}

func (d *Dependency) queryTest(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID, measurement string) ([]TestEvent, error) {
	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+measurement+`" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return []TestEvent{}, fmt.Errorf("failed to query solution for measurement %s: %v", measurement, err)
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Err(err).Msg("closing querySolution rows")
		}
	}()

	var outputTests []TestEvent

	for rows.Next() {
		record := rows.Record()

		if record.Time().Year() != 2022 {
			log.Warn().
				Str("current time from record.Time() is not 2022, it's ", strconv.Itoa(record.Time().Year())).
				Msg("invalid date on querySolution")
		}

		questionNumber, ok := record.ValueByKey("question_number").(int64)
		if !ok {
			questionNumber = 0
		}

		language, ok := record.ValueByKey("language").(string)
		if !ok {
			language = ""
		}

		solution, ok := record.ValueByKey("solution").(string)
		if !ok {
			solution = ""
		}

		scratchpad, ok := record.ValueByKey("scratch_pad").(string)
		if !ok {
			scratchpad = ""
		}

		serializedTestResult, ok := record.ValueByKey("serialized_test_result").(string)
		if !ok {
			serializedTestResult = ""
		}

		outputTests = append(
			outputTests,
			TestEvent{
				SessionID:            sessionID.String(),
				Timestamp:            record.Time(),
				QuestionNumber:       questionNumber,
				Language:             language,
				Solution:             solution,
				ScratchPad:           scratchpad,
				SerializedTestResult: serializedTestResult,
				Measurement:          measurement,
			},
		)
	}

	return outputTests, nil
}
