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

type MouseClick struct {
	SessionID          string             `json:"session_id" csv:"session_id"`
	MouseUpX           int64              `json:"mouse_up_x" csv:"mouse_up_x"`
	MouseDownX         int64              `json:"mouse_down_x" csv:"mouse_down_x"`
	MouseUpY           int64              `json:"mouse_up_y" csv:"mouse_up_y"`
	MouseDownY         int64              `json:"mouse_down_y" csv:"mouse_down_y"`
	MouseUpTimestamp   time.Time          `json:"mouse_up_timestamp" csv:"mouse_up_timestamp"`
	MouseDownTimestamp time.Time          `json:"mouse_down_timestamp" csv:"mouse_down_timestamp"`
	MouseUpButton      common.MouseButton `json:"mouse_up_button" csv:"mouse_up_button"`
	MouseDownButton    common.MouseButton `json:"mouse_down_button" csv:"mouse_down_button"`
	Duration           int64              `json:"duration" csv:"duration"`
}

func (d *Dependency) QueryMouseDown(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (*[]Mouse, error) {
	return d.queryMouse(ctx, queryAPI, sessionID, common.MeasurementMouseDown)
}

func (d *Dependency) QueryMouseUp(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (*[]Mouse, error) {
	return d.queryMouse(ctx, queryAPI, sessionID, common.MeasurementMouseUp)
}

func (d *Dependency) queryMouse(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID, measurement string) (*[]Mouse, error) {
	// Get the value of the time that the user started and ended the session.
	examStartedRow, err := queryAPI.Query(
		ctx,
		`from (bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => (r["_measurement"] == "`+common.MeasurementExamStarted+`" and
			                  r["session_id"] == "`+sessionID.String()+`"))`,
	)
	if err != nil {
		return &[]Mouse{}, fmt.Errorf("failed to query session start time: %w", err)
	}
	defer examStartedRow.Close()

	var startTime int64
	if examStartedRow.Next() {
		startTime = examStartedRow.Record().Time().Unix()
	}

	mouseRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketInputEvents+`")
			|> range(start: `+fmt.Sprint(startTime)+`)
			|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`" and
								 r["_measurement"] == "`+measurement+`")
			|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
			|> drop(columns: ["question_number", "_start", "_stop"])
			|> sort(columns: ["_time"])`,
	)
	if err != nil {
		return &[]Mouse{}, fmt.Errorf("failed to query mouse up: %w", err)
	}
	defer mouseRows.Close()

	var outputMouse []Mouse
	for mouseRows.Next() {
		record := mouseRows.Record()

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

	return &outputMouse, nil
}

func (d *Dependency) QueryMouseClick(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (*[]MouseClick, error) {
	mouseUpRows, err := d.QueryMouseUp(ctx, queryAPI, sessionID)
	if err != nil {
		return &[]MouseClick{}, fmt.Errorf("failed to query mouse up: %w", err)
	}

	mouseDownRows, err := d.QueryMouseDown(ctx, queryAPI, sessionID)
	if err != nil {
		return &[]MouseClick{}, fmt.Errorf("failed to query mouse down: %w", err)
	}

	// zip mouseUp and mouseDown since they're coming in pair
	var outputMouse []MouseClick
	for i, mouseUp := range *mouseUpRows {
		// just in case there's a mouse down without a mouse up
		// this should never happen
		if i >= len(*mouseDownRows) {
			break
		}

		mouseDown := (*mouseDownRows)[i]

		outputMouse = append(outputMouse, MouseClick{
			SessionID:          sessionID.String(),
			MouseUpX:           mouseUp.X,
			MouseDownX:         mouseDown.X,
			MouseUpY:           mouseUp.Y,
			MouseDownY:         mouseDown.Y,
			MouseUpButton:      mouseUp.Button,
			MouseDownButton:    mouseDown.Button,
			MouseUpTimestamp:   mouseUp.Timestamp,
			MouseDownTimestamp: mouseDown.Timestamp,
			Duration:           mouseUp.Timestamp.Sub(mouseDown.Timestamp).Milliseconds(),
		})
	}

	return &outputMouse, nil
}
