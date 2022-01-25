package file_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestQuerySamTest(t *testing.T) {
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
			"sam_test",
			map[string]string{
				"session_id": id.String(),
			},
			map[string]interface{}{
				"aroused_level": rand.Int31n(3),
				"pleased_level": rand.Int31n(3),
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		)

		writeSessionAPI.WritePoint(ctx, p)
	}

	readInputAPI := db.QueryAPI(deps.DBOrganization)
	result, err := deps.QuerySAMTest(ctx, readInputAPI, id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 50 {
		t.Errorf("expected 50 results, got %d", len(result))
	}
}
