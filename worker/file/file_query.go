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
