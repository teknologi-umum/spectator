package file_test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

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
	wg.Add(11)

	// Seeding session events for each user
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
		go func(sessionID string) {
			// Personal info
			// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/PersonalInfoSubmittedEvent.cs
			personalInfoPoint := influxdb2.NewPoint(
				"personal_info_submitted",
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
				"before_exam_sam_submitted",
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"aroused_level": "2",
					"pleased_level": "5",
				},
				eventStart.Add(time.Minute*2),
			)

			afterExamSAMPoint := influxdb2.NewPoint(
				"after_exam_sam_submitted",
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"aroused_level": "2",
					"pleased_level": "5",
				},
				eventStart.Add(time.Minute*3),
			)

			// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamStartedEvent.cs

			examStartedPoint := influxdb2.NewPoint(
				"exam_started",
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"_time": time.Now(),
				},
				eventStart.Add(time.Minute*4),
			)

			// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamEndedEvent.cs
			examEndedPoint := influxdb2.NewPoint(
				"exam_ended",
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
				"exam_forfeited",
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
				"exam_ide_reloaded",
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
				"exam_passed",
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
				"locale_set",
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
					"solution_accepted",
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
					"solution_rejected",
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
				"deadline_passed",
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
					"keystroke",
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
					"mouse_moved",
					map[string]string{
						"session_id": sessionID,
					},
					map[string]interface{}{
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
					"mouse_down",
					map[string]string{
						"session_id": sessionID,
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
					"mouse_up",
					map[string]string{
						"session_id": sessionID,
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
					"mouse_scrolled",
					map[string]string{
						"session_id": sessionID,
					},
					map[string]interface{}{
						"x": "1",
						"y": "2",
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
					"window_resize",
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

	wg.Wait()
	return nil
}
