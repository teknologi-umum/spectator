package file_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestQueryBeforeExamSAM(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for _, sessionID := range []uuid.UUID{globalID, globalID2} {
		readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
		result, err := deps.QueryBeforeExamSam(ctx, readInputAPI, sessionID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		if result.PleasedLevel != 5 {
			t.Logf("expecting PleasedLevel to be 5, got: %d", result.PleasedLevel)
		}

		if result.ArousedLevel != 2 {
			t.Logf("expecting ArousedLevel to be 2, got: %d", result.ArousedLevel)
		}
	}
}

func TestQueryAfterExamSAM(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for _, sessionID := range []uuid.UUID{globalID, globalID2} {
		readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
		result, err := deps.QueryAfterExamSam(ctx, readInputAPI, sessionID)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		if result.PleasedLevel != 5 {
			t.Logf("expecting PleasedLevel to be 5, got: %d", result.PleasedLevel)
		}

		if result.ArousedLevel != 2 {
			t.Logf("expecting ArousedLevel to be 2, got: %d", result.ArousedLevel)
		}
	}
}
