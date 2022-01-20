package funfact

import (
	logger "worker/logger"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// Dependency contains the dependency injection
// to be used on this package.
type Dependency struct {
	Environment         string
	DB                  influxdb2.Client
	DBOrganization      string
	Logger              *logger.Logger
	LoggerToken         string
	BucketInputEvents   string
	BucketSessionEvents string
}
