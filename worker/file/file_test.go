package file_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
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
	deps     *file.Dependency
	globalID uuid.UUID
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

	// Check for bucket existence
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

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

	// err = cleanup(ctx)
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

	_, err := bucketsAPI.FindBucketByName(ctx, deps.BucketInputEvents)
	if err != nil && err.Error() != "bucket '"+deps.BucketInputEvents+"' not found" {
		return fmt.Errorf("finding bucket: %v", err)
	}

	if err != nil && err.Error() == "bucket '"+deps.BucketInputEvents+"' not found" {
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
		orgDomain, err := organizationAPI.FindOrganizationByName(ctx, org)
		if err != nil {
			return fmt.Errorf("finding organization: %v", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, deps.BucketSessionEvents)
		if err != nil {
			return fmt.Errorf("creating bucket: %v", err)
		}
	}

	_, err = bucketsAPI.FindBucketByName(ctx, deps.BucketFileEvents)
	if err != nil && err.Error() != "bucket '"+deps.BucketFileEvents+"' not found" {
		return fmt.Errorf("finding bucket: %v", err)
	}

	if err != nil && err.Error() == "bucket '"+deps.BucketFileEvents+"' not found" {
		orgDomain, err := organizationAPI.FindOrganizationByName(ctx, org)
		if err != nil {
			return fmt.Errorf("finding organization: %v", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, deps.BucketFileEvents)
		if err != nil {
			return fmt.Errorf("creating bucket: %v", err)
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

// seedData seeds the database with test data
func seedData(ctx context.Context) error {
	sessionWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketSessionEvents)
	inputWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketInputEvents)
	fileWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketFileEvents)

	// We generate two pieces of UUID, each of them have their own
	// specific use case.
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %v", err)
	}

	globalID = id

	eventStart := time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC)
	// eventEnd := time.Date(2020, 1, 2, 13, 0, 0, 0, time.UTC)

	var wg sync.WaitGroup
	wg.Add(19)

	// Personal info
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/PersonalInfoSubmittedEvent.cs
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
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
			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// SAM Test before Exam
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/BeforeExamSAMSubmittedEvent.cs
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
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
			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// SAM Test after Exam
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/AfterExamSAMSubmittedEvent.cs
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
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
			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Exam Started
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamStartedEvent.cs
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"exam_started",
				map[string]string{
					"session_id": globalID.String(),
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*4+time.Second*time.Duration(i)),
			)
			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Exam Ended
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamEndedEvent.cs
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"exam_ended",
				map[string]string{
					"session_id": globalID.String(),
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*5+time.Second*time.Duration(i)),
			)
			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Exam Forfeited
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamForfeitedEvent.cs
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"exam_forfeited",
				map[string]string{
					"session_id": globalID.String(),
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*6+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Exam Passed
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamPassedEvent.cs
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"exam_passed",
				map[string]string{
					"session_id": globalID.String(),
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*7+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Exam IDE Reloaded
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamIDEReloadedEvent.cs
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"exam_ide_reloaded",
				map[string]string{
					"session_id": globalID.String(),
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*8+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Locale Set
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/LocaleSetEvent.cs
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"locale_set",
				map[string]string{
					"session_id": globalID.String(),
				},
				map[string]interface{}{
					"locale": "en-US",
				},
				eventStart.Add(time.Minute*9+time.Second*time.Duration(i)),
			)

			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Solution Accepted
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/SolutionAcceptedEvent.cs
	go func() {
		var points []*write.Point
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
					"direction":     "right",
					"x_position":    "20",
					"y_position":    "30",
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
					"x":     "1",
					"y":     "2",
					"delta": "3",
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
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"window_resize",
				map[string]string{
					"session_id": id.String(),
				},
				map[string]interface{}{
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

	// Test Result
	go func() {
		var points []*write.Point
		for i := 0; i < 50; i++ {
			studentNumber := fmt.Sprintf("%08d", rand.Intn(100000000))
			for _, x := range []string{"keystroke", "mouse_click", "mouse_move", "personal_info", "sam_test"} {
				point := influxdb2.NewPoint(
					"exported_data",
					map[string]string{
						"session_id":     globalID.String(),
						"student_number": studentNumber,
					},
					map[string]interface{}{
						"file_csv_url":  "/public/" + studentNumber + "_" + x + ".csv",
						"file_json_url": "/public/" + studentNumber + "_" + x + ".json",
					},
					time.Now(),
				)
				points = append(points, point)

				csvHandle, err := os.Create("./results/" + studentNumber + "_" + x + ".csv")
				if err != nil {
					log.Fatalf("creating a file: %v", err)
					return
				}
				defer csvHandle.Close()

				jsonHandle, err := os.Create("./results/" + studentNumber + "_" + x + ".json")
				if err != nil {
					log.Fatalf("creating a file: %v", err)
					return
				}
				defer jsonHandle.Close()

				_, err = csvHandle.Write([]byte(x))
				if err != nil {
					log.Fatalf("writing to a file: %v", err)
					return
				}

				_, err = jsonHandle.Write([]byte(x))
				if err != nil {
					log.Fatalf("writing to a file: %v", err)
					return
				}

				err = csvHandle.Sync()
				if err != nil {
					log.Fatalf("syncing a file: %v", err)
					return
				}

				err = jsonHandle.Sync()
				if err != nil {
					log.Fatalf("syncing a file: %v", err)
					return
				}

				_, err = csvHandle.Stat()
				if err != nil {
					log.Fatalf("getting file stat: %v", err)
					return
				}

				_, err = jsonHandle.Stat()
				if err != nil {
					log.Fatalf("getting file stat: %v", err)
					return
				}
			}
		}

		err := fileWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}
