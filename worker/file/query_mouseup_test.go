package file_test

import (
	"context"
	"testing"
	"time"
)

func TestQueryMouseUp(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id := globalID

	readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryMouseUp(ctx, readInputAPI, id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result) != 50 {
		t.Errorf("expected 50 results, got %d", len(result))
	}
}
