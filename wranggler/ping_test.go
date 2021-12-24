package main_test

import (
	"net/http"
	"net/http/httptest"
	main "rori"
	"testing"
)

// TestPing will test the ping handler for this worker.
func TestPing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	dependencies := &main.Dependency{}

	defer func() {
		r := recover()
		if r != nil {
			t.Fatal("panicked!", r)
		}
	}()

	dependencies.Ping(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expecting http status response of %d, got: %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != "Hello world" {
		t.Errorf("expecting http respohsne body of \"Hello world\", got: %s", rec.Body.String())
	}
}
