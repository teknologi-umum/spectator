package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
	"worker/file"
	"worker/funfact"
	"worker/logger"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	loggerpb "worker/logger_proto"
	pb "worker/worker_proto"
)

// Dependency contains the dependency injection
// to be used on this package.
type Dependency struct {
	Environment    string
	DB             influxdb2.Client
	Bucket         *minio.Client
	DBOrganization string
	Logger         *logger.Logger
	LoggerToken    string
	Funfact        *funfact.Dependency
	File           *file.Dependency
	pb.UnimplementedWorkerServer
}

func main() {
	// Lookup environment variables
	influxToken, ok := os.LookupEnv("INFLUX_TOKEN")
	if !ok {
		log.Fatalln("INFLUX_TOKEN environment variable missing")
	}

	influxHost, ok := os.LookupEnv("INFLUX_HOST")
	if !ok {
		log.Fatalln("INFLUX_HOST environment variable missing")
	}

	influxOrg, ok := os.LookupEnv("INFLUX_ORG")
	if !ok {
		log.Fatalln("INFLUX_ORG environment variable missing")
	}

	minioHost, ok := os.LookupEnv("MINIO_HOST")
	if !ok {
		log.Fatalln("MINIO_HOST environment variable missing")
	}

	minioID, ok := os.LookupEnv("MINIO_ACCESS_ID")
	if !ok {
		log.Fatalln("MINIO_ACCESS_ID environment variable missing")
	}

	minioSecret, ok := os.LookupEnv("MINIO_SECRET_KEY")
	if !ok {
		log.Fatalln("MINIO_SECRET_KEY environment variable missing")
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

	minioToken, ok := os.LookupEnv("MINIO_TOKEN")
	if !ok {
		log.Fatalln("MINIO_TOKEN environment variable missing")
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

	// Dial the logger service
	loggerConn, err := grpc.Dial(
		loggerServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer loggerConn.Close()

	loggerClient := logger.New(
		loggerpb.NewLoggerClient(loggerConn),
		loggerToken,
		environment,
	)

	// Initialize dependency injection struct
	dependencies := &Dependency{
		DB:             influxConn,
		DBOrganization: influxOrg,
		Bucket:         minioConn,
		Logger:         loggerClient,
		LoggerToken:    loggerToken,
		Environment:    environment,
		Funfact: &funfact.Dependency{
			Environment:    environment,
			DB:             influxConn,
			DBOrganization: influxOrg,
			Logger:         loggerClient,
			LoggerToken:    loggerToken,
		},
		File: &file.Dependency{
			Environment:    environment,
			Bucket:         minioConn,
			DB:             influxConn,
			DBOrganization: influxOrg,
			Logger:         loggerClient,
			LoggerToken:    loggerToken,
		},
	}

	// Check for bucket existence
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	bucketFound, err := minioConn.BucketExists(ctx, "spectator")
	if err != nil {
		log.Fatalf("Error checking MinIO bucket: %s\n", err)
	}

	if !bucketFound {
		err = minioConn.MakeBucket(ctx, "spectator", minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("Error creating MinIObucket: %s\n", err)
		}
	}

	err = dependencies.prepareBuckets(ctx)
	if err != nil {
		log.Fatalf("Error preparing InfluxDB buckets: %s\n", err)
	}

	portNumber, ok := os.LookupEnv("PORT")
	if !ok {
		portNumber = "4444"
	}

	// gRPC uses TCP connection.
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", portNumber))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	defer listener.Close()

	// Initialize gRPC server
	server := grpc.NewServer()

	// Register the service with the server, including injecting service dependencies.
	pb.RegisterWorkerServer(server, dependencies)
	log.Printf("Server listening at %s", listener.Addr().String())

	if err := server.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
