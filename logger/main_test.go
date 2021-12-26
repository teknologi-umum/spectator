package main_test

import (
	"os"
	"testing"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var db influxdb2.Client
var influxOrg string
var accessToken string

func TestMain(m *testing.M) {
	influxURL, ok := os.LookupEnv("INFLUX_URL")
	if !ok {
		influxURL = "http://localhost:8086"
	}

	influxToken, ok := os.LookupEnv("INFLUX_TOKEN")
	if !ok {
		influxToken = "H76G7mEgcyeV2ffM%E#Vd8U^eA6ZY8GH"
	}

	influxOrg, ok = os.LookupEnv("INFLUX_ORG")
	if !ok {
		influxOrg = "teknum_spectator"
	}

	accessToken, ok = os.LookupEnv("ACCESS_TOKEN")
	if !ok {
		accessToken = "testing"
	}

	db = influxdb2.NewClient(influxURL, influxToken)
	defer db.Close()

	os.Exit(m.Run())
}
