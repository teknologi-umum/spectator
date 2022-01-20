package file

import (
	"strconv"
	"strings"
	"time"
)

type queries struct {
	Level     string
	SessionID string
	Buckets   string
	TimeFrom  time.Time
	TimeTo    time.Time
}

func (d *Dependency) IsDebug() bool {
	return d.Environment == "DEVELOPMENT"
}

func reinaldysBuildQuery(q queries) string {
	var str strings.Builder
	str.WriteString("from(bucket: \"" + q.Buckets + "\")\n")
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

	if q.Level != "" {
		str.WriteString(`|> filter(fn: (r) => r["_measurement"] == "` + q.Level + `")` + "\n")
	}

	if q.SessionID != "" {
		str.WriteString(`|> filter(fn: (r) => r["session_id"] == "` + q.SessionID + `")` + "\n")
	}

	str.WriteString("|> yield()\n")

	return str.String()
}
