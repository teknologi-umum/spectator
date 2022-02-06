// file package provides the domain for creating JSON and CSV files
// for the user. Those created files are meant to be uploaded into
// MinIO or any other S3 compatible storage.
package file

import (
	"worker/logger"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
)

// Dependency contains the dependency injection
// to be used on this package.
type Dependency struct {
	Environment    string
	DB             influxdb2.Client
	Bucket         *minio.Client
	DBOrganization string
	Logger         *logger.Logger
	LoggerToken    string
}
