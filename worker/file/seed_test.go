package file_test

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"worker/common"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// seedData seeds the database with test data
func seedData(ctx context.Context) error {
	sessionWriteAPI := deps.DB.WriteAPI(deps.DBOrganization, common.BucketSessionEvents)
	inputWriteAPI := deps.DB.WriteAPI(deps.DBOrganization, common.BucketInputEvents)
	statisticWriteAPI := deps.DB.WriteAPI(deps.DBOrganization, common.BucketInputStatisticEvents)

	// We generate two pieces of UUID, each of them have their own
	// specific use case.
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %w", err)
	}

	id2, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate uuid: %w", err)
	}

	globalID = id
	globalID2 = id2

	eventStart := time.Date(2022, 1, 2, 12, 0, 0, 0, time.UTC)
	// eventEnd := time.Date(2020, 1, 2, 13, 0, 0, 0, time.UTC)

	// Seeding session events for each user
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
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
				"wallet_number":       "0812131415",
				"wallet_type":         "Gopay",
			},
			eventStart.Add(time.Minute),
		)
		sessionWriteAPI.WritePoint(personalInfoPoint)

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
		sessionWriteAPI.WritePoint(beforeExamSAMPoint)

		afterExamSAMPoint := influxdb2.NewPoint(
			common.MeasurementAfterExamSAMSubmitted,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"aroused_level": 2,
				"pleased_level": 5,
			},
			eventStart.Add(time.Minute*30),
		)
		sessionWriteAPI.WritePoint(afterExamSAMPoint)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamStartedEvent.cs
		examStartedPoint := influxdb2.NewPoint(
			common.MeasurementExamStarted,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"question_numbers": "1,2,3,4,5",
				"deadline":         eventStart.Add(time.Minute * 180).UnixNano(),
			},
			eventStart.Add(time.Minute*4),
		)
		sessionWriteAPI.WritePoint(examStartedPoint)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamEndedEvent.cs
		examEndedPoint := influxdb2.NewPoint(
			common.MeasurementExamEnded,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"arbitrary": "arbitrary",
			},
			eventStart.Add(time.Minute*5),
		)
		sessionWriteAPI.WritePoint(examEndedPoint)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamForfeited.cs
		examForfeitedPoint := influxdb2.NewPoint(
			common.MeasurementExamForfeited,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"arbitrary": "arbitrary",
			},
			eventStart.Add(time.Minute*6),
		)
		sessionWriteAPI.WritePoint(examForfeitedPoint)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamIDEReloadedEvent.cs
		examIDEReloadedPoint := influxdb2.NewPoint(
			common.MeasurementExamIDEReloaded,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"arbitrary": "arbitrary",
			},
			eventStart.Add(time.Minute*8),
		)
		sessionWriteAPI.WritePoint(examIDEReloadedPoint)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamPassedEvent.cs
		examPassedPoint := influxdb2.NewPoint(
			common.MeasurementExamPassed,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"arbitrary": "arbitrary",
			},
			eventStart.Add(time.Minute*7),
		)
		sessionWriteAPI.WritePoint(examPassedPoint)

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
		sessionWriteAPI.WritePoint(localeSetPoint)
	}

	// Because I want to make the sessionWriteAPI's buffer full, we should create more dummy users.
	for i := 0; i < 10; i++ {
		sessionID := uuid.NewString()

		var studentNumber = make([]byte, 16)
		rand.Read(studentNumber)

		// Personal info
		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/PersonalInfoSubmittedEvent.cs
		personalInfoPoint := influxdb2.NewPoint(
			common.MeasurementPersonalInfoSubmitted,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"student_number":      string(studentNumber),
				"hours_of_practice":   rand.Intn(23) + 1,
				"years_of_experience": rand.Intn(10),
				"familiar_languages":  "java,kotlin,swift",
				"wallet_number":       strconv.Itoa(rand.Int()),
				"wallet_type":         "Gopay",
			},
			eventStart.Add(time.Minute),
		)
		sessionWriteAPI.WritePoint(personalInfoPoint)

		beforeExamSAMPoint := influxdb2.NewPoint(
			common.MeasurementBeforeExamSAMSubmitted,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"aroused_level": rand.Intn(10),
				"pleased_level": rand.Intn(10),
			},
			eventStart.Add(time.Minute*2),
		)
		sessionWriteAPI.WritePoint(beforeExamSAMPoint)

		afterExamSAMPoint := influxdb2.NewPoint(
			common.MeasurementAfterExamSAMSubmitted,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"aroused_level": rand.Intn(10),
				"pleased_level": rand.Intn(10),
			},
			eventStart.Add(time.Minute*30),
		)
		sessionWriteAPI.WritePoint(afterExamSAMPoint)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamStartedEvent.cs
		examStartedPoint := influxdb2.NewPoint(
			common.MeasurementExamStarted,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"question_numbers": "1,2,3,4,5",
				"deadline":         eventStart.Add(time.Minute * 180).UnixNano(),
			},
			eventStart.Add(time.Minute*4),
		)
		sessionWriteAPI.WritePoint(examStartedPoint)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamEndedEvent.cs
		examEndedPoint := influxdb2.NewPoint(
			common.MeasurementExamEnded,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"arbitrary": "arbitrary",
			},
			eventStart.Add(time.Minute*5),
		)
		sessionWriteAPI.WritePoint(examEndedPoint)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamForfeited.cs
		examForfeitedPoint := influxdb2.NewPoint(
			common.MeasurementExamForfeited,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"arbitrary": "arbitrary",
			},
			eventStart.Add(time.Minute*6),
		)
		sessionWriteAPI.WritePoint(examForfeitedPoint)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamIDEReloadedEvent.cs
		examIDEReloadedPoint := influxdb2.NewPoint(
			common.MeasurementExamIDEReloaded,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"arbitrary": "arbitrary",
			},
			eventStart.Add(time.Minute*8),
		)
		sessionWriteAPI.WritePoint(examIDEReloadedPoint)

		// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/ExamPassedEvent.cs
		examPassedPoint := influxdb2.NewPoint(
			common.MeasurementExamPassed,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"arbitrary": "arbitrary",
			},
			eventStart.Add(time.Minute*7),
		)
		sessionWriteAPI.WritePoint(examPassedPoint)

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
		sessionWriteAPI.WritePoint(localeSetPoint)
	}

	// Solution Accepted
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/SolutionAcceptedEvent.cs
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
		for i := 0; i < 5; i++ {
			point := influxdb2.NewPoint(
				common.MeasurementSolutionAccepted,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"question_number":         i + 1,
					"language":                "PHP",
					"solution":                "echo 'Hello World!';",
					"scratchpad":              "Lorem ipsum dolot sit amet",
					"serialized_test_results": "Hello World!",
				},
				eventStart.Add(time.Minute*10+time.Second*time.Duration(i)),
			)

			sessionWriteAPI.WritePoint(point)
		}
	}

	// Solution Rejected
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/SolutionRejectedEvent.cs
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
		for i := 0; i < 20; i++ {
			point := influxdb2.NewPoint(
				common.MeasurementSolutionRejected,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"question_number":         i + 1,
					"language":                "PHP",
					"solution":                "echo 'Hello World!';",
					"scratchpad":              "Lorem ipsum dolot sit amet",
					"serialized_test_results": "Hello World!",
				},
				eventStart.Add(time.Minute*11+time.Second*time.Duration(i)),
			)

			sessionWriteAPI.WritePoint(point)
		}
	}

	// Deadline Passed
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/SessionDomain/DeadlinePassedEvent.cs
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
		point := influxdb2.NewPoint(
			common.MeasurementDeadlinePassed,
			map[string]string{
				"session_id": sessionID,
			},
			map[string]interface{}{
				"arbitrary": "arbitrary",
			},
			eventStart.Add(time.Minute*12),
		)

		sessionWriteAPI.WritePoint(point)
	}

	// Keystroke Event
	// https://github.com/teknologi-umum/spectator/blob/master/backend/Spectator.DomainEvents/InputDomain/KeystrokeEvent.cs
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
		var questionNumber = 1
		for i := 0; i < 500; i++ {
			point := influxdb2.NewPoint(
				common.MeasurementKeystroke,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"question_number": questionNumber,
					"key_char":        "a",
					"key_code":        "65",
					"alt":             false,
					"control":         false,
					"shift":           false,
					"meta":            false,
					"unrelated_key":   false,
				},
				eventStart.Add((time.Minute*13)+(250*time.Millisecond*time.Duration(i))),
			)

			inputWriteAPI.WritePoint(point)

			if questionNumber == 6 {
				questionNumber = 1
			} else {
				questionNumber++
			}
		}
	}

	// Mouse Move Event
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
		var questionNumber = 1
		for i := 0; i < 500; i++ {
			point := influxdb2.NewPoint(
				common.MeasurementMouseMoved,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"question_number": questionNumber,
					"direction":       "right",
					"x":               rand.Intn(1920),
					"y":               rand.Intn(1080),
				},
				eventStart.Add((time.Minute*13)+(250*time.Millisecond*time.Duration(i))),
			)

			inputWriteAPI.WritePoint(point)

			if questionNumber == 6 {
				questionNumber = 1
			} else {
				questionNumber++
			}
		}
	}

	// Mouse Down Event
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
		var questionNumber int = 1
		for i := 0; i < 100; i++ {
			point := influxdb2.NewPoint(
				common.MeasurementMouseDown,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"question_number": questionNumber,
					"x":               rand.Intn(1920),
					"y":               rand.Intn(1080),
					"button":          int64(common.MouseButtonRight),
				},
				eventStart.Add((time.Minute*13)+(time.Second*time.Duration(i))),
			)

			inputWriteAPI.WritePoint(point)

			if questionNumber == 6 {
				questionNumber = 1
			} else {
				questionNumber++
			}
		}
	}

	// Mouse Up Event
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
		var questionNumber = 1
		for i := 0; i < 100; i++ {
			point := influxdb2.NewPoint(
				common.MeasurementMouseUp,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"question_number": questionNumber,
					"x":               rand.Intn(1920),
					"y":               rand.Intn(1080),
					"button":          int64(common.MouseButtonMiddle),
				},
				eventStart.Add((time.Minute*13)+(time.Second*time.Duration(i))),
			)

			inputWriteAPI.WritePoint(point)

			if questionNumber == 6 {
				questionNumber = 1
			} else {
				questionNumber++
			}
		}
	}

	// Mouse Scrolled Event
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
		var questionNumber = 1
		for i := 0; i < 100; i++ {
			point := influxdb2.NewPoint(
				common.MeasurementMouseScrolled,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"question_number": questionNumber,
					"x":               rand.Intn(1920),
					"y":               rand.Intn(1080),
				},
				eventStart.Add((time.Minute*13)+(time.Second*time.Duration(i))),
			)

			inputWriteAPI.WritePoint(point)

			if questionNumber == 6 {
				questionNumber = 1
			} else {
				questionNumber++
			}
		}
	}

	// Window Resize Event
	for _, sessionID := range []string{globalID.String(), globalID2.String()} {
		for i := 0; i < 5; i++ {
			point := influxdb2.NewPoint(
				common.MeasurementWindowSized,
				map[string]string{
					"session_id": sessionID,
				},
				map[string]interface{}{
					"question_number": rand.Intn(5) + 1,
					"width":           i,
					"height":          i,
				},
				eventStart.Add((time.Minute*18)+(time.Second*time.Duration(i))),
			)

			inputWriteAPI.WritePoint(point)
		}
	}

	// Funfact Event
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

		statisticWriteAPI.WritePoint(point)
	}

	sessionWriteAPI.Flush()
	inputWriteAPI.Flush()
	statisticWriteAPI.Flush()

	return nil
}
