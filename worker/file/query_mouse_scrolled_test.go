package file_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestQueryMouseScrolled(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for _, sessionID := range []uuid.UUID{globalID, globalID2} {
		readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
		result, err := deps.QueryMouseScrolled(ctx, readInputAPI, sessionID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(*result) != 50 {
			t.Errorf("expected 50 results, got %d", len(*result))
		}
	}
}
