package file

import (
	"context"
	"fmt"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Mouse struct {
	Measurement string             `json:"_measurement" csv:"_measurement"`
	SessionID   string             `json:"session_id" csv:"session_id"`
	X           int64              `json:"x" csv:"x"`
	Y           int64              `json:"y" csv:"y"`
	Button      common.MouseButton `json:"button" csv:"button"`
	Timestamp   time.Time          `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryMouseDown(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]Mouse, error) {
	return d.queryMouse(ctx, queryAPI, sessionID, common.MeasurementMouseDown)
}

func (d *Dependency) QueryMouseUp(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]Mouse, error) {
	return d.queryMouse(ctx, queryAPI, sessionID, common.MeasurementMouseUp)
}

func (d *Dependency) queryMouse(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID, measurement string) ([]Mouse, error) {
	mouseClickRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+measurement+`" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return []Mouse{}, fmt.Errorf("failed to query mouse down: %w", err)
	}

	var outputMouse []Mouse

	for mouseClickRows.Next() {
		record := mouseClickRows.Record()

		button, ok := record.ValueByKey("button").(int64)
		if !ok {
			button = 0
		}

		x, ok := record.ValueByKey("x").(int64)
		if !ok {
			x = 0
		}

		y, ok := record.ValueByKey("y").(int64)
		if !ok {
			y = 0
		}

		outputMouse = append(outputMouse, Mouse{
			Measurement: measurement,
			SessionID:   sessionID.String(),
			X:           x,
			Y:           y,
			Button:      common.MouseButton(button),
			Timestamp:   record.Time(),
		})
	}

	return outputMouse, nil
}
