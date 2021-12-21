package main

import (
	"encoding/json"
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

	http.HandleFunc("/fun-fact", func(w http.ResponseWritter, r *http.Request) {

		type Member struct {
			MemId string "json:`member_id`"
		}

		var x Member

		err := json.NewDecoder(r.Body).Decode(&x)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		wpm := make(chan int8)
		delRate := make(chan float32)

		go func() {

		}()
		// aggregate WPM

		go func() {

		}()
		// aggregate Delete keys

		var result = struct {
			Wpm     int8
			DelRate float32
		}{
			<-wpm,
			<-delRate,
		}

		res, _ := json.Marshall(result)

		w.Write(res)

		return
	})

	http.HandleFunc("/all-user-shit", func(w http.ResponseWritter, r *http.Request) {

		return
	})

	http.HandleFunc("/ping", func(w http.ResponseWritter, r *http.Request) {

		return
	})

	potnum, ok := os.LookupEnv("PORT")
	if ok {
		http.ListenAndServe(":"+potnum, nil)
	} else {
		http.ListenAndServe(":4444", nil)
	}
}
