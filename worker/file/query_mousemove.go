package file

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type MouseMovement struct {
	SessionID      string    `json:"session_id" csv:"session_id"`
	Type           string    `json:"type" csv:"-"`
	QuestionNumber string    `json:"question_number" csv:"question_number"`
	Direction      string    `json:"direction" csv:"direction"`
	XPosition      int64     `json:"x_position" csv:"x_position"`
	YPosition      int64     `json:"y_position" csv:"y_position"`
	WindowWidth    int64     `json:"window_width" csv:"window_width"`
	WindowHeight   int64     `json:"window_height" csv:"window_height"`
	Timestamp      time.Time `json:"timestamp" csv:"_timestamp"`
}

func (d *Dependency) QueryMouseMove(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseMovement, error) {

	var outputMouseMove []MouseMovement

	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`"
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "mousemove" and r["session_id"] == `+fmt.Sprintf("\"%s\"", sessionID.String())+`)
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"])`,
	)
	if err != nil {
		return []MouseMovement{}, fmt.Errorf("failed to query mouse move - direction: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		rows := rows.Record()

		direction, ok := rows.ValueByKey("direction").(string)
		if !ok {
			// FIXME: add default value
		}

		x, ok := rows.ValueByKey("x").(int64)
		if !ok {
			// FIXME: add default value
		}

		y, ok := rows.ValueByKey("y").(int64)
		if !ok {
			// FIXME: add default value
		}

		windowWidth, ok := rows.ValueByKey("window_width").(int64)
		if !ok {
			// FIXME: add default value
		}

		windowHeight, ok := rows.ValueByKey("window_height").(int64)
		if !ok {
			// FIXME: add default value
		}

		questionNumber, ok := rows.ValueByKey("question_number").(string)
		if !ok {
			questionNumber = ""
		}

		sessionID, ok := rows.ValueByKey("session_id").(string)
		if !ok {
			sessionID = ""
		}
		timestamp := rows.Time()

		outputMouseMove = append(outputMouseMove, MouseMovement{
			SessionID:      sessionID,
			Type:           "mousemove",
			QuestionNumber: questionNumber,
			Direction:      direction,
			XPosition:      x,
			YPosition:      y,
			WindowWidth:    windowWidth,
			WindowHeight:   windowHeight,
			Timestamp:      timestamp,
		})

	}
	return outputMouseMove, nil
}
