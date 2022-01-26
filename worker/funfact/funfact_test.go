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
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var deps *funfact.Dependency
var globalID uuid.UUID
var globalID2 uuid.UUID

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
	_, err := bucketsAPI.FindBucketByName(ctx, deps.BucketInputEvents)
	if err != nil && err.Error() != "bucket '"+deps.BucketInputEvents+"' not found" {
		return fmt.Errorf("finding bucket: %v", err)
	}

	if err != nil && err.Error() == "bucket '"+deps.BucketInputEvents+"' not found" {
		organizationAPI := deps.DB.OrganizationsAPI()
		orgDomain, err := organizationAPI.FindOrganizationByName(ctx, deps.DBOrganization)
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
		orgDomain, err := organizationAPI.FindOrganizationByName(ctx, deps.DBOrganization)
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

// seedData will fill the InfluxDB database with the generated data
// from this function. Why create separate one instead of seeding it on every
// test cases? Because we want to reduce HTTP write calls into the InfluxDB
func seedData(ctx context.Context) error {
	sessionWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketSessionEvents)
	inputWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketInputEvents)

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
		var points []*write.Point
		for i := 0; i < 25; i++ {
			point := influxdb2.NewPoint(
				"code_test_attempt",
				map[string]string{
					"session_id":  globalID.String(),
					"question_id": strconv.Itoa(rand.Intn(5)),
				},
				map[string]interface{}{
					"code":     "console.log('Hello world!');",
					"language": "javascript",
				},
				time.Unix(rand.Int63n(delta)+min, 0),
			)
			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
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

		err := sessionWriteAPI.WritePoint(ctx, point)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
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

		err := sessionWriteAPI.WritePoint(ctx, point)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
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

		err := sessionWriteAPI.WritePoint(ctx, point)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
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

		err := sessionWriteAPI.WritePoint(ctx, point)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
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

			// Write 70 occurrence of normal keystrokes
			go func() {
				var points []*write.Point
				for j := 0; j < 70; j++ {
					point := influxdb2.NewPoint(
						"keystroke",
						map[string]string{
							"session_id": globalID.String(),
						},
						map[string]interface{}{
							"key_char":      keystrokesNormal[rand.Intn(len(keystrokesNormal)-1)],
							"unrelated_key": false,
						},
						temporaryDate,
					)
					points = append(points, point)
				}

				err := inputWriteAPI.WritePoint(ctx, points...)
				if err != nil {
					log.Fatalf("Error writing point: %v", err)
				}
				childWg.Done()
			}()

			// Write 5 occurrence of misc keystrokes
			go func() {
				var points []*write.Point
				for j := 0; j < 5; j++ {
					point := influxdb2.NewPoint(
						"keystroke",
						map[string]string{
							"session_id": globalID.String(),
						},
						map[string]interface{}{
							"key_char":      keystrokesMisc[rand.Intn(len(keystrokesMisc)-1)],
							"unrelated_key": false,
						},
						temporaryDate,
					)

					points = append(points, point)

				}
				err := inputWriteAPI.WritePoint(ctx, points...)
				if err != nil {
					log.Fatalf("Error writing point: %v", err)
				}
				childWg.Done()
			}()

			// Write 25 occurrence of deletion keystrokes
			go func() {
				var points []*write.Point
				for j := 0; j < 25; j++ {
					point := influxdb2.NewPoint(
						"keystroke",
						map[string]string{
							"session_id": globalID.String(),
						},
						map[string]interface{}{
							"key_char":      keystrokesDelete[rand.Intn(len(keystrokesDelete)-1)],
							"unrelated_key": false,
						},
						temporaryDate,
					)

					points = append(points, point)
				}

				err := inputWriteAPI.WritePoint(ctx, points...)
				if err != nil {
					log.Fatalf("Error writing point: %v", err)
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
		var points []*write.Point
		for i := 0; i < 50; i++ {
			point := influxdb2.NewPoint(
				"keystroke",
				map[string]string{
					"session_id": globalID2.String(),
				},
				map[string]interface{}{
					"key_char": keystrokesNormal[rand.Intn(len(keystrokesNormal)-1)],
				},
				time.Unix(rand.Int63n(delta)+min, 0),
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
