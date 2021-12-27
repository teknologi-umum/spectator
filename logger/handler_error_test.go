package main_test

import (
	"errors"
	logger "logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	err := errors.New("something happened")
	logger.HandleError(err, rec, req)
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expecting status code of %d, got: %d", http.StatusInternalServerError, rec.Code)
	}

	if rec.Body.String() != "something happened" {
		t.Errorf("expecting body of 'something happened', got: %s", rec.Body.String())
	}
}
