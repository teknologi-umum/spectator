package file_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestQueryExamStarted(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for _, sessionID := range []uuid.UUID{globalID, globalID2} {
		readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
		_, err := deps.QueryExamStarted(ctx, readInputAPI, sessionID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	}
}
