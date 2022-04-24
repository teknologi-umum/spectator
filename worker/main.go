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

	bucketFound, err := minioConn.BucketExists(ctx, "public")
	if err != nil {
		log.Fatalf("Error checking MinIO bucket: %s\n", err)
	}

	if !bucketFound {
		err = minioConn.MakeBucket(ctx, "public", minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("Error creating MinIO bucket: %s\n", err)
		}

		policy := `{
			"Version":"2012-10-17",
			"Statement":[
			  {
				"Sid": "AddPerm",
				"Effect": "Allow",
				"Principal": "*",
				"Action":["s3:GetObject"],
				"Resource":["arn:aws:s3:::public/*"]
			  }
			]
		  }`

		err = minioConn.SetBucketPolicy(ctx, "public", policy)
		if err != nil {
			log.Fatalf("Error setting bucket policy: %s\n", err)
		}
	}

	err = dependencies.prepareBuckets(ctx)
	if err != nil {
		log.Fatalf("Error preparing InfluxDB buckets: %s\n", err)
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
