package file

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type MouseMovement struct {
	SessionID    string    `json:"session_id" csv:"session_id"`
	Type         string    `json:"type" csv:"-"`
	Direction    string    `json:"direction" csv:"direction"`
	X            int64     `json:"x" csv:"x"`
	Y            int64     `json:"y" csv:"y"`
	WindowWidth  int64     `json:"window_width" csv:"window_width"`
	WindowHeight int64     `json:"window_height" csv:"window_height"`
	Timestamp    time.Time `json:"timestamp" csv:"_timestamp"`
}

func (d *Dependency) QueryMouseMove(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseMovement, error) {
	mouseMoveRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "mouse_move" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"])`,
	)
	if err != nil {
		return []MouseMovement{}, fmt.Errorf("failed to query mouse_move: %w", err)
	}
	defer mouseMoveRows.Close()

	var outputMouseMove []MouseMovement

	for mouseMoveRows.Next() {
		rows := mouseMoveRows.Record()

		direction, ok := rows.ValueByKey("direction").(string)
		if !ok {
			direction = ""
		}

		x, ok := rows.ValueByKey("x").(int64)
		if !ok {
			x = 0
		}

		y, ok := rows.ValueByKey("y").(int64)
		if !ok {
			y = 0
		}

		windowWidth, ok := rows.ValueByKey("window_width").(int64)
		if !ok {
			windowWidth = 0
		}

		windowHeight, ok := rows.ValueByKey("window_height").(int64)
		if !ok {
			windowHeight = 0
		}

		outputMouseMove = append(
			outputMouseMove,
			MouseMovement{
				SessionID:    sessionID.String(),
				Type:         "mouse_move",
				Direction:    direction,
				X:            x,
				Y:            y,
				WindowWidth:  windowWidth,
				WindowHeight: windowHeight,
				Timestamp:    rows.Time(),
			},
		)
	}

	return outputMouseMove, nil
}
