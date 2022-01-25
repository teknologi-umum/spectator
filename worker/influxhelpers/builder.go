package influxhelpers

import (
	"strconv"
	"strings"
	"time"
)

type Queries struct {
	Buckets     string
	Measurement string
	SessionID   string
	Field       string
	TimeFrom    time.Time
	TimeTo      time.Time
}

// ReinaldysBuildQuery builds an InfluxDB (or a Flux) query from the predefined
// Queries struct. It will do an SQL injection escaping (or sanitizing), to prevent
// malicioous attacks on event that the string inside the Queries files includes
// a double-slash (//) and/or a double quotes.
//
// It will panic if the "Buckets" field is empty.
func ReinaldysBuildQuery(q Queries) string {
	if q.Buckets == "" {
		panic("query builder: bucket shall not be empty")
	}

	var str strings.Builder
	str.WriteString("from(bucket: \"" + sanitize(q.Buckets) + "\")\n")
	// range query
	str.WriteString("|> range(")
	if !q.TimeFrom.IsZero() {
		str.WriteString("start: " + strconv.FormatInt(q.TimeFrom.Unix(), 10))
	} else {
		str.WriteString("start: 0")
	}

	if !q.TimeTo.IsZero() {
		str.WriteString(", stop: " + strconv.FormatInt(q.TimeTo.Unix(), 10))
	}

	str.WriteString(")\n")

	str.WriteString("|> sort(columns: [\"_time\"])\n")

	if q.SessionID == "" {
		str.WriteString("|> group(columns: [\"session_id\", \"_time\"])\n")
	} else {
		str.WriteString("|> group(columns: [\"_time\"])\n")
	}

	if q.Measurement != "" {
		str.WriteString(`|> filter(fn: (r) => r["_measurement"] == "` + sanitize(q.Measurement) + `")` + "\n")
	}

	if q.SessionID != "" {
		str.WriteString(`|> filter(fn: (r) => r["session_id"] == "` + sanitize(q.SessionID) + `")` + "\n")
	}

	if q.Field != "" {
		str.WriteString(`|> filter(fn: (r) => r["_field"] == "` + sanitize(q.Field) + `")` + "\n")
	}

	str.WriteString("|> yield()")

	return str.String()
}

// sanitize will sanitize the incoming string that is concatenated on the query
// to prevent malicious SQL injection attack.
func sanitize(str string) string {
	s := str

	if strings.Contains(s, "\\") {
		s = strings.ReplaceAll(s, "\\", "\\\\")
	}

	if strings.Contains(s, "//") {
		s = strings.ReplaceAll(s, "//", "\\/\\/")
	}

	if strings.Contains(s, "\"") {
		s = strings.ReplaceAll(s, "\"", "\\\"")
	}

	return s
}
