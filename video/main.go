package main

import (
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Dependency contains the dependency injection
// to be used on this package.
type Dependency struct {
	Environment string
}

func main() {
	// Lookup environment variables
	err := loadEnvironment()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

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

	loggerServerAddr, ok := os.LookupEnv("LOGGER_SERVER_ADDRESS")
	if !ok {
		log.Fatalln("LOGGER_SERVER_ADDRESS environment variable missing")
	}

	loggerToken, ok := os.LookupEnv("LOGGER_TOKEN")
	if !ok {
		log.Fatalln("LOGGER_TOKEN environment variable missing")
	}

	environment, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		environment = "DEVELOPMENT"
	}

	portNumber, ok := os.LookupEnv("PORT")
	if !ok {
		portNumber = "3000"
	}

	// Create InfluxDB instance
	influxConn := influxdb2.NewClient(influxHost, influxToken)
	defer influxConn.Close()

	// Create Minio instance
	minioConn, err := minio.New(minioHost, &minio.Options{
		Creds: credentials.NewStaticV4(minioID, minioSecret, minioToken),
	})
	if err != nil {
		log.Fatalln(err)
	}
}
