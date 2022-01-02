package main_test

import (
	logger "logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	logger.NotAllowed(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expecting status code of %d, got: %d", http.StatusMethodNotAllowed, rec.Code)
	}

	if rec.Body.String() != "method not allowed" {
		t.Errorf("expecting body of 'method not allowed', got: %s", rec.Body.String())
	}
}
