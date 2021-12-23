package main

import (
	"context"
	"encoding/json"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"net/http"
	"os"
	"runtime"
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
	defer influxConn.Close()

	minioConn, err := minio.New(mhost, &minio.Options{
		Creds:  credentials.NewStaticV4(maid, maidsex, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/fun-fact", func(w http.ResponseWriter, r *http.Request) {
		runtime.GOMAXPROCS(2)

		type Member struct {
			MemId string "json:`member_id`"
		}

		var x Member

		err := json.NewDecoder(r.Body).Decode(&x)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		queryAPI := influxConn.QueryAPI(influxOrg)

		wpm := make(chan int8)
		delRate := make(chan float32)
		attempt := make(chan int8)

		go func() {
			// TODO:  ini buat ngambil nganu, jangan lupa result
			_, err := queryAPI.Query(context.Background(), "from()")
			if err != nil {
				panic(err)
			}

		}()
		// aggregate WPM

		go func() {
			// TODO:  ini buat ngambil nganu, jangan lupa result
			_, err := queryAPI.Query(context.Background(), "from()")
			if err != nil {
				panic(err)
			}

		}()
		// aggregate Delete keys

		go func() {
			// TODO:  ini buat ngambil nganu, jangan lupa result
			_, err := queryAPI.Query(context.Background(), "from()")
			if err != nil {
				panic(err)
			}

		}()
		// question attempt

		var result = struct {
			Wpm     int8
			DelRate float32
			Attempt int8
		}{
			<-wpm,
			<-delRate,
			<-attempt,
		}

		res, _ := json.Marshal(result)

		w.Write(res)

		return
	})

	http.HandleFunc("/all-user-shit", func(w http.ResponseWriter, r *http.Request) {
		// TODO: nanti ditanya lagi masih males

		return
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.
		return
	})

	potnum, ok := os.LookupEnv("PORT")
	if ok {
		http.ListenAndServe(":"+potnum, nil)
	} else {
		http.ListenAndServe(":4444", nil)
	}
}
