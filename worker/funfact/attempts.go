package funfact

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// MeasurementSolution provides the name for the solution measurement
// type that is used for storing rejected or accepted test result.
type MeasurementSolution string

const (
	MeasurementSolutionRejected MeasurementSolution = "solution_rejected"
	MeasurementSolutionAccepted MeasurementSolution = "solution_accepted"
)

// Solution provides a union struct to host the solution_rejected
// and solution_accepted measurement from the InfluxDB or
// from any other type of input.
type Solution struct {
	Measurement          MeasurementSolution `json:"measurement"`
	SessionId            uuid.UUID           `json:"session_id"`
	QuestionNumber       int64               `json:"question_number"`
	Language             string              `json:"language"`
	Solution             string              `json:"solution"`
	Scratchpad           string              `json:"scratchpad"`
	SerializedTestResult string              `json:"serialized_test_result"`
}

func (d *Dependency) CalculateSubmissionAttempts(ctx context.Context, sessionID uuid.UUID, result chan uint32) error {
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// NOTE(2022-01-30): code_test_attempt has been changed into 2 measurements:
	// solution_accepted and solution_rejected, which contains a few data based on
	// the new Solution struct.

	// output contains the number of accepted and rejected solutions
	var output uint32

	solutionAcceptedRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "solution_accepted")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> yield()`,
	)
	if err != nil {
		return fmt.Errorf("failed to query solution_accepted measurement: %w", err)
	}
	defer solutionAcceptedRows.Close()

	for solutionAcceptedRows.Next() {
		output += 1
	}

	solutionRejectedRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "solution_rejected")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> yield()`,
	)
	if err != nil {
		return fmt.Errorf("failed to query solution_rejected measurement: %w", err)
	}
	defer solutionRejectedRows.Close()

	for solutionRejectedRows.Next() {
		output += 1
	}

	result <- output

	return nil
}
