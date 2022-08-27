package file_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestQueryKeystrokes(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for _, sessionID := range []uuid.UUID{globalID, globalID2} {
		readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
		result, err := deps.QueryKeystrokes(ctx, readInputAPI, sessionID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		if len(result) != 50 {
			t.Logf("Expected 50 results, got %d", len(result))
		}
	}
}
