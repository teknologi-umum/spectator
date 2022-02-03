package file

import (
	"context"
	"fmt"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type MouseScrolled struct {
	SessionID string    `json:"session_id" csv:"session_id"`
	Type      string    `json:"type" csv:"-"`
	X         int64     `json:"x" csv:"x"`
	Y         int64     `json:"y" csv:"y"`
	Timestamp time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryMouseScrolled(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseScrolled, error) {
	mouseClickRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementMouseScrolled+`" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return []MouseScrolled{}, fmt.Errorf("failed to query mouse down: %w", err)
	}

	var outputMouseScrolled []MouseScrolled

	for mouseClickRows.Next() {
		record := mouseClickRows.Record()

		x, ok := record.ValueByKey("x").(int64)
		if !ok {
			x = 0
		}

		y, ok := record.ValueByKey("y").(int64)
		if !ok {
			y = 0
		}

		outputMouseScrolled = append(outputMouseScrolled, MouseScrolled{
			SessionID: sessionID.String(),
			Type:      "mouse_scrolled",
			X:         x,
			Y:         y,
			Timestamp: record.Time(),
		})
	}

	return outputMouseScrolled, nil
}
