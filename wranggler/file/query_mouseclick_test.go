package file_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestQueryMouseClick(t *testing.T) {
	t.Cleanup(cleanup)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id, err := uuid.NewUUID()
	if err != nil {
		t.Error(err)
	}

	writeInputAPI := db.WriteAPIBlocking(deps.DBOrganization, deps.BucketInputEvents)

	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	for i := 0; i < 50; i++ {
		p := influxdb2.NewPoint(
			"coding_event_mouseclick",
			map[string]string{
				"session_id":      id.String(),
				"question_number": "1",
			},
			map[string]interface{}{
				"key_char":     "a",
				"right_click":  false,
				"left_click":   false,
				"middle_click": false,
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		)

		writeInputAPI.WritePoint(ctx, p)
	}

	readInputAPI := db.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryMouseClick(ctx, readInputAPI, id)
	if err != nil {
		t.Fatal("Test Query Mouse Click", err)
		return
	}

	if len(result) == 50 {
		t.Log("Test Query Mouse Click done")
	} else {
		t.Fatal("Data not 50")
	}
}
