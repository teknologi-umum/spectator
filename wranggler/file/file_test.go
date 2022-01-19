package file_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
	"worker/file"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var deps *file.Dependency
var db influxdb2.Client

func TestMain(m *testing.M) {
	// Lookup environment variables
	influxToken, ok := os.LookupEnv("INFLUX_TOKEN")
	if !ok {
		influxToken = "nMfrRYVcTyqFwDARAdqB92Ywj6GNMgPEd"
	}

	influxHost, ok := os.LookupEnv("INFLUX_HOST")
	if !ok {
		influxHost = "http://localhost:8086"
	}

	influxOrg, ok := os.LookupEnv("INFLUX_ORG")
	if !ok {
		influxOrg = "teknum_spectator"
	}

	minioHost, ok := os.LookupEnv("MINIO_HOST")
	if !ok {
		log.Fatalln("MINIO_HOST envar missing")
	}

	minioID, ok := os.LookupEnv("MINIO_ACCESS_ID")
	if !ok {
		log.Fatalln("MINIO_ACCESS_ID envar missing")
	}

	minioSecret, ok := os.LookupEnv("MINIO_SECRET_KEY")
	if !ok {
		log.Fatalln("MINIO_SECRET_KEY envar missing")
	}

	minioToken, ok := os.LookupEnv("MINIO_TOKEN")
	if !ok {
		log.Fatalln("MINIO_TOKEN envar missing")
	}

	db = influxdb2.NewClient(influxHost, influxToken)

	bucket, err := minio.New(
		minioHost,
		&minio.Options{
			Secure: false,
			Creds:  credentials.NewStaticV4(minioID, minioSecret, minioToken),
		},
	)
	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}

	deps = &file.Dependency{
		DB:                  db,
		DBOrganization:      influxOrg,
		Bucket:              bucket,
		BucketInputEvents:   "input_events",
		BucketSessionEvents: "session_events",
	}

	code := m.Run()

	db.Close()

	os.Exit(code)
}

func cleanup() {
	// create new context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// find current organization
	currentOrganization, err := deps.DB.OrganizationsAPI().FindOrganizationByName(ctx, deps.DBOrganization)
	if err != nil {
		log.Fatalf("finding organization: %v", err)
	}

	// find input_events bucket
	inputEventsBucket, err := deps.DB.BucketsAPI().FindBucketByName(ctx, deps.BucketInputEvents)
	if err != nil {
		log.Fatalf("finding bucket: %v", err)
	}

	// delete bucket data
	deleteAPI := db.DeleteAPI()

	inputEventMeasurements := []string{"ERROR", "WARNING", "INFO", "DEBUG", "CRITICAL"}
	for _, measurement := range inputEventMeasurements {
		err = deleteAPI.Delete(ctx, currentOrganization, inputEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\""+measurement+"\"")
		if err != nil {
			log.Fatalf("deleting bucket data: [%s] %v", measurement, err)
		}
	}

	// find input_events bucket
	sessionEventsBucket, err := db.BucketsAPI().FindBucketByName(ctx, deps.BucketSessionEvents)
	if err != nil {
		log.Fatalf("finding bucket: %v", err)
	}

	sessionEventMeasurements := []string{"ERROR", "WARNING", "INFO", "DEBUG", "CRITICAL"}
	for _, measurement := range sessionEventMeasurements {
		err = deleteAPI.Delete(ctx, currentOrganization, sessionEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\""+measurement+"\"")
		if err != nil {
			log.Fatalf("deleting bucket data: [%s] %v", measurement, err)
		}
	}
}
