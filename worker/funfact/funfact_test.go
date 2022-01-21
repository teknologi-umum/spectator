package funfact_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
	"worker/funfact"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var deps *funfact.Dependency
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

	db = influxdb2.NewClient(influxHost, influxToken)

	deps = &funfact.Dependency{
		DB:                  db,
		DBOrganization:      influxOrg,
		BucketInputEvents:   "input_events",
		BucketSessionEvents: "session_events",
		Environment:         "testing",
	}

	rand.Seed(time.Now().Unix())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	err := prepareBuckets(ctx, deps.DB, influxOrg)
	if err != nil {
		log.Fatalf("Error preparing influxdb buckets: %v", err)
	}

	code := m.Run()

	db.Close()

	os.Exit(code)
}

func prepareBuckets(ctx context.Context, db influxdb2.Client, org string) error {
	bucketsAPI := db.BucketsAPI()
	_, err := bucketsAPI.FindBucketByName(ctx, deps.BucketInputEvents)
	if err != nil && err.Error() != "bucket '"+deps.BucketInputEvents+"' not found" {
		return fmt.Errorf("finding bucket: %v", err)
	}

	if err != nil && err.Error() == "bucket '"+deps.BucketInputEvents+"' not found" {
		organizationAPI := db.OrganizationsAPI()
		orgDomain, err := organizationAPI.FindOrganizationByName(ctx, org)
		if err != nil {
			return fmt.Errorf("finding organization: %v", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, deps.BucketInputEvents)
		if err != nil {
			return fmt.Errorf("creating bucket: %v", err)
		}
	}

	_, err = bucketsAPI.FindBucketByName(ctx, deps.BucketSessionEvents)
	if err != nil && err.Error() != "bucket '"+deps.BucketSessionEvents+"' not found" {
		return fmt.Errorf("finding bucket: %v", err)
	}

	if err != nil && err.Error() == "bucket '"+deps.BucketSessionEvents+"' not found" {
		organizationAPI := db.OrganizationsAPI()
		orgDomain, err := organizationAPI.FindOrganizationByName(ctx, org)
		if err != nil {
			return fmt.Errorf("finding organization: %v", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, deps.BucketSessionEvents)
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
