package file_test

import (
	"context"
	"testing"
	"time"
)

func TestQuerySolutionAccepted(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)

	solutions, err := deps.QuerySolutionAccepted(ctx, readInputAPI, globalID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(solutions) == 0 {
		t.Logf("expected 'solutions' to have a slice length greater than zero, got zero instead")
	}
}

func TestQuerySolutionRejected(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	readInputAPI := deps.DB.QueryAPI(deps.DBOrganization)

	solutions, err := deps.QuerySolutionRejected(ctx, readInputAPI, globalID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(solutions) == 0 {
		t.Logf("expected 'solutions' to have a slice length greater than zero, got zero instead")
	}
}
