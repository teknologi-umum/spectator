package file_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
	"worker/file"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	deps      *file.Dependency
	globalID  uuid.UUID
	globalID2 uuid.UUID
)

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
		BucketFileEvents:    "file_results",
		Environment:         "testing",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	// Check for bucket existence
	err = prepareBuckets(ctx, deps.DB, influxOrg)
	if err != nil {
		log.Fatalf("Failed to prepare influxdb buckets: %v", err)
	}

	err = seedData(ctx)
	if err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	code := m.Run()

	fmt.Println("Cleaning up...")

	err = cleanup(ctx)
	if err != nil {
		log.Fatalf("Failed to cleanup: %v", err)
	}

	deps.DB.Close()

	os.Exit(code)
}

// prepareBuckets creates the buckets if they don't exist
func prepareBuckets(ctx context.Context, db influxdb2.Client, org string) error {
	bucketsAPI := deps.DB.BucketsAPI()
	organizationAPI := deps.DB.OrganizationsAPI()

	bucketNames := []string{deps.BucketInputEvents, deps.BucketSessionEvents, deps.BucketFileEvents}

	for _, bucket := range bucketNames {
		_, err := bucketsAPI.FindBucketByName(ctx, bucket)
		if err != nil && err.Error() != "bucket '"+bucket+"' not found" {
			return fmt.Errorf("finding bucket: %v", err)
		}

		if err != nil && err.Error() == "bucket '"+bucket+"' not found" {
			orgDomain, err := organizationAPI.FindOrganizationByName(ctx, org)
			if err != nil {
				return fmt.Errorf("finding organization: %v", err)
			}

			_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, bucket)
			if err != nil && err.Error() != "conflict: bucket with name "+bucket+" already exists" {
				return fmt.Errorf("creating bucket: %v", err)
			}
		}
	}

	return nil
}

