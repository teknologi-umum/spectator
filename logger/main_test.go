package main_test

import (
	"context"
	"log"
	logger "logger"
	"os"
	"testing"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var db influxdb2.Client
var influxOrg string
var accessToken string

func TestMain(m *testing.M) {
	influxURL, ok := os.LookupEnv("INFLUX_URL")
	if !ok {
		influxURL = "http://localhost:8086"
		err := os.Setenv("INFLUX_URL", influxURL)
		if err != nil {
			log.Fatalf("error setting INFLUX_URL environment variable: %v", err)
		}
	}

	influxToken, ok := os.LookupEnv("INFLUX_TOKEN")
	if !ok {
		influxToken = "H76G7mEgcyeV2ffM%E#Vd8U^eA6ZY8GH"
		err := os.Setenv("INFLUX_TOKEN", influxToken)
		if err != nil {
			log.Fatalf("error setting INFLUX_TOKEN environment variable: %v", err)
		}
	}

	influxOrg, ok = os.LookupEnv("INFLUX_ORG")
	if !ok {
		influxOrg = "teknum_spectator"
		err := os.Setenv("INFLUX_ORG", influxOrg)
		if err != nil {
			log.Fatalf("error setting INFLUX_ORG environment variable: %v", err)
		}
	}

	accessToken, ok = os.LookupEnv("ACCESS_TOKEN")
	if !ok {
		accessToken = "testing"
		err := os.Setenv("ACCESS_TOKEN", accessToken)
		if err != nil {
			log.Fatalf("error setting ACCESS_TOKEN environment variable: %v", err)
		}
	}

	err := os.Setenv("TZ", "UTC")
	if err != nil {
		log.Fatalf("setting TZ environment variable to UTC: %v", err)
	}

	db = influxdb2.NewClient(influxURL, influxToken)
	defer db.Close()

	deps := logger.Dependency{
		DB:          db,
		Org:         influxOrg,
		AccessToken: accessToken,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = deps.PrepareBucket(ctx)
	if err != nil {
		log.Fatalf("failed preparing bucket: %v", err)
	}

	os.Exit(m.Run())
}

func cleanup() {
	// create new context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// find current organization
	currentOrganization, err := db.OrganizationsAPI().FindOrganizationByName(ctx, influxOrg)
	if err != nil {
		log.Fatalf("finding organization: %v", err)
	}

	// find current bucket
	currentBucket, err := db.BucketsAPI().FindBucketByName(ctx, "log")
	if err != nil {
		log.Fatalf("finding bucket: %v", err)
	}

	// delete bucket data
	deleteAPI := db.DeleteAPI()
	err = deleteAPI.Delete(ctx, currentOrganization, currentBucket, time.UnixMilli(0), time.Now(), "")
	if err != nil {
		log.Fatalf("deleting bucket data: %v", err)
	}
}
