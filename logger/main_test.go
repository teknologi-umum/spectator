package main_test

import (
	"context"
	"log"
	logger "logger"
	pb "logger/proto"
	"net"
	"os"
	"testing"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var db influxdb2.Client
var influxOrg string
var accessToken string

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

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
		influxToken = "nMfrRYVcTyqFwDARAdqB92Ywj6GNMgPEd"
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

	deps := &logger.Dependency{
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

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterLoggerServer(s, deps)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

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

	measurements := []string{"ERROR", "WARNING", "INFO", "DEBUG", "CRITICAL"}
	for _, measurement := range measurements {
		err = deleteAPI.Delete(ctx, currentOrganization, currentBucket, time.UnixMilli(0), time.Now(), "_measurement=\""+measurement+"\"")
		if err != nil {
			log.Fatalf("deleting bucket data: [%s] %v", measurement, err)
		}
	}
}