// cleanup deletes the buckets' data
func cleanup(ctx context.Context) error {
	// find current organization
	currentOrganization, err := deps.DB.OrganizationsAPI().FindOrganizationByName(ctx, deps.DBOrganization)
	if err != nil {
		return fmt.Errorf("finding organization: %v", err)
	}

	// delete bucket data
	deleteAPI := deps.DB.DeleteAPI()

	// find input_events bucket
	inputEventsBucket, err := deps.DB.BucketsAPI().FindBucketByName(ctx, deps.BucketInputEvents)
	if err != nil {
		return fmt.Errorf("finding bucket: %v", err)
	}

	fileEventMeasurement := []string{
		"keystroke",
		"mouse_down",
		"mouse_up",
		"mouse_moved",
		"mouse_scrolled",
		"window_sized",
	}
	for _, measurement := range fileEventMeasurement {
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

	// find file_results bucket
	fileEventsBucket, err := deps.DB.BucketsAPI().FindBucketByName(ctx, deps.BucketFileEvents)
	if err != nil {
		return fmt.Errorf("finding bucket: %v", err)
	}

	for _, measurement := range fileEventMeasurement {
		err = deleteAPI.Delete(ctx, currentOrganization, fileEventsBucket, time.UnixMilli(0), time.Now(), "_measurement=\"exported_data\"")
		if err != nil {
			return fmt.Errorf("deleting bucket data: [%s] %v", measurement, err)
		}
	}

	// delete json/csv files
	pathJSON, err := filepath.Glob("./*_*.json")
	if err != nil {
		return fmt.Errorf("unexpected error: %v", err)
	}
	pathCSV, err := filepath.Glob("./*_*.csv")
	if err != nil {
		return fmt.Errorf("unexpected error: %v", err)
	}

	for _, path := range append(pathJSON, pathCSV...) {
		err = os.Remove(path)
		if err != nil {
			return fmt.Errorf("unexpected error: %v", err)
		}
	}

	return nil
}

// seedData seeds the database with test data
func seedData(ctx context.Context) error {
	sessionWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketSessionEvents)
	inputWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketInputEvents)

	// We generate two pieces of UUID, each of them have their own
	// specific use case.
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %v", err)
	}

	id2, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %v", err)
	}

	globalID = id
	globalID2 = id2

	eventStart := time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC)
	// eventEnd := time.Date(2020, 1, 2, 13, 0, 0, 0, time.UTC)

	var wg sync.WaitGroup
	// FIXME: correct this waitgroup number
	wg.Add(19)

	// Seeding session events for user with globalID
	go func() {
		// Personal info
		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/PersonalInfoSubmittedEvent.cs
		personalInfoPoint := influxdb2.NewPoint(
			"personal_info_submitted",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"student_number":      "1202213133",
				"hours_of_practice":   4,
				"years_of_experience": 3,
				"familiar_languages":  "java,kotlin,swift",
			},
			eventStart.Add(time.Minute+time.Second*time.Duration(i)),
		)

		beforeExamSAMPoint := influxdb2.NewPoint(
			"before_exam_sam_submitted",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"aroused_level": "2",
				"pleased_level": "5",
			},
			eventStart.Add(time.Minute*2+time.Second*time.Duration(i)),
		)

		afterExamSAMPoint := influxdb2.NewPoint(
			"after_exam_sam_submitted",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"aroused_level": "2",
				"pleased_level": "5",
			},
			eventStart.Add(time.Minute*3+time.Second*time.Duration(i)),
		)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamStartedEvent.cs

		examStartedPoint := influxdb2.NewPoint(
			"exam_started",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"_time": time.Now(),
			},
			eventStart.Add(time.Minute*4+time.Second*time.Duration(i)),
		)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamEndedEvent.cs
		examEndedPoint := influxdb2.NewPoint(
			"exam_ended",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"_time": time.Now(),
			},
			eventStart.Add(time.Minute*5+time.Second*time.Duration(i)),
		)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamIDEReloadedEvent.cs
		examIDEReloadedPoint := influxdb2.NewPoint(
			"exam_ide_reloaded",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"_time": time.Now(),
			},
			eventStart.Add(time.Minute*8+time.Second*time.Duration(i)),
		)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamPassedEvent.cs
		examPassedPoint := influxdb2.NewPoint(
			"exam_passed",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"_time": time.Now(),
			},
			eventStart.Add(time.Minute*7+time.Second*time.Duration(i)),
		)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/LocaleSetEvent.cs
		localeSetPoint := influxdb2.NewPoint(
			"locale_set",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"locale": "en-US",
			},
			eventStart.Add(time.Minute*9+time.Second*time.Duration(i)),
		)

		err := sessionWriteAPI.WritePoint(
			ctx,
			personalInfoPoint,
			beforeExamSAMPoint,
			afterExamSAMPoint,
			examStartedPoint,
			examEndedPoint,
			examIDEReloadedPoint,
			examPassedPoint,
			localeSetPoint,
		)
		if err != nil {
			log.Fatalf("Error writing session events point for globalID: %v", err)
		}
		wg.Done()
	}()

	// Seeding session events for user with globalID2
	go func() {
		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamForfeitedEvent.cs
		examForfeitedPoint := influxdb2.NewPoint(
			"exam_forfeited",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"_time": time.Now(),
			},
			eventStart.Add(time.Minute*6+time.Second*time.Duration(i)),
		)

		err := sessionWriteAPI.WritePoint(
			ctx,
			examForfeitedPoint,
		)
		if err != nil {
			log.Fatalf("Error writing session event points for globalID2: %v", err)
		}
		wg.Done()
	}()

	// Solution Accepted
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/SolutionAcceptedEvent.cs
	go func() {
		var points []*write.Point
		// TODO: lower the chance of this being generated
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"solution_accepted",
				map[string]string{
					"session_id": globalID.String(),
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*10+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Solution Rejected
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/SolutionRejectedEvent.cs
	go func() {
		var points []*write.Point
		// TODO: lower the chance of this being generated
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"solution_rejected",
				map[string]string{
					"session_id": globalID.String(),
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*11+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Deadline Passed
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/DeadlinePassedEvent.cs
	go func() {
		var points []*write.Point
		// TODO: lower the chance of this being generated
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"deadline_passed",
				map[string]string{
					"session_id": globalID.String(),
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*12+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Keystroke Event
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/InputDomain/KeystrokeEvent.cs
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"keystroke",
				map[string]string{
					"session_id": id.String(),
				},
				map[string]interface{}{
					// TODO: generate for each session id (there are 2 different uuid being generated, right?)
					"key_char":      "a",
					"key_code":      "65",
					"alt":           false,
					"control":       false,
					"shift":         false,
					"meta":          false,
					"unrelated_key": false,
				},
				eventStart.Add(time.Minute*13+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := inputWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Mouse Move Event
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"mouse_move",
				map[string]string{
					"session_id": id.String(),
				},
				map[string]interface{}{
					// TODO: generate for each session id (there are 2 different uuid being generated, right?)
					"direction":     "right",
					"x":             "20",
					"y":             "30",
					"window_width":  "100",
					"window_height": "200",
				},
				eventStart.Add(time.Minute*14+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := inputWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Mouse Down Event
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"mouse_down",
				map[string]string{
					"session_id": id.String(),
				},
				map[string]interface{}{
					// TODO: generate for each session id (there are 2 different uuid being generated, right?)
					"x":      "1",
					"y":      "2",
					"button": "0",
				},
				eventStart.Add(time.Minute*15+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := inputWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Mouse Up Event
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"mouse_up",
				map[string]string{
					"session_id": id.String(),
				},
				map[string]interface{}{
					// TODO: generate for each session id (there are 2 different uuid being generated, right?)
					"x":      "1",
					"y":      "2",
					"button": "0",
				},
				eventStart.Add(time.Minute*16+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := inputWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Mouse Scrolled Event
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"mouse_scrolled",
				map[string]string{
					"session_id": id.String(),
				},
				map[string]interface{}{
					"x": "1",
					"y": "2",
				},
				eventStart.Add(time.Minute*17+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := inputWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Window Resize Event
	go func() {
		var points []*write.Point
		// TODO: lower the chance of this being generated
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"window_resize",
				map[string]string{
					"session_id": id.String(),
				},
				map[string]interface{}{
					// TODO: generate for each session id (there are 2 different uuid being generated, right?)
					"width":  i,
					"height": i,
				},
				eventStart.Add(time.Minute*18+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := inputWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}
