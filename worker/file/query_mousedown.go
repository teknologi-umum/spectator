package file

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type MouseButton int

const (
	Left MouseButton = iota
	Right
	Middle
)

type MouseDown struct {
	SessionID      string      `json:"session_id" csv:"session_id"`
	Type           string      `json:"type" csv:"-"`
	X              int         `json:"x" csv:"x"`
	Y              int         `json:"y" csv:"y"`
	Button         MouseButton `json:"button" csv:"button"`
	Timestamp      time.Time   `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryMouseDown(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseDown, error) {
	mouseClickRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "mouse_down" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return []MouseDown{}, fmt.Errorf("failed to query mouse down: %w", err)
	}

	var outputMouseDown []MouseDown

	for mouseClickRows.Next() {
		rows := mouseClickRows.Record()

		button, ok := rows.ValueByKey("button").(MouseButton)
		if !ok {
			button = 0
		}

		x, ok := rows.ValueByKey("x").(int)
		if !ok {
			x = 0
		}

		y, ok := rows.ValueByKey("y").(int)
		if !ok {
			y = 0
		}

		outputMouseDown = append(outputMouseDown, MouseDown{
			SessionID:      sessionID.String(),
			Type:           "mouse_down",
			X:              x,
			Y:              y,
			Button:         button,
			Timestamp:      rows.Time(),
		})
	}

	return outputMouseDown, nil
}
