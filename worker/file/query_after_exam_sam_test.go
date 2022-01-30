package file_test

import (
	"context"
	"testing"
	"time"
)

func TestQueryAfterExamSAM(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id := globalID

	readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryAfterExamSam(ctx, readInputAPI, id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(result) != 50 {
		t.Errorf("Expected 50 results, got %d", len(result))
	}
}
