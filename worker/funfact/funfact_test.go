package funfact_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
	"worker/funfact"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	deps      *funfact.Dependency
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

	db := influxdb2.NewClient(influxHost, influxToken)

	deps = &funfact.Dependency{
		DB:                  db,
		DBOrganization:      influxOrg,
		BucketInputEvents:   "input_events",
		BucketSessionEvents: "session_events",
		Environment:         "testing",
	}

	rand.Seed(time.Now().Unix())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)

	err := prepareBuckets(ctx)
	if err != nil {
		log.Fatalf("Error preparing influxdb buckets: %v", err)
	}

	// First cleanup to ensure that there are no data in the database
	err = cleanup(ctx)
	if err != nil {
		log.Fatalf("Error cleaning up buckets: %v", err)
	}

	err = seedData(ctx)
	if err != nil {
		log.Fatalf("Error seeding data: %v", err)
	}

	db.Close()

	cancel()

	code := m.Run()

	// Refresh context
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)

	// Second cleanup to ensure that there are no data in the database
	err = cleanup(ctx)
	if err != nil {
		log.Fatalf("Error cleaning up buckets: %v", err)
	}
	cancel()

	db.Close()

	os.Exit(code)
}

// prepareBuckets will check and create the buckets if they do not exist.
func prepareBuckets(ctx context.Context) error {
	bucketsAPI := deps.DB.BucketsAPI()
	organizationAPI := deps.DB.OrganizationsAPI()

	bucketNames := []string{deps.BucketInputEvents, deps.BucketSessionEvents, deps.BucketInputStatisticEvents}

	for _, bucket := range bucketNames {
		_, err := bucketsAPI.FindBucketByName(ctx, bucket)
		if err != nil && err.Error() != "bucket '"+bucket+"' not found" {
			return fmt.Errorf("finding bucket: %v", err)
		}

		if err != nil && err.Error() == "bucket '"+bucket+"' not found" {
			orgDomain, err := organizationAPI.FindOrganizationByName(ctx, deps.DBOrganization)
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

// seedData will fill the InfluxDB database with the generated data
// from this function. Why create separate one instead of seeding it on every
// test cases? Because we want to reduce HTTP write calls into the InfluxDB
func seedData(ctx context.Context) error {
	sessionWriteAPI := deps.DB.WriteAPI(deps.DBOrganization, deps.BucketSessionEvents)
	inputWriteAPI := deps.DB.WriteAPI(deps.DBOrganization, deps.BucketInputEvents)

	// We generate two pieces of UUID, each of them have their own
	// specific use case.
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %v", err)
	}

	globalID = id

	id, err = uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %v", err)
	}

	globalID2 = id

	var wg sync.WaitGroup
	wg.Add(11)

	// Random date between range
	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	// Seed coding test attempts
	go func() {
		for i := 0; i < 20; i++ {
			point := influxdb2.NewPoint(
				string(funfact.MeasurementSolutionAccepted),
				map[string]string{
					"session_id":  globalID.String(),
					"question_id": strconv.Itoa(rand.Intn(5)),
				},
				map[string]interface{}{
					"solution":               "console.log('Hello world!');",
					"language":               "javascript",
					"scratchpad":             "Lorem ipsum dolor sit amet",
					"serialized_test_result": "{\"stderr\":\"Hello world!\"}",
				},
				time.Unix(rand.Int63n(delta)+min, 0),
			)
			sessionWriteAPI.WritePoint(point)
		}

		for i := 0; i < 5; i++ {
			point := influxdb2.NewPoint(
				string(funfact.MeasurementSolutionRejected),
				map[string]string{
					"session_id":  globalID.String(),
					"question_id": strconv.Itoa(rand.Intn(5)),
				},
				map[string]interface{}{
					"solution":               "console.log('Hello world!');",
					"language":               "javascript",
					"scratchpad":             "Lorem ipsum dolor sit amet",
					"serialized_test_result": "{\"stderr\":\"Hello world!\"}",
				},
				time.Unix(rand.Int63n(delta)+min, 0),
			)
			sessionWriteAPI.WritePoint(point)
		}

		wg.Done()
	}()

	go func() {
		point := influxdb2.NewPoint(
			"exam_started",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"exam_id": "1",
			},
			time.Unix(min, 0),
		)

		sessionWriteAPI.WritePoint(point)
		wg.Done()
	}()

	go func() {
		point := influxdb2.NewPoint(
			"exam_started",
			map[string]string{
				"session_id": globalID2.String(),
			},
			map[string]interface{}{
				"exam_id": "1",
			},
			time.Unix(min, 0),
		)

		sessionWriteAPI.WritePoint(point)

		wg.Done()
	}()

	go func() {
		point := influxdb2.NewPoint(
			"exam_ended",
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"exam_id": "1",
			},
			time.Unix(max, 0),
		)

		sessionWriteAPI.WritePoint(point)

		wg.Done()
	}()

	go func() {
		point := influxdb2.NewPoint(
			"exam_forfeited",
			map[string]string{
				"session_id": globalID2.String(),
			},
			map[string]interface{}{
				"exam_id": "1",
			},
			time.Unix(max, 0),
		)

		sessionWriteAPI.WritePoint(point)

		wg.Done()
	}()

	// Keystroke events
	keystrokesNormal := []string{"a", "b", "c", "d", "e", "f"}
	keystrokesMisc := []string{"Space", "PageUp", "PageDown", "ArrowLeft", "ArrowRight", "ArrowUp", "ArrowDown"}
	keystrokesDelete := []string{"Backspace", "Delete"}

	// For the first user, he has a lot of keystrokes
	// During their session, they spent 5 minutes (the first for loop)
	// then they did a 70 keystroke-count of keystrokesNormal
	// 5 keystroke-count of keystrokesMisc and another 5 keystroke-count of keystrokesDelete
	temporaryDate := time.Unix(min, 0)
	for i := 0; i < 5; i++ {
		var anotherWg sync.WaitGroup
		anotherWg.Add(1)
		go func() {
			var childWg sync.WaitGroup
			childWg.Add(3)

			// Write 200 occurrence of normal keystrokes
			go func() {
				for j := 0; j < 200; j++ {
					point := influxdb2.NewPoint(
						"keystroke",
						map[string]string{
							"session_id": globalID.String(),
						},
						map[string]interface{}{
							"key_char":      keystrokesNormal[rand.Intn(len(keystrokesNormal)-1)],
							"unrelated_key": false,
						},
						time.Now(),
					)
					inputWriteAPI.WritePoint(point)
				}

				childWg.Done()
			}()

			// Write 100 occurrence of misc keystrokes
			go func() {
				for j := 0; j < 100; j++ {
					point := influxdb2.NewPoint(
						"keystroke",
						map[string]string{
							"session_id": globalID.String(),
						},
						map[string]interface{}{
							"key_char":      keystrokesMisc[rand.Intn(len(keystrokesMisc)-1)],
							"unrelated_key": false,
						},
						time.Now(),
					)

					inputWriteAPI.WritePoint(point)
				}

				childWg.Done()
			}()

			// Write 50 occurrence of deletion keystrokes
			go func() {
				for j := 0; j < 50; j++ {
					point := influxdb2.NewPoint(
						"keystroke",
						map[string]string{
							"session_id": globalID.String(),
						},
						map[string]interface{}{
							"key_char":      keystrokesDelete[rand.Intn(len(keystrokesDelete)-1)],
							"unrelated_key": false,
						},
						time.Now(),
					)

					inputWriteAPI.WritePoint(point)
				}

				childWg.Done()
			}()

			childWg.Wait()
			anotherWg.Done()
		}()
		anotherWg.Wait()
		temporaryDate = temporaryDate.Add(1 * time.Minute)
		wg.Done()
	}

	go func() {
		// temporaryDate := time.Unix(min, 0)
		for k := 0; k < 3; k++ {
			for i := 0; i < 100; i++ {
				point := influxdb2.NewPoint(
					"keystroke",
					map[string]string{
						"session_id": globalID2.String(),
					},
					map[string]interface{}{
						"key_char": keystrokesNormal[rand.Intn(len(keystrokesNormal)-1)],
					},
					time.Now(),
				)

				inputWriteAPI.WritePoint(point)
			}
		}

		wg.Done()
	}()

	wg.Wait()

	deps.DB.Close()

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
