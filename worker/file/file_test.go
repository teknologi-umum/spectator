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

	g, gctx := errgroup.WithContext(prepareCtx)
	
	g.Go(func() error {
		// Check for InfluxDB buckets existence
		err = prepareBuckets(gctx, deps.DB, influxOrg)
		if err != nil {
			return fmt.Errorf("failed to prepare influxdb buckets: %v", err)
		}
	
		err = seedData(gctx)
		if err != nil {
			return fmt.Errorf("failed to seed data: %v", err)
		}

		return nil
	})

	g.Go(func() error {
		// Check for MinIO bucket existence
		bucketFound, err := bucket.BucketExists(gctx, "spectator")
		if err != nil {
			return fmt.Errorf("error checking MinIO bucket: %s\n", err)
		}

		if !bucketFound {
			err = bucket.MakeBucket(gctx, "spectator", minio.MakeBucketOptions{})
			if err != nil {
				return fmt.Errorf("error creating MinIObucket: %s\n", err)
			}
		}

		return nil
	})

	err = g.Wait()
	if err != nil {
		log.Fatalf("Failed to prepare test: %v", err)
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
				return fmt.Errorf("finding bucket: %v", err)
			}
	
			if err != nil && err.Error() == "bucket '"+b+"' not found" {
				orgDomain, err := organizationAPI.FindOrganizationByName(gctx, org)
				if err != nil {
					return fmt.Errorf("finding organization: %v", err)
				}
	
				_, err = bucketsAPI.CreateBucketWithName(gctx, orgDomain, b)
				if err != nil && err.Error() != "conflict: bucket with name "+b+" already exists" {
					return fmt.Errorf("creating bucket: %v", err)
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
		return fmt.Errorf("finding organization: %v", err)
	}

	// delete bucket data
	deleteAPI := deps.DB.DeleteAPI()

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		// find input_events bucket
		inputEventsBucket, err := deps.DB.BucketsAPI().FindBucketByName(gctx, common.BucketInputEvents)
		if err != nil {
			return fmt.Errorf("finding bucket: %v", err)
		}

		inputEventMeasurement := []string{
			common.MeasurementKeystroke,
			common.MeasurementMouseDown,
			common.MeasurementMouseUp,
			common.MeasurementMouseMoved,
			common.MeasurementMouseScrolled,
			common.MeasurementWindowSized,
		}
		for _, measurement := range inputEventMeasurement {
			err = deleteAPI.Delete(gctx, currentOrganization, inputEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\""+measurement+"\"")
			if err != nil {
				return fmt.Errorf("deleting bucket data: [%s] %v", measurement, err)
			}
		}

		return nil
	})

	g.Go(func() error {
		// find input_events bucket
		sessionEventsBucket, err := deps.DB.BucketsAPI().FindBucketByName(gctx, common.BucketSessionEvents)
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

		// More speed hack, we create a child errgroup
		c, cctx := errgroup.WithContext(gctx)
		c.Go(func() error {
			for _, measurement := range sessionEventMeasurements[:len(sessionEventMeasurements)/2] {
				err = deleteAPI.Delete(cctx, currentOrganization, sessionEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\""+measurement+"\"")
				if err != nil {
					return fmt.Errorf("deleting bucket data: [%s] %v", measurement, err)
				}
			}
			return nil
		})

		c.Go(func() error {
			for _, measurement := range sessionEventMeasurements[len(sessionEventMeasurements)/2:] {
				err = deleteAPI.Delete(cctx, currentOrganization, sessionEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\""+measurement+"\"")
				if err != nil {
					return fmt.Errorf("deleting bucket data: [%s] %v", measurement, err)
				}
			}
			return nil
		})

		return c.Wait()
	})

	g.Go(func() error {
		// find statistics bucket
		statisticBucket, err := deps.DB.BucketsAPI().FindBucketByName(gctx, common.BucketInputStatisticEvents)
		if err != nil {
			return fmt.Errorf("finding bucket: %v", err)
		}

		statisticEventMeasurements := []string{
			common.MeasurementFunfactProjection,
		}

		for _, measurement := range statisticEventMeasurements {
			err = deleteAPI.Delete(gctx, currentOrganization, statisticBucket, time.UnixMilli(0), time.Now(), "_measurement=\""+measurement+"\"")
			if err != nil {
				return fmt.Errorf("deleting bucket data: [%s] %v", measurement, err)
			}
		}

		return nil
	})

	g.Go(func() error {
		// find file_results bucket
		fileEventsBucket, err := deps.DB.BucketsAPI().FindBucketByName(gctx, common.BucketFileEvents)
		if err != nil {
			return fmt.Errorf("finding bucket: %v", err)
		}

		fileEventMeasurement := []string{common.MeasurementExportedData}

		for _, measurement := range fileEventMeasurement {
			err = deleteAPI.Delete(gctx, currentOrganization, fileEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\"exported_data\"")
			if err != nil {
				return fmt.Errorf("deleting bucket data: [%s] %v", measurement, err)
			}
		}

		return nil
	})

	return g.Wait()
}
