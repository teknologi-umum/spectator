package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "logger/proto"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

type Dependency struct {
	DB          influxdb2.Client
	Org         string
	AccessToken string
	pb.UnimplementedLoggerServer
}

type LogPayload struct {
	AccessToken string    `json:"access_token" msgpack:"access_token"`
	Data        []LogData `json:"data" msgpack:"data"`
}

type LogData struct {
	RequestID   string            `json:"request_id" msgpack:"request_id"`
	Application string            `json:"application" msgpack:"application"`
	Message     string            `json:"message" msgpack:"message"`
	Body        map[string]string `json:"body" msgpack:"body"`
	Level       string            `json:"level" msgpack:"level"`
	Environment string            `json:"environment" msgpack:"environment"`
	Language    string            `json:"language" msgpack:"language"`
	Timestamp   time.Time         `json:"timestamp" msgpack:"timestamp"`
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

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	db := influxdb2.NewClient(influxURL, influxToken)
	defer db.Close()

	deps := &Dependency{
		DB:          db,
		Org:         influxOrganization,
		AccessToken: accessToken,
	}

	// Prepare the log bucket
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := deps.PrepareBucket(ctx)
	if err != nil {
		log.Fatalf("preparing bucket: %v", err)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLoggerServer(grpcServer, deps)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		s := <-sigCh
		log.Println("Attempting graceful shutdown with SIGNAL:", s)
		grpcServer.GracefulStop()
		if err := listener.Close(); err != nil {
			log.Println("Failed to close listener:", err)
		}
	}()

	log.Println("gRPC server: Listening on port", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
