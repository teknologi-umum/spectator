package file

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type MouseMovement struct {
	SessionID string    `json:"session_id" csv:"session_id"`
	Type      string    `json:"type" csv:"-"`
	Direction string    `json:"direction" csv:"direction"`
	X         int64     `json:"x" csv:"x"`
	Y         int64     `json:"y" csv:"y"`
	Timestamp time.Time `json:"timestamp" csv:"_timestamp"`
}

func (d *Dependency) QueryMouseMove(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseMovement, error) {
	var outputMouseMove []MouseMovement

	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+string(MeasurementMouseMoved)+`" and r["session_id"] == `+fmt.Sprintf("\"%s\"", sessionID.String())+`)
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"])`,
	)
	if err != nil {
		return []MouseMovement{}, fmt.Errorf("failed to query mouse move - direction: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		record := rows.Record()
		direction, ok := record.ValueByKey("direction").(string)
		if !ok {
			direction = ""
		}

		x, ok := record.ValueByKey("x").(int64)
		if !ok {
			x = 0
		}

		y, ok := record.ValueByKey("y").(int64)
		if !ok {
			y = 0
		}

		outputMouseMove = append(
			outputMouseMove,
			MouseMovement{
				SessionID: sessionID.String(),
				Type:      "mouse_move",
				Direction: direction,
				X:         x,
				Y:         y,
				Timestamp: record.Time(),
			},
		)
	}

	return outputMouseMove, nil
}
