package file_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"
	"worker/file"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var deps *file.Dependency
var globalID uuid.UUID

func TestMain(m *testing.M) {
	// Lookup environment variables
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

	db := influxdb2.NewClient(influxHost, influxToken)

	bucket, err := minio.New(
		minioHost,
		&minio.Options{
			Secure: false,
			Creds:  credentials.NewStaticV4(minioID, minioSecret, minioToken),
		},
	)
	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}

	deps = &file.Dependency{
		DB:                  db,
		DBOrganization:      influxOrg,
		Bucket:              bucket,
		BucketInputEvents:   "input_events",
		BucketSessionEvents: "session_events",
		Environment:         "testing",
	}

	// Check for bucket existence
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	err = prepareBuckets(ctx, deps.DB, influxOrg)
	if err != nil {
		log.Fatalf("Failed to prepare influxdb buckets: %v", err)
	}

	code := m.Run()

	err = cleanup(ctx)
	if err != nil {
		log.Fatalf("Failed to cleanup: %v", err)
	}

	deps.DB.Close()

	os.Exit(code)
}

func prepareBuckets(ctx context.Context, db influxdb2.Client, org string) error {
	bucketsAPI := deps.DB.BucketsAPI()
	_, err := bucketsAPI.FindBucketByName(ctx, deps.BucketInputEvents)
	if err != nil && err.Error() != "bucket '"+deps.BucketInputEvents+"' not found" {
		return fmt.Errorf("finding bucket: %v", err)
	}

	if err != nil && err.Error() == "bucket '"+deps.BucketInputEvents+"' not found" {
		organizationAPI := deps.DB.OrganizationsAPI()
		orgDomain, err := organizationAPI.FindOrganizationByName(ctx, org)
		if err != nil {
			return fmt.Errorf("finding organization: %v", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, deps.BucketInputEvents)
		if err != nil {
			return fmt.Errorf("creating bucket: %v", err)
		}
	}

	_, err = bucketsAPI.FindBucketByName(ctx, deps.BucketSessionEvents)
	if err != nil && err.Error() != "bucket '"+deps.BucketSessionEvents+"' not found" {
		return fmt.Errorf("finding bucket: %v", err)
	}

	if err != nil && err.Error() == "bucket '"+deps.BucketSessionEvents+"' not found" {
		organizationAPI := deps.DB.OrganizationsAPI()
		orgDomain, err := organizationAPI.FindOrganizationByName(ctx, org)
		if err != nil {
			return fmt.Errorf("finding organization: %v", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, deps.BucketSessionEvents)
		if err != nil {
			return fmt.Errorf("creating bucket: %v", err)
		}
	}

	return nil
}

func cleanup(ctx context.Context) error {
	// find current organization
	currentOrganization, err := deps.DB.OrganizationsAPI().FindOrganizationByName(ctx, deps.DBOrganization)
	if err != nil {
		return fmt.Errorf("finding organization: %v", err)
	}

	// find input_events bucket
	inputEventsBucket, err := deps.DB.BucketsAPI().FindBucketByName(ctx, deps.BucketInputEvents)
	if err != nil {
		return fmt.Errorf("finding bucket: %v", err)
	}

	// delete bucket data
	deleteAPI := deps.DB.DeleteAPI()

	inputEventMeasurements := []string{
		"keystroke",
		"mouse_down",
		"mouse_up",
		"mouse_moved",
		"mouse_scrolled",
		"window_sized",
	}
	for _, measurement := range inputEventMeasurements {
		err = deleteAPI.Delete(ctx, currentOrganization, inputEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\""+measurement+"\"")
		if err != nil {
			return fmt.Errorf("deleting bucket data: [%s] %v", measurement, err)
		}
	}

	// find input_events bucket
	sessionEventsBucket, err := deps.DB.BucketsAPI().FindBucketByName(ctx, deps.BucketSessionEvents)
	if err != nil {
		return fmt.Errorf("finding bucket: %v", err)
	}

	sessionEventMeasurements := []string{
		"code_test_attempt",
		"exam_forfeited",
		"exam_ended",
		"exam_started",
		"solution_rejected",
		"solution_accepted",
		"session_started",
		"personal_info_submitted",
		"locale_set",
		"exam_ide_reloaded",
		"deadline_passed",
		"before_exam_sam_submitted",
		"after_exam_sam_submitted",
	}
	for _, measurement := range sessionEventMeasurements {
		err = deleteAPI.Delete(ctx, currentOrganization, sessionEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\""+measurement+"\"")
		if err != nil {
			return fmt.Errorf("deleting bucket data: [%s] %v", measurement, err)
		}
	}

	return nil
}

func seedData(ctx context.Context) error {
	sessionWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketSessionEvents)
	//inputWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketInputEvents)

	// We generate two pieces of UUID, each of them have their own
	// specific use case.
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %v", err)
	}

	globalID = id

	eventStart := time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC)
	//eventEnd := time.Date(2020, 1, 2, 13, 0, 0, 0, time.UTC)

	var wg sync.WaitGroup
	// FIXME: thin this one out
	wg.Add(200)

	// Personal info
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/PersonalInfoSubmittedEvent.cs
	go func() {
		point := influxdb2.NewPoint(
			"personal_info_submitted",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"student_number":      "1202213133",
				"years_of_experience": 1,
				"hours_of_practice":   4,
				"familiar_languages":  "java,kotlin,swift",
			},
			eventStart,
		)
		err := sessionWriteAPI.WritePoint(ctx, point)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// SAM Test before Exam
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/BeforeExamSAMSubmittedEvent.cs
	go func() {
		point := influxdb2.NewPoint(
			"before_exam_sam_submitted",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"aroused_level": 2,
				"pleased_level": 5,
			},
			eventStart.Add(time.Minute*2),
		)
		err := sessionWriteAPI.WritePoint(ctx, point)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}
