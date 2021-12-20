package main

import (
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"net/http"
	"os"
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
	mhost, ok := os.LookupEnv("MINIO_HOST")
	if !ok {
		log.Fatalln("MINIO_HOST envar missing")
	}
	maid, ok := os.LookupEnv("MINIO_ACCESS_ID")
	if !ok {
		log.Fatalln("MINIO_ACCESS_ID envar missing")
	}
	maidsex, ok := os.LookupEnv("MINIO_SECRET_KEY")
	if !ok {
		log.Fatalln("MINIO_SECRET_KEY envar missing")
	}

	influxConn := influxdb2.NewClient(influxHost, influxToken)

	minioConn, err := minio.New(mhost, &minio.Options{
		Creds:  credentials.NewStaticV4(maid, maidsex, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/sam-test", func(w http.ResponseWritter, r *http.Request) {
		return
	})

	http.HandleFunc("/sam-test", func(w http.ResponseWritter, r *http.Request) {

		return
	})

	http.ListenAndServe(":444", nil)
}
