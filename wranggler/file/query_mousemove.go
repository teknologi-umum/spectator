package file

import (
	"context"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func (d *Dependency) QueryMouseMove(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseMovement, error) {
	mouseMoveRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "coding_event_mousemove")
		`,
	)
	if err != nil {
		return []MouseMovement{}, err
	}

	outputMouseMove := []MouseMovement{}
	tempMouseMove := MouseMovement{}
	var tablePosition int64
	for mouseMoveRows.Next() {
		// TODO: remove this, just use normal stuffs instead of
		// reinventing the wheel. lol.

		rows := mouseMoveRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		switch rows.Field() {
		case "direction":
			tempMouseMove.Direction = rows.Value().(string)
		case "x_position":
			x, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
			}
			tempMouseMove.XPosition = x
		case "y_position":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
			}
			tempMouseMove.YPosition = y
		case "window_height":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
			}
			tempMouseMove.WindowHeight = y
		case "window_width":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
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

	return outputMouseMove, nil
}
