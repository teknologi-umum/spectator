package file_test

import (
	"context"
	"testing"
	"time"
)

func TestQueryTestAccepted(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)

	tests, err := deps.QueryTestAccepted(ctx, readInputAPI, globalID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(tests) == 0 {
		t.Logf("expected 'solutions' to have a slice length greater than zero, got zero instead")
	}
}

func TestQueryTestRejected(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)

	tests, err := deps.QueryTestRejected(ctx, readInputAPI, globalID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(tests) == 0 {
		t.Logf("expected 'solutions' to have a slice length greater than zero, got zero instead")
	}
}
