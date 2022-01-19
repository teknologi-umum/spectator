package funfact

import (
	logger "worker/logger_proto"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
)

// Dependency contains the dependency injection
// to be used on this package.
type Dependency struct {
	Environment         string
	DB                  influxdb2.Client
	Bucket              *minio.Client
	DBOrganization      string
	Logger              logger.LoggerClient
	LoggerToken         string
	BucketInputEvents   string
	BucketSessionEvents string
}
