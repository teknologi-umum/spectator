package file_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
	"worker/common"
	"worker/file"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/sync/errgroup"
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

	code := m.Run()

	prepareCancel()

	fmt.Println("Cleaning up...")

	// Setup a context for cleaning up things
	cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), time.Second*60)

	err = cleanup(cleanupCtx)
	if err != nil {
		log.Fatalf("Failed to cleanup: %v", err)
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
	}

	g, gctx := errgroup.WithContext(ctx)

	for _, bucket := range bucketNames {
		var b = bucket
		g.Go(func() error {
			_, err := bucketsAPI.FindBucketByName(gctx, b)
			if err != nil && err.Error() != "bucket '"+b+"' not found" {
				return fmt.Errorf("finding bucket: %w", err)
			}

			if err != nil && err.Error() == "bucket '"+b+"' not found" {
				orgDomain, err := organizationAPI.FindOrganizationByName(gctx, org)
				if err != nil {
					return fmt.Errorf("finding organization: %w", err)
				}

				_, err = bucketsAPI.CreateBucketWithName(gctx, orgDomain, b)
				if err != nil && err.Error() != "conflict: bucket with name "+b+" already exists" {
					return fmt.Errorf("creating bucket: %w", err)
				}
			}

			return nil
		})
	}

	return g.Wait()
}

// cleanup deletes the buckets' data
func cleanup(ctx context.Context) error {
	// find current organization
	currentOrganization, err := deps.DB.OrganizationsAPI().FindOrganizationByName(ctx, deps.DBOrganization)
	if err != nil {
		return fmt.Errorf("finding organization: %w", err)
	}

	for _, bucket := range []string{common.BucketFileEvents, common.BucketInputEvents, common.BucketInputStatisticEvents, common.BucketSessionEvents} {
		acquiredBucket, err := deps.DB.BucketsAPI().FindBucketByName(ctx, bucket)
		if err != nil {
			return fmt.Errorf("finding bucket: %w", err)
		}

		err = deps.DB.BucketsAPI().DeleteBucket(ctx, acquiredBucket)
		if err != nil {
			return fmt.Errorf("deleting bucket: %w", err)
		}

		_, err = deps.DB.BucketsAPI().CreateBucketWithName(ctx, currentOrganization, bucket)
		if err != nil {
			return fmt.Errorf("creating bucket: %w", err)
		}
	}

	return nil
}
