package file

import (
	"encoding/csv"
	"fmt"
	"strings"

	"worker/logger"

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
	Logger              *logger.Logger
	LoggerToken         string
	BucketInputEvents   string
	BucketSessionEvents string
}

// UnmarshalInfluxRow converts a row from InfluxDB into a map[string]interface{}
//
// Deprecated: use regular row parsing provided by InfluxDB client library
func UnmarshalInfluxRow(row string) (map[string]interface{}, error) {
	// because csv.NewReader() accepts io.Reader, we'll create one from strings pkg
	input := strings.NewReader(row)
	reader := csv.NewReader(input)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	records, err := reader.Read()
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("reading row value to csv: %v", err)
	}

	// find records length
	// because it's a jagged array, we'll do a nested one
	var recordsLength = len(records)

	output := make(map[string]interface{}, recordsLength)
	for _, rec := range records {
		kv := strings.Split(rec, ":")
		output[kv[0]] = kv[1]
	}

	return output, nil
}
