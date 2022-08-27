package funfact_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
	"worker/common"
	"worker/funfact"
	"worker/status"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
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
		DB:             db,
		DBOrganization: influxOrg,
		Environment:    "testing",
		Status: &status.Dependency{
			DB:             db,
			DBOrganization: influxOrg,
		},
	}

	rand.Seed(time.Now().Unix())

	// Setup a context for preparing things
	prepareCtx, prepareCancel := context.WithTimeout(context.Background(), time.Second*60)

	err := prepareBuckets(prepareCtx)
	if err != nil {
		log.Fatalf("Error preparing influxdb buckets: %v", err)
	}

	err = seedData(prepareCtx)
	if err != nil {
		log.Fatalf("Error seeding data: %v", err)
	}

	prepareCancel()

	code := m.Run()

	// Setup a context for cleaning up things
	cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), time.Second*60)

	// Second cleanup to ensure that there are no data in the database
	err = cleanup(cleanupCtx)
	if err != nil {
		log.Fatalf("Error cleaning up buckets: %v", err)
	}

	cleanupCancel()

	db.Close()

	os.Exit(code)
}

// prepareBuckets will check and create the buckets if they do not exist.
func prepareBuckets(ctx context.Context) error {
	bucketsAPI := deps.DB.BucketsAPI()
	organizationAPI := deps.DB.OrganizationsAPI()

	bucketNames := []string{common.BucketInputEvents, common.BucketSessionEvents, common.BucketInputStatisticEvents}

	for _, bucket := range bucketNames {
		_, err := bucketsAPI.FindBucketByName(ctx, bucket)
		if err != nil && !strings.Contains(err.Error(), "not found") {
			return fmt.Errorf("finding bucket: %w", err)
		}

		if err != nil && strings.Contains(err.Error(), "not found") {
			orgDomain, err := organizationAPI.FindOrganizationByName(ctx, deps.DBOrganization)
			if err != nil {
				return fmt.Errorf("finding organization: %w", err)
			}

			_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, bucket)
			if err != nil && err.Error() != "conflict: bucket with name "+bucket+" already exists" {
				return fmt.Errorf("creating bucket: %w", err)
			}
		}
	}

	return nil
}

