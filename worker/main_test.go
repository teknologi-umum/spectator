package main_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	worker "worker"
	"worker/common"
	pb "worker/worker_proto"
)

var db influxdb2.Client
var bucket *minio.Client
var dbOrganization string

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

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

	var err error

	db = influxdb2.NewClient(influxHost, influxToken)

	dbOrganization = influxOrg

	bucket, err = minio.New(
		minioHost,
		&minio.Options{
			Secure: false,
			Creds:  credentials.NewStaticV4(minioID, minioSecret, minioToken),
		},
	)
	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}

	prepareCtx, prepareCancel := context.WithTimeout(context.Background(), time.Second*30)

	err = prepareBuckets(prepareCtx, db, influxOrg)
	if err != nil {
		log.Fatalf("Failed to prepare buckets: %v", err)
	}

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterWorkerServer(s, &worker.Dependency{DB: db, Bucket: bucket, DBOrganization: influxOrg})
	go func() {
		if err := s.Serve(lis); err != nil && !strings.Contains(err.Error(), "closed") {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	prepareCancel()

	code := m.Run()

	// It turns out that defer doesn't work
	// when combined with os.Exit()
	cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), time.Second*30)

	err = cleanup(cleanupCtx, db, dbOrganization)
	if err != nil {
		log.Printf("cleaning up: %v", err)
	}

	db.Close()

	err = lis.Close()
	if err != nil {
		log.Printf("closing listener: %v", err)
	}

	cleanupCancel()

	os.Exit(code)
}

func prepareBuckets(ctx context.Context, db influxdb2.Client, org string) error {
	bucketsAPI := db.BucketsAPI()
	organizationAPI := db.OrganizationsAPI()

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
func cleanup(ctx context.Context, db influxdb2.Client, org string) error {
	// find current organization
	currentOrganization, err := db.OrganizationsAPI().FindOrganizationByName(ctx, org)
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
		acquiredBucket, err := db.BucketsAPI().FindBucketByName(ctx, bucket)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				continue
			}

			return fmt.Errorf("finding bucket: %w", err)
		}

		err = db.BucketsAPI().DeleteBucket(ctx, acquiredBucket)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				continue
			}

			return fmt.Errorf("deleting bucket: %w", err)
		}

		_, err = db.BucketsAPI().CreateBucketWithName(ctx, currentOrganization, bucket)
		if err != nil {
			return fmt.Errorf("creating bucket: %w", err)
		}
	}

	return nil
}
