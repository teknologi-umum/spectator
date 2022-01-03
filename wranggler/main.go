package main

import (
	"fmt"
	"log"
	"net"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"

	pb "worker/proto"
)

// Dependency contains the dependency injection
// to be used on this package.
type Dependency struct {
	DB             influxdb2.Client
	Bucket         *minio.Client
	DBOrganization string
	pb.UnimplementedWorkerServer
}

const (
	// BucketInputEvents is the bucket name for storing
	// keystroke events, window events, and mouse events.
	BucketInputEvents = "input_events"
	// BucketSessionEvents is the bucket name for storing
	// the session events, including their personal information.
	BucketSessionEvents = "session_events"
)

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
		Creds:  credentials.NewStaticV4(minioID, minioSecret, "spectator"),
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

	portNumber, ok := os.LookupEnv("PORT")
	if !ok {
		portNumber = "4444"
	}

	// gRPC uses TCP connection.
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "localhost", portNumber))
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
