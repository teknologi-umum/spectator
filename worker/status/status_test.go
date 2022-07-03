package status_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
	"worker/common"
	"worker/status"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var deps *status.Dependency

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

	db := influxdb2.NewClient(influxHost, influxToken)

	deps = &status.Dependency{
		DB:             db,
		DBOrganization: influxOrg,
	}

	// Setup a context for preparing things
	prepareCtx, prepareCancel := context.WithTimeout(context.Background(), time.Second*60)

	err := prepareBuckets(prepareCtx)
	if err != nil {
		log.Fatalf("Error preparing influxdb buckets: %v", err)
	}

	// First cleanup to ensure that there are no data in the database
	err = cleanup(prepareCtx)
	if err != nil {
		log.Fatalf("Error cleaning up buckets: %v", err)
	}

	db.Close()

	prepareCancel()

	code := m.Run()

	// Setup a context for cleaning up things
	cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), time.Second*60)

	// Second cleanup to ensure that there are no data in the database
	err = cleanup(cleanupCtx)
	if err != nil {
		log.Fatalf("Error cleaning up buckets: %v", err)
	}

	cleanupCancel()

	db.Close()

	os.Exit(code)
}

// prepareBuckets will check and create the buckets if they do not exist.
func prepareBuckets(ctx context.Context) error {
	bucketsAPI := deps.DB.BucketsAPI()
	organizationAPI := deps.DB.OrganizationsAPI()

	bucketNames := []string{common.BucketWorkerStatus}

	for _, bucket := range bucketNames {
		_, err := bucketsAPI.FindBucketByName(ctx, bucket)
		if err != nil && err.Error() != "bucket '"+bucket+"' not found" {
			return fmt.Errorf("finding bucket: %w", err)
		}

		if err != nil && err.Error() == "bucket '"+bucket+"' not found" {
			orgDomain, err := organizationAPI.FindOrganizationByName(ctx, deps.DBOrganization)
			if err != nil {
				return fmt.Errorf("finding organization: %w", err)
			}

			_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, bucket)
			if err != nil && err.Error() != "conflict: bucket with name "+bucket+" already exists" {
				return fmt.Errorf("creating bucket: %w", err)
			}
		}
	}

	return nil
}

func cleanup(ctx context.Context) error {
	// find current organization
	currentOrganization, err := deps.DB.OrganizationsAPI().FindOrganizationByName(ctx, deps.DBOrganization)
	if err != nil {
		return fmt.Errorf("finding organization: %w", err)
	}

	for _, bucket := range []string{common.BucketWorkerStatus} {
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

func TestAppendState(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	randomSessionID := uuid.New()

	t.Run("Normal", func(t *testing.T) {
		err := deps.AppendState(ctx, randomSessionID, "funfact", status.StatePending)
		if err != nil {
			t.Errorf("append state: %v", err)
		}
	})

	t.Run("EmptyParameter", func(t *testing.T) {
		err := deps.AppendState(ctx, randomSessionID, "", status.StateFailed)
		if err == nil {
			t.Errorf("expecting an error, got nil instead")
		}

		if !errors.Is(err, status.ErrEmptyFieldParameter) {
			t.Errorf("expecting an error of ErrEmptyFieldParameter, instead got: %v", err)
		}
	})
}
