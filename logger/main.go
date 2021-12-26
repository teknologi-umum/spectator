package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	_ "github.com/joho/godotenv/autoload"
)

type Dependency struct {
	DB  		influxdb2.Client
	Org 		string
	AccessToken string
}

func main() {
	influxURL, ok := os.LookupEnv("INFLUX_URL")
	if !ok {
		log.Fatal("INFLUX_URL environment variable must be set")
	}

	influxToken, ok := os.LookupEnv("INFLUX_TOKEN")
	if !ok {
		log.Fatal("INFLUX_TOKEN environment variable must be set")
	}

	influxOrganization, ok := os.LookupEnv("INFLUX_ORG")
	if !ok {
		log.Fatal("INFLUX_ORG environment variable must be set")
	}

	accessToken, ok := os.LookupEnv("ACCESS_TOKEN")
	if !ok {
		log.Fatal("ACCESS_TOKEN environment variable must be set")
	}

	db := influxdb2.NewClient(influxURL, influxToken)
	defer db.Close()

	deps := Dependency{
		DB:  db,
		Org: influxOrganization,
		AccessToken: accessToken,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			err := deps.ReadLogs(w, r)
			if err != nil {
				HandleError(err, w, r)
			}
			return
		case http.MethodPost:
			err := deps.InsertLog(w, r)
			if err != nil {
				HandleError(err, w, r)
			}
			return
		default:
			err := NotAllowed(w, r)
			if err != nil {
				HandleError(err, w, r)
			}
		}
	})
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			err := deps.Ping(w, r)
			if err != nil {
				HandleError(err, w, r)
			}
			return
		default:
			err := NotAllowed(w, r)
			if err != nil {
				HandleError(err, w, r)
			}
		}
	})

	server := http.Server{
		Handler:      mux,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 5,
		Addr:         ":5000",
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	
	go func(){
		log.Println("Starting server")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
		log.Println("Server closed")
	}()
	log.Printf("Logger service running on http://%s", server.Addr)

	<-done
	log.Printf("Server shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}
	log.Print("Server Exited Properly")
}
