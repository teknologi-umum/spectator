package file_test

import (
	"context"
	"testing"
	"time"
)

func TestQueryPersonalInfo(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id := globalID

	readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)
	result, err := deps.QueryPersonalInfo(ctx, readInputAPI, id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if result.SessionID != id.String() {
		t.Errorf("personal info not exist")
	}
}
