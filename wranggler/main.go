package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Dependency contains the dependency injection
// to be used on this package.
type Dependency struct {
	DB             influxdb2.Client
	Bucket         *minio.Client
	DBOrganization string
}

// Member contains the struct for member_id data
// that will be sent on the request body.
type Member struct {
	ID string `json:"member_id"`
}

func main() {
	// Lookup environment variables
	influxToken, ok := os.LookupEnv("INFLUX_TOKEN")
	if !ok {
		log.Fatalln("INFLUX_TOKEN envar missing")
	}

	influxHost, ok := os.LookupEnv("INFLUX_HOST")
	if !ok {
		log.Fatalln("INFLUX_HOST envar missing")
	}

	influxOrg, ok := os.LookupEnv("INFLUX_ORG")
	if !ok {
		log.Fatalln("INFLUX_ORG envar missing")
	}

	minioHost, ok := os.LookupEnv("MINIO_HOST")
	if !ok {
		log.Fatalln("MINIO_HOST envar missing")
	}

	minioID, ok := os.LookupEnv("MINIO_ACCESS_ID")
	if !ok {
		log.Fatalln("MINIO_ACCESS_ID envar missing")
	}

	minioSecret, ok := os.LookupEnv("MINIO_SECRET_KEY")
	if !ok {
		log.Fatalln("MINIO_SECRET_KEY envar missing")
	}

	// Create InfluxDB instance
	influxConn := influxdb2.NewClient(influxHost, influxToken)
	defer influxConn.Close()

	// Create Minio instance
	minioConn, err := minio.New(minioHost, &minio.Options{
		Creds:  credentials.NewStaticV4(minioID, minioSecret, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize dependency injection struct
	dependencies := &Dependency{
		DB:             influxConn,
		DBOrganization: influxOrg,
		Bucket:         minioConn,
	}

	// Create a new HTTP mux router with Chi
	r := chi.NewRouter()
	// Endpoint for healthchecks
	r.Get("/ping", dependencies.Ping)
	// Endpoint for generating fun fact about the user
	r.Post("/fun-fact", dependencies.FunFact)
	// Endpoint for generating file of a user
	r.Post("/generate", dependencies.GenerateFile)

	portNumber, ok := os.LookupEnv("PORT")
	if !ok {
		portNumber = "4444"
	}
	err = http.ListenAndServe(":"+portNumber, r)
	if err != nil {
		log.Fatal(err)
	}
}
