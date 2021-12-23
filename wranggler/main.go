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

func main() {
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

	influxConn := influxdb2.NewClient(influxHost, influxToken)
	defer influxConn.Close()
	minioConn, err := minio.New(minioHost, &minio.Options{
		Creds:  credentials.NewStaticV4(minioID, minioSecret, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	r := chi.NewRouter()
	r.Get("/ping", Ping)
	r.Post("/fun-fact", FunFact)
	r.Post("/generate", GenerateFile)

	potnum, ok := os.LookupEnv("PORT")
	if ok {
		http.ListenAndServe(":"+potnum, r)
	} else {
		http.ListenAndServe(":4444", r)
	}
}
