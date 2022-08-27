package file_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
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

	// Setup a context for preparing things
	prepareCtx, prepareCancel := context.WithTimeout(context.Background(), time.Second*120)

	// Check for InfluxDB buckets existence
	err = prepareBuckets(prepareCtx, deps.DB, influxOrg)
	if err != nil {
		log.Fatalf("failed to prepare influxdb buckets: %v", err)
	}

	err = seedData(prepareCtx)
	if err != nil {
		log.Fatalf("failed to seed data: %v", err)
	}

	// Setup minio
	bucketFound, err := bucket.BucketExists(prepareCtx, "public")
	if err != nil {
		log.Fatalf("Error checking MinIO bucket: %s\n", err)
	}

	if !bucketFound {
		err = bucket.MakeBucket(prepareCtx, "public", minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("Error creating MinIO bucket: %s\n", err)
		}

		policy := `{
			"Version":"2012-10-17",
			"Statement":[
			  {
				"Sid": "AddPerm",
				"Effect": "Allow",
				"Principal": "*",
				"Action":["s3:GetObject"],
				"Resource":["arn:aws:s3:::public/*"]
			  }
			]
		  }`

		err = bucket.SetBucketPolicy(prepareCtx, "public", policy)
		if err != nil {
			log.Fatalf("Error setting bucket policy: %s\n", err)
		}
	}

	code := m.Run()

	prepareCancel()

	fmt.Println("Cleaning up...")

	// Setup a context for cleaning up things
	cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), time.Second*60)

	err = cleanup(cleanupCtx)
	if err != nil {
		log.Fatalf("Failed on second cleanup: %v", err)
	}

	cleanupCancel()

	deps.DB.Close()

	os.Exit(code)
}

// prepareBuckets creates the buckets if they don't exist
func prepareBuckets(ctx context.Context, db influxdb2.Client, org string) error {
	bucketsAPI := deps.DB.BucketsAPI()
	organizationAPI := deps.DB.OrganizationsAPI()

	bucketNames := []string{
		common.BucketInputEvents,
		common.BucketSessionEvents,
		common.BucketFileEvents,
		common.BucketInputStatisticEvents,
		common.BucketWorkerStatus,
	}

	for _, bucket := range bucketNames {
		var b = bucket
		_, err := bucketsAPI.FindBucketByName(ctx, b)
		if err != nil && !strings.Contains(err.Error(), "not found") {
			return fmt.Errorf("finding bucket: %w", err)
		}

		if err != nil && strings.Contains(err.Error(), "not found") {
			orgDomain, err := organizationAPI.FindOrganizationByName(ctx, org)
			if err != nil {
				return fmt.Errorf("finding organization: %w", err)
			}

			_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, b)
			if err != nil && err.Error() != "conflict: bucket with name "+b+" already exists" {
				return fmt.Errorf("creating bucket: %w", err)
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
		return fmt.Errorf("finding organization: %w", err)
	}

	bucketNames := []string{
		common.BucketInputEvents,
		common.BucketSessionEvents,
		common.BucketFileEvents,
		common.BucketInputStatisticEvents,
		common.BucketWorkerStatus,
	}

	for _, bucket := range bucketNames {
		acquiredBucket, err := deps.DB.BucketsAPI().FindBucketByName(ctx, bucket)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				continue
			}

			return fmt.Errorf("finding bucket: %w", err)
		}

		err = deps.DB.BucketsAPI().DeleteBucket(ctx, acquiredBucket)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				continue
			}

			return fmt.Errorf("deleting bucket: %w", err)
		}

		_, err = deps.DB.BucketsAPI().CreateBucketWithName(ctx, currentOrganization, bucket)
		if err != nil {
			return fmt.Errorf("creating bucket: %w", err)
		}
	}

	return nil
}
