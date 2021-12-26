package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/vmihailenco/msgpack/v5"
)

type Payload struct {
	AccessToken string `json:"access_token" msgpack:"access_token"`
	Data        Data   `json:"data" msgpack:"data"`
}

type Data struct {
	RequestID   string                 `json:"request_id" msgpack:"request_id"`
	Application string                 `json:"application" msgpack:"application"`
	Message     string                 `json:"message" msgpack:"message"`
	Body        map[string]interface{} `json:"body" msgpack:"body"`
	Level       string                 `json:"level" msgpack:"level"`
	Environment string                 `json:"environment" msgpack:"environment"`
	Language    string                 `json:"language" msgpack:"language"`
	Timestamp   time.Time              `json:"timestamp" msgpack:"timestamp"`
}

func (d *Dependency) ValidatePayload(p Payload) error {
	if p.AccessToken != d.AccessToken {
		return fmt.Errorf("access token must be provided")
	}

	var missing []string
	if p.Data.RequestID == "" {
		missing = append(missing, "request_id")
	}

	if p.Data.Application == "" {
		missing = append(missing, "application")
	}

	if p.Data.Message == "" {
		missing = append(missing, "message")
	}

	if len(missing) == 0 {
		return nil
	}

	return fmt.Errorf("%s must be provided", strings.Join(missing, ", "))
}

func (d *Dependency) InsertLog(w http.ResponseWriter, r *http.Request) error {
	var body Payload
	switch r.Header.Get("content-type") {
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			return err
		}
	case "application/msgpack":
		err := msgpack.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			return err
		}
	default:
		w.WriteHeader(http.StatusTeapot)
		w.Header().Set("content-type", "text/plain")
		w.Write([]byte("content-type header must be provided"))
		return nil
	}

	err := d.ValidatePayload(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			return err
		}
		return nil
	}

	err = d.writeIntoLog(r.Context(), body)
	if err != nil {
		return fmt.Errorf("writing log: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "text/plain")
	w.Write([]byte("OK"))

	return nil
}

type queries struct {
	Level       string
	RequestID   string
	Application string
	TimeFrom    time.Time
	TimeTo      time.Time
}

func (d *Dependency) ReadLogs(w http.ResponseWriter, r *http.Request) error {
	urlQuery := r.URL.Query()
	query := queries{
		Level:       urlQuery.Get("level"),
		RequestID:   urlQuery.Get("request_id"),
		Application: urlQuery.Get("application"),
	}

	if from := urlQuery.Get("from"); from != "" {
		timeFrom, err := time.Parse(time.RFC3339, from)
		if err != nil {
			return fmt.Errorf("parsing from query: %v", err)
		}
		query.TimeFrom = timeFrom
	}

	if to := urlQuery.Get("to"); to != "" {
		timeTo, err := time.Parse(time.RFC3339, to)
		if err != nil {
			return fmt.Errorf("parsing to query: %v", err)
		}
		query.TimeTo = timeTo
	}

	logs, err := d.fetchLog(r.Context(), query)
	if err != nil {
		return err
	}

	var contentType string
	if r.Header.Get("accept") == "" {
		contentType = "application/json"
	}

	var data []byte
	switch contentType {
	case "application/msgpack":
		data, err = msgpack.Marshal(logs)
		if err != nil {
			return err
		}
	default:
		data, err = json.Marshal(logs)
		if err != nil {
			return err
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", contentType)
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
