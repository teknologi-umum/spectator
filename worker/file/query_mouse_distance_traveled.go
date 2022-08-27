package file

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/rs/zerolog/log"
)

type MouseDistanceTraveled struct {
	Measurement    string  `json:"_measurement" csv:"_measurement"`
	SessionID      string  `json:"session_id" csv:"session_id"`
	QuestionNumber int     `json:"question_number" csv:"question_number"`
	Distance       float64 `json:"distance" csv:"distance"`
}

func (d *Dependency) QueryMouseDistanceTraveled(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseDistanceTraveled, error) {
	var outputDistanceTraveled []MouseDistanceTraveled

	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementMouseMoved+`" and r["session_id"] == `+fmt.Sprintf("\"%s\"", sessionID.String())+`)
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time", "question_number"])`,
	)
	if err != nil {
		return []MouseDistanceTraveled{}, fmt.Errorf("failed to query mouse move - direction: %w", err)
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Err(err).Msg("closing mouseDistanceTraveledRows")
		}
	}()

	var currentQuestionNumber int
	var currentDistance float64
	var lastX int64
	var lastY int64

	var questionNumberIndexes = make(map[int]int)

	for rows.Next() {
		record := rows.Record()

		if record.Time().Year() != 2022 {
			log.Warn().
				Str("current time from record.Time() is not 2022, it's ", strconv.Itoa(record.Time().Year())).
				Msg("invalid date on QueryMouseDistanceTraveled")
		}

		questionNumber, ok := record.ValueByKey("question_number").(int64)
		if !ok {
			// If this happened, this would be an invalid case
			// as the question number would starts at 1
			questionNumber = 0
		}

		x, ok := record.ValueByKey("x").(int64)
		if !ok {
			x = 0
		}

		y, ok := record.ValueByKey("y").(int64)
		if !ok {
			y = 0
		}

		if currentQuestionNumber == 0 {
			// This means it's the first iteration.
			// We should continue but keep every current data
			// to the established variables.

			currentQuestionNumber = int(questionNumber)
			lastX = x
			lastY = y
			continue
		}

		dx := x - lastX
		dy := y - lastY

		if dx == 0 && dy == 0 {
			// This means the mouse didn't move.
			continue
		}

		distance := math.Sqrt(math.Pow(float64(dx), 2) + math.Pow(float64(dy), 2))

		currentDistance += distance
		lastX = x
		lastY = y

		if len(outputDistanceTraveled) > 0 && outputDistanceTraveled[len(outputDistanceTraveled)-1].QuestionNumber == int(questionNumber) {
			outputDistanceTraveled[len(outputDistanceTraveled)-1].Distance += distance
		} else {
			// We should handle the case if the question number is not consecutive.
			if len(outputDistanceTraveled) > 0 {
				questionNumberIndex, found := questionNumberIndexes[int(questionNumber)]

				if found {
					outputDistanceTraveled[questionNumberIndex].Distance += distance
					continue
				}
			}

			outputDistanceTraveled = append(outputDistanceTraveled, MouseDistanceTraveled{
				Measurement:    common.MeasurementMouseDistanceTraveled,
				SessionID:      sessionID.String(),
				QuestionNumber: int(questionNumber),
				Distance:       distance,
			})

			questionNumberIndexes[int(questionNumber)] = len(outputDistanceTraveled) - 1
		}
	}

	return outputDistanceTraveled, nil
}
