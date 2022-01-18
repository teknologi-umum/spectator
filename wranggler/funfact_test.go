package main_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	worker "worker"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestCalculateWordsPerMinute(t *testing.T) {
	t.Cleanup(cleanup)
	rand.Seed(time.Now().Unix())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	deps := worker.Dependency{
		DB:             db,
		DBOrganization: dbOrganization,
		Bucket:         bucket,
	}

	id, err := uuid.NewUUID()
	if err != nil {
		t.Error(err)
	}

	writeAPI := db.WriteAPIBlocking(deps.DBOrganization, worker.BucketInputEvents)

	// Random date between range
	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < 50; i++ {
		point := influxdb2.NewPoint(
			"coding_event_keystroke",
			map[string]string{
				"session_id": id.String(),
			},
			map[string]interface{}{
				"key_char": "a",
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		)
		writeAPI.WritePoint(ctx, point)
	}

	res := make(chan uint32, 1)
	err = deps.CalculateWordsPerMinute(ctx, id, res)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	t.Logf("average wpm: %v", <-res)
}
