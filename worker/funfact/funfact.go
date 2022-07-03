package funfact

import (
	logger "worker/logger"
	"worker/status"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// Dependency contains the dependency injection
// to be used on this package.
type Dependency struct {
	Environment    string
	DB             influxdb2.Client
	DBOrganization string
	Logger         *logger.Logger
	LoggerToken    string
	Status         *status.Dependency
}

// KeystrokeInput contains the data of
// keystroke measurement.
type KeystrokeInput struct {
	UnrelatedKey bool
	Shift        bool
	Alt          bool
	Control      bool
	Meta         bool
	KeyChar      string
}
