package file_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestQueryExamForfeited(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for _, sessionID := range []uuid.UUID{globalID, globalID2} {
		readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
		_, err := deps.QueryExamForfeited(ctx, readInputAPI, sessionID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	}
}

func TestQueryIDEReloaded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for _, sessionID := range []uuid.UUID{globalID, globalID2} {
		readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
		_, err := deps.QueryExamIDEReloaded(ctx, readInputAPI, sessionID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	}
}

func TestQueryExamEnded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for _, sessionID := range []uuid.UUID{globalID, globalID2} {
		readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
		_, err := deps.QueryExamEnded(ctx, readInputAPI, sessionID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	}
}
