package file_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestQueryExamEnded(t *testing.T) {
	// TODO:
	// 1. insert some fake data into the influx db
	// 2. query the data with the function
	// 3. compare the length of the result and the length of fake data
	// add another test (maybe a subtest, or another test function)
	// that checks if there is no data to be queried.
	// we must check if that (rare and edgy) event happen,
	// so what would the software react?
	t.Cleanup(cleanup)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id, err := uuid.NewUUID()
	if err != nil {
		t.Errorf("failed to generate uuid: %v", err)
	}

	writeSessionAPI := db.WriteAPIBlocking(deps.DBOrganization, deps.BucketSessionEvents)

	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < 50; i++ {
		p := influxdb2.NewPoint(
			"exam_ended",
			map[string]string{
				"session_id": id.String(),
			},
			map[string]interface{}{},
			time.Unix(rand.Int63n(delta)+min, 0),
		)

		writeSessionAPI.WritePoint(ctx, p)
	}

	readInputAPI := db.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryKeystrokes(ctx, readInputAPI, id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(result) != 50 {
		t.Errorf("Expected 50 keystrokes, got %d", len(result))
	}
}
