package file_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestQueryPersonalInfo(t *testing.T) {
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
			"personal_info",
			map[string]string{
				"session_id": id.String(),
			},
			map[string]interface{}{
				"student_number":      "",
				"hours_of_practice":   rand.Int31n(666),
				"years_of_experience": rand.Int31n(5),
				"familiar_languages":  "",
			},
			time.Unix(rand.Int63n(delta)+min, 0),
		)

		writeSessionAPI.WritePoint(ctx, p)
	}

	readInputAPI := db.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryPersonalInfo(ctx, readInputAPI, id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if result.SessionID != id.String() {
		t.Errorf("personal info not exist")
	}
}
