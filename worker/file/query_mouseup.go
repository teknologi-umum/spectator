package file

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type MouseUp struct {
	SessionID string      `json:"session_id" csv:"session_id"`
	Type      string      `json:"type" csv:"-"`
	X         int         `json:"x" csv:"x"`
	Y         int         `json:"y" csv:"y"`
	Button    MouseButton `json:"button" csv:"button"`
	Timestamp time.Time   `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryMouseUp(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseUp, error) {
	mouseClickRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+string(MeasurementMouseUp)+`" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return []MouseUp{}, fmt.Errorf("failed to query mouse up: %w", err)
	}

	var outputMouseUp []MouseUp

	for mouseClickRows.Next() {
		record := mouseClickRows.Record()

		button, ok := record.ValueByKey("button").(MouseButton)
		if !ok {
			button = 0
		}

		x, ok := record.ValueByKey("x").(int)
		if !ok {
			x = 0
		}

		y, ok := record.ValueByKey("y").(int)
		if !ok {
			y = 0
		}

		outputMouseUp = append(outputMouseUp, MouseUp{
			SessionID: sessionID.String(),
			Type:      "mouse_up",
			X:         x,
			Y:         y,
			Button:    button,
			Timestamp: record.Time(),
		})
	}

	return outputMouseUp, nil
}
