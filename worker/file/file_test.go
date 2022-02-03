package file_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
	"worker/common"
	"worker/file"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	deps      *file.Dependency
	globalID  uuid.UUID
	globalID2 uuid.UUID
)

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
		minioHost = "localhost:9000"
	}

	minioID, ok := os.LookupEnv("MINIO_ACCESS_ID")
	if !ok {
		minioID = "teknum"
	}

	minioSecret, ok := os.LookupEnv("MINIO_SECRET_KEY")
	if !ok {
		minioSecret = "c2N9Xz8bzHPkgNcgDtKgwGPTdb76GjD48"
	}

	minioToken, ok := os.LookupEnv("MINIO_TOKEN")
	if !ok {
		minioToken = ""
	}

	db := influxdb2.NewClient(influxHost, influxToken)

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
		DB:             db,
		DBOrganization: influxOrg,
		Bucket:         bucket,
		Environment:    "testing",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	// Check for bucket existence
	err = prepareBuckets(ctx, deps.DB, influxOrg)
	if err != nil {
		log.Fatalf("Failed to prepare influxdb buckets: %v", err)
	}

	err = seedData(ctx)
	if err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	code := m.Run()

	fmt.Println("Cleaning up...")

	err = cleanup(ctx)
	if err != nil {
		log.Fatalf("Failed to cleanup: %v", err)
	}

	deps.DB.Close()

	os.Exit(code)
}

// prepareBuckets creates the buckets if they don't exist
func prepareBuckets(ctx context.Context, db influxdb2.Client, org string) error {
	bucketsAPI := deps.DB.BucketsAPI()
	organizationAPI := deps.DB.OrganizationsAPI()

	bucketNames := []string{common.BucketInputEvents, common.BucketSessionEvents, common.BucketFileEvents}

	for _, bucket := range bucketNames {
		_, err := bucketsAPI.FindBucketByName(ctx, bucket)
		if err != nil && err.Error() != "bucket '"+bucket+"' not found" {
			return fmt.Errorf("finding bucket: %v", err)
		}

		if err != nil && err.Error() == "bucket '"+bucket+"' not found" {
			orgDomain, err := organizationAPI.FindOrganizationByName(ctx, org)
			if err != nil {
				return fmt.Errorf("finding organization: %v", err)
			}

			_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, bucket)
			if err != nil && err.Error() != "conflict: bucket with name "+bucket+" already exists" {
				return fmt.Errorf("creating bucket: %v", err)
			}
		}
	}

	return nil
}

// cleanup deletes the buckets' data
func cleanup(ctx context.Context) error {
	// find current organization
	currentOrganization, err := deps.DB.OrganizationsAPI().FindOrganizationByName(ctx, deps.DBOrganization)
	if err != nil {
		return fmt.Errorf("finding organization: %v", err)
	}

	// delete bucket data
	deleteAPI := deps.DB.DeleteAPI()

	// find input_events bucket
	inputEventsBucket, err := deps.DB.BucketsAPI().FindBucketByName(ctx, common.BucketInputEvents)
	if err != nil {
		return fmt.Errorf("finding bucket: %v", err)
	}

	fileEventMeasurement := []string{
		common.MeasurementKeystroke,
		common.MeasurementMouseDown,
		common.MeasurementMouseUp,
		common.MeasurementMouseMoved,
		common.MeasurementMouseScrolled,
		common.MeasurementWindowSized,
	}
	for _, measurement := range fileEventMeasurement {
		err = deleteAPI.Delete(ctx, currentOrganization, inputEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\""+measurement+"\"")
		if err != nil {
			return fmt.Errorf("deleting bucket data: [%s] %v", measurement, err)
		}
	}

	// find input_events bucket
	sessionEventsBucket, err := deps.DB.BucketsAPI().FindBucketByName(ctx, common.BucketSessionEvents)
	if err != nil {
		return fmt.Errorf("finding bucket: %v", err)
	}

	sessionEventMeasurements := []string{
		common.MeasurementCodeTestAttempt,
		common.MeasurementExamForfeited,
		common.MeasurementExamEnded,
		common.MeasurementExamStarted,
		common.MeasurementSolutionRejected,
		common.MeasurementSolutionAccepted,
		common.MeasurementSessionStarted,
		common.MeasurementPersonalInfoSubmitted,
		common.MeasurementLocaleSet,
		common.MeasurementExamIDEReloaded,
		common.MeasurementDeadlinePassed,
		common.MeasurementBeforeExamSAMSubmitted,
		common.MeasurementAfterExamSAMSubmitted,
	}
	for _, measurement := range sessionEventMeasurements {
		err = deleteAPI.Delete(ctx, currentOrganization, sessionEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\""+measurement+"\"")
		if err != nil {
			return fmt.Errorf("deleting bucket data: [%s] %v", measurement, err)
		}
	}

	// find file_results bucket
	fileEventsBucket, err := deps.DB.BucketsAPI().FindBucketByName(ctx, common.BucketFileEvents)
	if err != nil {
		return fmt.Errorf("finding bucket: %v", err)
	}

	for _, measurement := range fileEventMeasurement {
		err = deleteAPI.Delete(ctx, currentOrganization, fileEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\"exported_data\"")
		if err != nil {
			return fmt.Errorf("deleting bucket data: [%s] %v", measurement, err)
		}
	}

	// delete json/csv files
	pathJSON, err := filepath.Glob("./*_*.json")
	if err != nil {
		return fmt.Errorf("unexpected error: %v", err)
	}
	pathCSV, err := filepath.Glob("./*_*.csv")
	if err != nil {
		return fmt.Errorf("unexpected error: %v", err)
	}

	for _, path := range append(pathJSON, pathCSV...) {
		err = os.Remove(path)
		if err != nil {
			return fmt.Errorf("unexpected error: %v", err)
		}
	}

	return nil
}
