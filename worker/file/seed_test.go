package file_test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"worker/common"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

// seedData seeds the database with test data
func seedData(ctx context.Context) error {
	sessionWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, common.BucketSessionEvents)
	inputWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, common.BucketInputEvents)
	statisticWriteAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, common.BucketInputStatisticEvents)

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
	wg.Add(12)

	// Seeding session events for each user
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
		go func(sessionID string) {
			// Personal info
			// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/PersonalInfoSubmittedEvent.cs
			personalInfoPoint := influxdb2.NewPoint(
				common.MeasurementPersonalInfoSubmitted,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"student_number":      "1202213133",
					"hours_of_practice":   4,
					"years_of_experience": 3,
					"familiar_languages":  "java,kotlin,swift",
				},
				eventStart.Add(time.Minute),
			)

			beforeExamSAMPoint := influxdb2.NewPoint(
				common.MeasurementBeforeExamSAMSubmitted,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"aroused_level": 2,
					"pleased_level": 5,
				},
				eventStart.Add(time.Minute*2),
			)

			afterExamSAMPoint := influxdb2.NewPoint(
				common.MeasurementAfterExamSAMSubmitted,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"aroused_level": 2,
					"pleased_level": 5,
				},
				eventStart.Add(time.Minute*3),
			)

			// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamStartedEvent.cs
			examStartedPoint := influxdb2.NewPoint(
				common.MeasurementExamStarted,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"question_numbers": "1,2,3,4,5",
					"deadline":         eventStart.Add(time.Minute * 180).Unix(),
				},
				eventStart.Add(time.Minute*4),
			)

			// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamEndedEvent.cs
			examEndedPoint := influxdb2.NewPoint(
				common.MeasurementExamEnded,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*5),
			)

			// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamForfeited.cs
			examForfeitedPoint := influxdb2.NewPoint(
				common.MeasurementExamForfeited,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*6),
			)

			// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamIDEReloadedEvent.cs
			examIDEReloadedPoint := influxdb2.NewPoint(
				common.MeasurementExamIDEReloaded,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*8),
			)

			// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamPassedEvent.cs
			examPassedPoint := influxdb2.NewPoint(
				common.MeasurementExamPassed,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*7),
			)

			// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/LocaleSetEvent.cs
			localeSetPoint := influxdb2.NewPoint(
				common.MeasurementLocaleSet,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"locale": "en-US",
				},
				eventStart.Add(time.Minute*9),
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
				examForfeitedPoint,
				localeSetPoint,
			)
			if err != nil {
				log.Fatalf("Error writing session events point for %s: %v", sessionID, err)
			}
			wg.Done()
		}(sessionID)
	}

	// Solution Accepted
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/SolutionAcceptedEvent.cs
	go func() {
		var points []*write.Point
		for _, sessionID := range []string{globalID.String(), globalID2.String()} {
			for i := 0; i < 5; i++ {
				point := influxdb2.NewPoint(
					common.MeasurementSolutionAccepted,
					map[string]string{
						"session_id": sessionID,
					},
					map[string]interface{}{
						"_time": time.Now(),
					},
					eventStart.Add(time.Minute*10+time.Second*time.Duration(i)),
				)

				points = append(points, point)
			}
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
		for _, sessionID := range []string{globalID.String(), globalID2.String()} {
			for i := 0; i < 5; i++ {
				point := influxdb2.NewPoint(
					common.MeasurementSolutionRejected,
					map[string]string{
						"session_id": sessionID,
					},
					map[string]interface{}{
						"_time": time.Now(),
					},
					eventStart.Add(time.Minute*11+time.Second*time.Duration(i)),
				)

				points = append(points, point)
			}
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
		for _, sessionID := range []string{globalID.String(), globalID2.String()} {
			point := influxdb2.NewPoint(
				common.MeasurementDeadlinePassed,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*12),
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
		for _, sessionID := range []string{globalID.String(), globalID2.String()} {
			for i := 0; i < 50; i++ {
				point := influxdb2.NewPoint(
					common.MeasurementKeystroke,
					map[string]string{
						"session_id": sessionID,
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
		for _, sessionID := range []string{globalID.String(), globalID2.String()} {
			for i := 0; i < 50; i++ {
				point := influxdb2.NewPoint(
					common.MeasurementMouseMoved,
					map[string]string{
						"session_id": sessionID,
					},
					map[string]interface{}{
						"direction": "right",
						"x":         20,
						"y":         30,
					},
					eventStart.Add(time.Minute*14+time.Second*time.Duration(i)),
				)

				points = append(points, point)
			}
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
		for _, sessionID := range []string{globalID.String(), globalID2.String()} {
			for i := 0; i < 50; i++ {
				point := influxdb2.NewPoint(
					common.MeasurementMouseDown,
					map[string]string{
						"session_id": sessionID,
					},
					map[string]interface{}{
						"x":      6,
						"y":      5,
						"button": int64(common.MouseButtonRight),
					},
					eventStart.Add(time.Minute*15+time.Second*time.Duration(i)),
				)

				points = append(points, point)
			}
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
		for _, sessionID := range []string{globalID.String(), globalID2.String()} {
			for i := 0; i < 50; i++ {
				point := influxdb2.NewPoint(
					common.MeasurementMouseUp,
					map[string]string{
						"session_id": sessionID,
					},
					map[string]interface{}{
						"x":      1,
						"y":      2,
						"button": int64(common.MouseButtonMiddle),
					},
					eventStart.Add(time.Minute*16+time.Second*time.Duration(i)),
				)

				points = append(points, point)
			}
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
		for _, sessionID := range []string{globalID.String(), globalID2.String()} {
			for i := 0; i < 50; i++ {
				point := influxdb2.NewPoint(
					common.MeasurementMouseScrolled,
					map[string]string{
						"session_id": sessionID,
					},
					map[string]interface{}{
						"x": 1,
						"y": 2,
					},
					eventStart.Add(time.Minute*17+time.Second*time.Duration(i)),
				)

				points = append(points, point)
			}
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
		for _, sessionID := range []string{globalID.String(), globalID2.String()} {
			for i := 0; i < 4; i++ {
				point := influxdb2.NewPoint(
					common.MeasurementWindowSized,
					map[string]string{
						"session_id": sessionID,
					},
					map[string]interface{}{
						"width":  i,
						"height": i,
					},
					eventStart.Add(time.Minute*18+time.Second*time.Duration(i)),
				)

				points = append(points, point)
			}
		}

		err := inputWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	// Funfact Event
	go func() {
		var points []*write.Point

		for _, sessionID := range []string{globalID.String(), globalID2.String()} {
			point := influxdb2.NewPoint(
				common.MeasurementFunfactProjection,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"words_per_minute":    60,
					"deletion_rate":       0.7,
					"submission_attempts": 30,
				},
				time.Now(),
			)

			points = append(points, point)
		}

		err := statisticWriteAPI.WritePoint(ctx, points...)
		if err != nil {
			log.Fatalf("Error writing point: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}
