package file

import (
	"context"
	"fmt"
	"log"
	"time"
	"worker/influxhelpers"

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

	outputMouseMove := []MouseMovement{}
	for _, x := range []string{"direction", "x_position", "y_position", "window_width", "window_height"} {
		mouseMoveRows, err := queryAPI.Query(
			ctx,
			influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
				Measurement: "coding_event_mousemove",
				SessionID:   sessionID.String(),
				Buckets:     d.BucketInputEvents,
				Field:       x,
			}),
		)
		if err != nil {
			return []MouseMovement{}, fmt.Errorf("failed to query mouse moves: %w", err)
		}

		tempMouseMove := MouseMovement{}
		var tablePosition int64
		for mouseMoveRows.Next() {

			rows := mouseMoveRows.Record()
			table, ok := rows.ValueByKey("table").(int64)
			if !ok {
				table = 0
			}

			switch x {
			case "direction":
				tempMouseMove.Direction, ok = rows.Value().(string)
				if !ok {
					tempMouseMove.Direction = ""
				}
			case "x_position":
				x, ok := rows.Value().(int64)
				if !ok {
					return []MouseMovement{}, fmt.Errorf("failed to parse x position type")
				}
				tempMouseMove.XPosition = x
			case "y_position":
				y, ok := rows.Value().(int64)
				if !ok {
					return []MouseMovement{}, fmt.Errorf("failed to parse y position type")
				}
				tempMouseMove.YPosition = y
			case "window_height":
				y, ok := rows.Value().(int64)
				if !ok {
					return []MouseMovement{}, fmt.Errorf("failed to parse window height type")
				}
				tempMouseMove.WindowHeight = y
			case "window_width":
				y, ok := rows.Value().(int64)
				if !ok {
					return []MouseMovement{}, fmt.Errorf("failed to parse window width type")
				}
				tempMouseMove.WindowWidth = y
			}

			if d.IsDebug() {
				log.Println(rows.String())
				log.Printf("table %d\n", rows.Table())
			}

			if table != 0 && table > tablePosition {
				outputMouseMove = append(outputMouseMove, tempMouseMove)
				tablePosition = table
			} else {
				var ok bool

				tempMouseMove.QuestionNumber, ok = rows.ValueByKey("question_number").(string)
				if !ok {
					tempMouseMove.QuestionNumber = ""
				}

				tempMouseMove.SessionID, ok = rows.ValueByKey("session_id").(string)
				if !ok {
					tempMouseMove.SessionID = ""
				}
				tempMouseMove.Timestamp = rows.Time()
			}
		}

		// ? : this part ask Reynaldi's i had no ideas.
		if len(outputMouseMove) > 0 || tempMouseMove.SessionID != "" {
			outputMouseMove = append(outputMouseMove, tempMouseMove)
		}

	}
	return outputMouseMove, nil
}
