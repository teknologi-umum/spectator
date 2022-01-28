package file_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestQueryExamForfeited(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id, err := uuid.NewUUID()
	if err != nil {
		t.Errorf("failed to generate uuid: %v", err)
	}

	writeSessionAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketSessionEvents)

	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < 50; i++ {
		p := influxdb2.NewPointWithMeasurement("exam_forfeited")
		p.AddTag("session_id", id.String())
		p.AddField("_time", time.Unix(rand.Int63n(delta)+min, 0))

		err := writeSessionAPI.WritePoint(ctx, p)
		if err != nil {
			t.Error(err)
			return
		}
	}

	readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryExamForfeited(ctx, readInputAPI, id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(result) != 50 {
		t.Errorf("Expected 50 keystrokes, got %d", len(result))
	}
}
