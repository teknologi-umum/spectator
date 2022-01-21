package main_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	worker "worker"
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
		minioID = "diPj59zJzm2kwUZxcg5QRAUtpbVx5Uxd"
	}

	minioSecret, ok := os.LookupEnv("MINIO_SECRET_KEY")
	if !ok {
		minioSecret = "xLxBHSp2vAdX2TJSy6EptamrNk5ZXzXo"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	err = prepareBuckets(ctx, db, influxOrg)
	if err != nil {
		log.Fatalf("Failed to prepare buckets: %v", err)
	}

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterWorkerServer(s, &worker.Dependency{DB: db, Bucket: bucket, DBOrganization: influxOrg})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	code := m.Run()

	// It turns out that defer doesn't work
	// when combined with os.Exit()
	cancel()
	db.Close()

	os.Exit(code)
}

func prepareBuckets(ctx context.Context, db influxdb2.Client, org string) error {
	bucketsAPI := db.BucketsAPI()
	_, err := bucketsAPI.FindBucketByName(ctx, worker.BucketInputEvents)
	if err != nil && err.Error() != "bucket '"+worker.BucketInputEvents+"' not found" {
		return fmt.Errorf("finding bucket: %v", err)
	}

	if err != nil && err.Error() == "bucket '"+worker.BucketInputEvents+"' not found" {
		organizationAPI := db.OrganizationsAPI()
		orgDomain, err := organizationAPI.FindOrganizationByName(ctx, org)
		if err != nil {
			return fmt.Errorf("finding organization: %v", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, worker.BucketInputEvents)
		if err != nil {
			return fmt.Errorf("creating bucket: %v", err)
		}
	}

	_, err = bucketsAPI.FindBucketByName(ctx, worker.BucketSessionEvents)
	if err != nil && err.Error() != "bucket '"+worker.BucketSessionEvents+"' not found" {
		return fmt.Errorf("finding bucket: %v", err)
	}

	if err != nil && err.Error() == "bucket '"+worker.BucketSessionEvents+"' not found" {
		organizationAPI := db.OrganizationsAPI()
		orgDomain, err := organizationAPI.FindOrganizationByName(ctx, org)
		if err != nil {
			return fmt.Errorf("finding organization: %v", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, worker.BucketSessionEvents)
		if err != nil {
			return fmt.Errorf("creating bucket: %v", err)
		}
	}

	return nil
}

func cleanup() {
	// create new context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// find current organization
	currentOrganization, err := db.OrganizationsAPI().FindOrganizationByName(ctx, dbOrganization)
	if err != nil {
		log.Fatalf("finding organization: %v", err)
	}

	// find input_events bucket
	inputEventsBucket, err := db.BucketsAPI().FindBucketByName(ctx, worker.BucketInputEvents)
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
	sessionEventsBucket, err := db.BucketsAPI().FindBucketByName(ctx, worker.BucketSessionEvents)
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
