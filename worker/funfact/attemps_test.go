package funfact_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestCalculateSubmissionAttempts(t *testing.T) {
	t.Cleanup(cleanup)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	id, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("failed to generate uuid: %v", err)
	}

	writeAPI := db.WriteAPIBlocking(deps.DBOrganization, deps.BucketSessionEvents)

	// Random date between range
	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < 25; i++ {
		point := influxdb2.NewPoint(
			"code_test_attempt",
			map[string]string{
				"session_id":  id.String(),
				"question_id": strconv.Itoa(rand.Intn(5)),
			},
			map[string]interface{}{
				"code":     "console.log('Hello world!');",
				"language": "javascript",
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		)

		err = writeAPI.WritePoint(ctx, point)
		if err != nil {
			t.Fatalf("writing point: %v", err)
		}
	}

	res := make(chan uint32, 1)
	err = deps.CalculateSubmissionAttempts(ctx, id, res)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	attemps := <-res
	t.Logf("submission attemps: %v", attemps)
}