// seedData will fill the InfluxDB database with the generated data
// from this function. Why create separate one instead of seeding it on every
// test cases? Because we want to reduce HTTP write calls into the InfluxDB
func seedData(ctx context.Context) error {
	sessionWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, common.BucketSessionEvents)
	inputWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, common.BucketInputEvents)

	// We generate two pieces of UUID, each of them have their own
	// specific use case.
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %w", err)
	}

	globalID = id

	id, err = uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %w", err)
	}

	globalID2 = id

	var wg sync.WaitGroup
	wg.Add(8)

	// Random date between range
	min := time.Date(2019, 5, 2, 1, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2019, 5, 2, 1, 4, 0, 0, time.UTC).Unix()
	delta := max - min

	// Seed coding test attempts
	go func() {
		var points []*write.Point
		for i := 0; i < 20; i++ {
			point := influxdb2.NewPoint(
				common.MeasurementSolutionAccepted,
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

			points = append(points, point)
		}

		for i := 0; i < 5; i++ {
			point := influxdb2.NewPoint(
				common.MeasurementSolutionRejected,
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

			points = append(points, point)
		}

		err := sessionWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("failed to write point: %v", err)
		}

		wg.Done()
	}()

	go func() {
		point1 := influxdb2.NewPoint(
			common.MeasurementExamStarted,
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"exam_id": "1",
			},
			time.Unix(min, 0),
		)

		point2 := influxdb2.NewPoint(
			common.MeasurementExamStarted,
			map[string]string{
				"session_id": globalID2.String(),
			},
			map[string]interface{}{
				"exam_id": "1",
			},
			time.Unix(min, 0),
		)

		point3 := influxdb2.NewPoint(
			common.MeasurementExamEnded,
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"exam_id": "1",
			},
			time.Unix(max, 0),
		)

		point4 := influxdb2.NewPoint(
			common.MeasurementExamForfeited,
			map[string]string{
				"session_id": globalID2.String(),
			},
			map[string]interface{}{
				"exam_id": "1",
			},
			time.Unix(max, 0),
		)

		point5 := influxdb2.NewPoint(
			common.MeasurementPersonalInfoSubmitted,
			map[string]string{
				"session_id": globalID.String(),
			},
			map[string]interface{}{
				"student_number":      "1202213133",
				"hours_of_practice":   4,
				"years_of_experience": 3,
				"familiar_languages":  "java,kotlin,swift",
			},
			time.Unix(min, 0),
		)

		err := sessionWriteAPI.WritePoint(ctx, point1, point2, point3, point4, point5)
		if err != nil {
			log.Fatalf("failed to write point: %v", err)
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
		go func() {
			var childWg sync.WaitGroup
			childWg.Add(3)

			// Write 200 occurrence of normal keystrokes
			go func() {
				var points []*write.Point
				for j := 0; j < 200; j++ {
					point := influxdb2.NewPoint(
						common.MeasurementKeystroke,
						map[string]string{
							"session_id": globalID.String(),
						},
						map[string]interface{}{
							"key_char":      keystrokesNormal[rand.Intn(len(keystrokesNormal)-1)],
							"unrelated_key": false,
						},
						time.Now(),
					)

					points = append(points, point)
				}

				err := inputWriteAPI.WritePoint(ctx, points...)
				if err != nil {
					log.Fatalf("failed to write point: %v", err)
				}
				childWg.Done()
			}()

			// Write 100 occurrence of misc keystrokes
			go func() {
				var points []*write.Point
				for j := 0; j < 100; j++ {
					point := influxdb2.NewPoint(
						common.MeasurementKeystroke,
						map[string]string{
							"session_id": globalID.String(),
						},
						map[string]interface{}{
							"key_char":      keystrokesMisc[rand.Intn(len(keystrokesMisc)-1)],
							"unrelated_key": false,
						},
						time.Now(),
					)

					points = append(points, point)
				}

				err := inputWriteAPI.WritePoint(ctx, points...)
				if err != nil {
					log.Fatalf("failed to write point: %v", err)
				}
				childWg.Done()
			}()

			// Write 50 occurrence of deletion keystrokes
			go func() {
				var points []*write.Point
				for j := 0; j < 50; j++ {
					point := influxdb2.NewPoint(
						common.MeasurementKeystroke,
						map[string]string{
							"session_id": globalID.String(),
						},
						map[string]interface{}{
							"key_char":      keystrokesDelete[rand.Intn(len(keystrokesDelete)-1)],
							"unrelated_key": false,
						},
						time.Now(),
					)

					points = append(points, point)
				}
				err := inputWriteAPI.WritePoint(ctx, points...)
				if err != nil {
					log.Fatalf("failed to write point: %v", err)
				}

				childWg.Done()
			}()

			childWg.Wait()
			wg.Done()
		}()

		temporaryDate = temporaryDate.Add(1 * time.Minute)
	}

	go func() {
		// temporaryDate := time.Unix(min, 0)
		for k := 0; k < 3; k++ {
			var points []*write.Point
			for i := 0; i < 100; i++ {
				point := influxdb2.NewPoint(
					common.MeasurementKeystroke,
					map[string]string{
						"session_id": globalID2.String(),
					},
					map[string]interface{}{
						"key_char": keystrokesNormal[rand.Intn(len(keystrokesNormal)-1)],
					},
					time.Now(),
				)

				points = append(points, point)
			}

			err := inputWriteAPI.WritePoint(ctx, points...)
			if err != nil {
				log.Fatalf("failed to write point: %v", err)
			}
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
		return fmt.Errorf("finding organization: %w", err)
	}

	for _, bucket := range []string{common.BucketInputEvents, common.BucketSessionEvents, common.BucketInputStatisticEvents} {
		acquiredBucket, err := deps.DB.BucketsAPI().FindBucketByName(ctx, bucket)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				continue
			}

			return fmt.Errorf("finding bucket: %w", err)
		}

		err = deps.DB.BucketsAPI().DeleteBucket(ctx, acquiredBucket)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				continue
			}

			return fmt.Errorf("deleting bucket: %w", err)
		}

		_, err = deps.DB.BucketsAPI().CreateBucketWithName(ctx, currentOrganization, bucket)
		if err != nil {
			return fmt.Errorf("creating bucket: %w", err)
		}
	}

	return nil
}
