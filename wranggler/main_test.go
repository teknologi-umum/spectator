package main_test

import (
	"context"
	"log"
	"net"
	"os"

	//"log"
	//"os"
	"testing"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	worker "worker"
	pb "worker/proto"
	// "github.com/minio/minio-go/v7"
	// "github.com/minio/minio-go/v7/pkg/credentials"
)

type Point struct {
	Type  string `json:"t"`
	Event string `json:"e"`
	Actor string `json:"a"`
	Value string `json:"v"`
}

type Submission struct {
	Type           string `json:"t"`
	Event          string `json:"e"`
	Actor          string `json:"a"`
	QuestionNumber string `json:"q"`
	Value          string `json:"v"`
}

type PersonalInfo struct {
	Type              string `json:"type"`
	SessionID         string `json:"session_id"`
	StudentNumber     string `json:"student_number"`
	HoursOfPractice   string `json:"hours_of_practice"`
	YearsOfExperience string `json:"years_of_experience"`
	FamiliarLanguages string `json:"familiar_languages"`
}

var db influxdb2.Client
var bucket *minio.Client
var dbOrganization string

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestMain(m *testing.M) {
	// Lookup environment variables
	influxToken, ok := os.LookupEnv("INFLUX_TOKEN")
	if !ok {
		influxToken = "H76G7mEgcyeV2ffM%E#Vd8U^eA6ZY8GH"
	}

	influxHost, ok := os.LookupEnv("INFLUX_HOST")
	if !ok {
		influxHost = "http://localhost:8086"
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

	minioToken, ok := os.LookupEnv("MINIO_TOKEN")
	if !ok {
		log.Fatalln("MINIO_TOKEN envar missing")
	}

	var err error

	db = influxdb2.NewClient(influxHost, influxToken)
	defer db.Close()

	dbOrganization = influxOrg

	bucket, err = minio.New(
		minioHost,
		&minio.Options{
			Secure: false,
			Creds:  credentials.NewStaticV4(minioID, minioSecret, minioToken),
		},
	)
	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterWorkerServer(s, &worker.Dependency{DB: db, Bucket: bucket, DBOrganization: influxOrg})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	os.Exit(m.Run())
}
