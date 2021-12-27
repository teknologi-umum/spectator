package main_test

import (
	"bytes"
	"encoding/json"
	logger "logger"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/vmihailenco/msgpack/v5"
)

// Running it by each function is not helpful.
// So I guess it's better to run the whole log function
// in a single test.

func TestLogJSON(t *testing.T) {
	t.Cleanup(cleanup)

	deps := logger.Dependency{
		DB:          db,
		Org:         influxOrg,
		AccessToken: accessToken,
	}

	var payloads = []logger.Payload{
		{
			AccessToken: accessToken,
			Data: logger.Data{
				RequestID:   "a1",
				Application: "core",
				Message:     "A quick brown fox jumps over the lazy dog",
				Level:       "info",
				Environment: "production",
				Language:    "C#",
				Timestamp:   time.Now(),
			},
		},
		{
			AccessToken: accessToken,
			Data: logger.Data{
				RequestID:   "a1",
				Application: "worker",
				Message:     "Oh no, something went wrong",
				Level:       "error",
				Environment: "production",
				Language:    "Javascript",
				Body: map[string]interface{}{
					"stack_trace": "file.js:70 anotherfile.js:30",
					"why":         "I don't know",
				},
			},
		},
		{
			AccessToken: accessToken,
			Data: logger.Data{
				RequestID:   "b2",
				Application: "core",
				Message:     "Well, hello there. General Kenobi.",
			},
		},
		{
			AccessToken: accessToken,
			Data: logger.Data{
				RequestID:   "c3",
				Application: "worker",
				Message:     "This happened in the past",
				Timestamp:   time.Now().Add(time.Hour * 6 * -1),
			},
		},
	}

	// insert all data first
	for index, payload := range payloads {
		body, err := json.Marshal(payload)
		if err != nil {
			t.Errorf("marshaling json: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		req.Header.Set("content-type", "application/json")
		rec := httptest.NewRecorder()
		err = deps.InsertLog(rec, req)
		if err != nil {
			t.Errorf("[%d] an error was thrown: %v", index, err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("[%d] expecting status code of %d, got: %d", index, http.StatusOK, rec.Code)
		}
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	err := deps.ReadLogs(rec, req)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expecting status code of %d, got: %d", http.StatusOK, rec.Code)
	}

	urlQuery := url.Values{}
	urlQuery.Add("from", time.Time{}.Format(time.RFC3339))
	urlQuery.Add("to", time.Now().Format(time.RFC3339))
	urlQuery.Add("request_id", "a1")
	req = httptest.NewRequest(http.MethodGet, "/?"+urlQuery.Encode(), nil)
	rec = httptest.NewRecorder()
	err = deps.ReadLogs(rec, req)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expecting status code of %d, got: %d", http.StatusOK, rec.Code)
	}
}

func TestLogMsgpack(t *testing.T) {
	t.Cleanup(cleanup)

	deps := logger.Dependency{
		DB:          db,
		Org:         influxOrg,
		AccessToken: accessToken,
	}

	var payloads = []logger.Payload{
		{
			AccessToken: accessToken,
			Data: logger.Data{
				RequestID:   "a1",
				Application: "core",
				Message:     "A quick brown fox jumps over the lazy dog",
				Level:       "info",
				Environment: "production",
				Language:    "C#",
				Timestamp:   time.Now(),
			},
		},
		{
			AccessToken: accessToken,
			Data: logger.Data{
				RequestID:   "a1",
				Application: "worker",
				Message:     "Oh no, something went wrong",
				Level:       "error",
				Environment: "production",
				Language:    "Javascript",
				Body: map[string]interface{}{
					"stack_trace": "file.js:70 anotherfile.js:30",
					"why":         "I don't know",
				},
			},
		},
		{
			AccessToken: accessToken,
			Data: logger.Data{
				RequestID:   "b2",
				Application: "core",
				Message:     "Well, hello there. General Kenobi.",
			},
		},
		{
			AccessToken: accessToken,
			Data: logger.Data{
				RequestID:   "c3",
				Application: "worker",
				Message:     "This happened in the past",
				Timestamp:   time.Now().Add(time.Hour * 6 * -1),
			},
		},
	}

	// insert all data first
	for index, payload := range payloads {
		body, err := msgpack.Marshal(payload)
		if err != nil {
			t.Errorf("marshaling json: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		req.Header.Set("content-type", "application/msgpack")
		rec := httptest.NewRecorder()
		err = deps.InsertLog(rec, req)
		if err != nil {
			t.Errorf("[%d] an error was thrown: %v", index, err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("[%d] expecting status code of %d, got: %d", index, http.StatusOK, rec.Code)
		}
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("accept", "application/msgpack")
	rec := httptest.NewRecorder()
	err := deps.ReadLogs(rec, req)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expecting status code of %d, got: %d", http.StatusOK, rec.Code)
	}

	urlQuery := url.Values{}
	urlQuery.Add("from", time.Time{}.Format(time.RFC3339))
	urlQuery.Add("to", time.Now().Format(time.RFC3339))
	urlQuery.Add("request_id", "a1")
	req = httptest.NewRequest(http.MethodGet, "/?"+urlQuery.Encode(), nil)
	req.Header.Set("accept", "application/msgpack")
	rec = httptest.NewRecorder()
	err = deps.ReadLogs(rec, req)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expecting status code of %d, got: %d", http.StatusOK, rec.Code)
	}
}

// Should supply explicit Content-Type headers
func TestContentTypeTeapot(t *testing.T) {
	t.Cleanup(cleanup)

	deps := logger.Dependency{
		DB:          db,
		Org:         influxOrg,
		AccessToken: accessToken,
	}
	payload := logger.Payload{
		AccessToken: accessToken,
		Data: logger.Data{
			RequestID:   "a1",
			Application: "core",
			Message:     "A quick brown fox jumps over the lazy dog",
			Level:       "info",
			Environment: "production",
			Language:    "C#",
			Timestamp:   time.Now(),
		},
	}

	body, err := msgpack.Marshal(payload)
	if err != nil {
		t.Errorf("marshaling json: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	err = deps.InsertLog(rec, req)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if rec.Code != http.StatusTeapot {
		t.Errorf("expecting status code of %d, got: %d", http.StatusTeapot, rec.Code)
	}
}

func TestValidatePayload(t *testing.T) {
	deps := logger.Dependency{
		DB:          db,
		Org:         influxOrg,
		AccessToken: accessToken,
	}

	t.Run("empty", func(t *testing.T) {
		p := logger.Payload{}
		err := deps.ValidatePayload(p)
		if err == nil || err.Error() != "access token must be provided" {
			t.Errorf("expecting an error, instead got: %v", err)
		}
	})

	t.Run("missing", func(t *testing.T) {
		p := logger.Payload{AccessToken: accessToken}
		err := deps.ValidatePayload(p)
		if err == nil || err.Error() != "proper request_id, application, message must be provided" {
			t.Errorf("expecting an error, instead got: %v", err)
		}
	})

	t.Run("commas", func(t *testing.T) {
		p := logger.Payload{
			AccessToken: accessToken,
			Data: logger.Data{
				RequestID:   "bla,bla",
				Application: "asd,asd",
				Message:     "hello there",
			},
		}
		err := deps.ValidatePayload(p)
		if err == nil || err.Error() != "proper request_id, application must be provided" {
			t.Errorf("expecting an error, nistead got: %v", err)
		}
	})
}
