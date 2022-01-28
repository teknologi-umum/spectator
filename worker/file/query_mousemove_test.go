package file_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestQueryMouseMove(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id, err := uuid.NewUUID()
	if err != nil {
		t.Errorf("failed to generate uuid: %v", err)
	}

	writeInputAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketInputEvents)

	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < 50; i++ {
		p := influxdb2.NewPoint(
			"coding_event_mousemove",
			map[string]string{
				"session_id":      id.String(),
				"question_number": "1",
			},
			map[string]interface{}{
				"direction":     "right",
				"x_position":    rand.Int31n(1337),
				"y_position":    rand.Int31n(768),
				"window_width":  rand.Int31n(1337),
				"window_height": rand.Int31n(768),
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		)

		err := writeInputAPI.WritePoint(ctx, p)
		if err != nil {
			t.Error(err)
			return
		}
	}

	readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryMouseMove(ctx, readInputAPI, id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 50 {
		t.Errorf("expected 50 results, got %d", len(result))
	}
}
