package main_test

import (
	logger "logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	deps := logger.Dependency{
		DB:          db,
		Org:         influxOrg,
		AccessToken: accessToken,
	}

	err := deps.Ping(rec, req)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expecting status code of %d, got: %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() == "" {
		t.Error("expecting a body got empty body")
	}
}
