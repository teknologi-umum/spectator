package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	loggerpb "video/logger_proto"
	pb "video/video_proto"

	"github.com/joncrlsn/dque"
)

// Dependency contains the dependency injection
// to be used on this package.
type Dependency struct {
	Ffmpeg         *Ffmpeg
	DB             influxdb2.Client
	Bucket         *minio.Client
	Logger         *Logger
	Queue          *dque.DQue
	Environment    string
	DBOrganization string
	pb.UnimplementedVideoServiceServer
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

	ffmpegClient, err := NewFfmpeg()
	if err != nil {
		log.Fatalln(err)
	}

	// Create embedded queue instance
	queue, err := dque.NewOrOpen("video-queue", ".", 50, VideoJobBuilder)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := queue.Close()
		if err != nil && !errors.Is(err, dque.ErrQueueClosed) {
			log.Printf("closing queue: %v", err)
		}
	}()

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
	defer func() {
		err := loggerConn.Close()
		if err != nil {
			log.Printf("closing logger service connection: %v", err)
		}
	}()

	loggerClient := loggerpb.NewLoggerClient(loggerConn)

	dependencies := &Dependency{
		Environment:    environment,
		Ffmpeg:         ffmpegClient,
		Bucket:         minioConn,
		DB:             influxConn,
		DBOrganization: influxOrg,
		Logger:         NewLogger(loggerClient, loggerToken, environment),
		Queue:          queue,
	}

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)

	// gRPC uses TCP connection.
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", portNumber))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Printf("closing listener: %v", err)
		}
	}()

	// Initialize gRPC server
	server := grpc.NewServer()

	// Register the service with the server, including injecting service dependencies.
	pb.RegisterVideoServiceServer(server, dependencies)

	go func() {
		log.Printf("Server listening at %s", listener.Addr().String())

		if err := server.Serve(listener); err != nil {
			log.Fatalln("Failed to serve:", err)
		}
	}()

	go func() {
		log.Printf("Starting video processing queue")
		for {
			iface, err := queue.DequeueBlock()
			if err != nil {
				log.Printf("error dequeuing block: %v", err)
				time.Sleep(time.Second)
				continue
			}

			block, ok := iface.(*VideoJob)
			if !ok {
				log.Printf("iface is not a type of VideoJob")
				time.Sleep(time.Second)
				continue
			}

			time.Sleep(time.Second * 30)
			dependencies.executeVideoJob(block.SessionId)
		}
	}()

	<-exitSignal

	log.Println("Shutting down server...")

	err = queue.Close()
	if err != nil {
		log.Printf("closing queue: %v", err)
	}

	server.GracefulStop()
}
