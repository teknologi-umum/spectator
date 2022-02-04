package file

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Solution struct {
	SessionID            string    `json:"session_id" csv:"session_id"` // tag
	Timestamp            time.Time `json:"timestamp" csv:"timestamp"`
	QuestionNumber       int64     `json:"question_number" csv:"question_number"`
	Language             string    `json:"language" csv:"language"`
	Solution             string    `json:"solution" csv:"solution"`
	ScratchPad           string    `json:"scratch_pad" csv:"scratch_pad"`
	SerializedTestResult string    `json:"serialized_test_result" csv:"serialized_test_result"`
	Measurement          string    `json:"_measurement" csv:"_measurement"`
}

// TODO: implement
func (d *Dependency) QuerySolutionAccepted(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]Solution, error) {
	return []Solution{}, nil
}

// TODO: implement
func (d *Dependency) QuerySolutionRejected(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]Solution, error) {
	return []Solution{}, nil
}
