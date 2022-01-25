package file

import (
	"time"
)

func (d *Dependency) IsDebug() bool {
	return d.Environment == "DEVELOPMENT"
}

type Queries struct {
	Level     string
	SessionID string
	Buckets   string
	TimeFrom  time.Time
	TimeTo    time.Time
}

// `from(bucket: "`+d.BucketInputEvents+`")
// |> range(start: 0)
// |> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
// |> filter(fn : (r) => r["_measurement"] == "coding_event_mousemove")
