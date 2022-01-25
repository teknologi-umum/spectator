package influxhelpers_test

import (
	"testing"
	"time"
	"worker/influxhelpers"
)

func TestReinaldysBuildQuery(t *testing.T) {
	query := influxhelpers.Queries{
		Buckets:     "events",
		Measurement: "android",
		SessionID:   "5f90e248-9759-499e-bf25-04dafb3e0f94",
		Field:       "roma",
		TimeFrom:    time.Date(2019, 10, 9, 8, 7, 6, 0, time.UTC),
		TimeTo:      time.Date(2019, 12, 11, 10, 9, 8, 0, time.UTC),
	}

	s := influxhelpers.ReinaldysBuildQuery(query)

	expected := `from(bucket: "events")
|> range(start: 1570608426, stop: 1576058948)
|> sort(columns: ["_time"])
|> group(columns: ["_time"])
|> filter(fn: (r) => r["_measurement"] == "android")
|> filter(fn: (r) => r["session_id"] == "5f90e248-9759-499e-bf25-04dafb3e0f94")
|> filter(fn: (r) => r["_field"] == "roma")
|> yield()`

	if s != expected {
		t.Errorf("expecting: %s, got: %s", expected, s)
	}
}

func TestReinaldysBuildQuery_Panic(t *testing.T) {
	defer func() {
		r := recover()
		if r.(string) != "query builder: bucket shall not be empty" {
			t.Fatalf(
				"expecting a panic of: %s, got: %v",
				"query builder: bucket shall not be empty",
				r,
			)
		}
	}()

	_ = influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{})
}

func TestReinaldysBuildQuery_Alternate(t *testing.T) {
	query := influxhelpers.Queries{
		Buckets: "events",
	}

	s := influxhelpers.ReinaldysBuildQuery(query)

	expected := `from(bucket: "events")
|> range(start: 0)
|> sort(columns: ["_time"])
|> group(columns: ["session_id", "_time"])
|> yield()`

	if s != expected {
		t.Errorf("expecting: %s, got: %s", expected, s)
	}
}

func TestSanitize(t *testing.T) {
	query := influxhelpers.Queries{
		Buckets:     "events",
		Measurement: `") // hello world`,
		SessionID:   `"lorem ipsum""`,
		Field:       `")\n\t\t// hello world`,
	}

	s := influxhelpers.ReinaldysBuildQuery(query)

	expected := `from(bucket: "events")
|> range(start: 0)
|> sort(columns: ["_time"])
|> group(columns: ["_time"])
|> filter(fn: (r) => r["_measurement"] == "\") \/\/ hello world")
|> filter(fn: (r) => r["session_id"] == "\"lorem ipsum\"\"")
|> filter(fn: (r) => r["_field"] == "\")\\n\\t\\t\/\/ hello world")
|> yield()`

	if s != expected {
		t.Errorf("expecting: %s, got: %s", expected, s)
	}
}
