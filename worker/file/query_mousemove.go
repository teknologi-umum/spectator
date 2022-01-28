package file

import (
	"context"
	"fmt"
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

	var outputMouseMove []MouseMovement

	directionRows, err := queryAPI.Query(
		ctx,
		influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
			Measurement: "mousemove",
			SessionID:   sessionID.String(),
			Buckets:     d.BucketInputEvents,
			Field:       "direction",
			SortByTime:  true,
		}),
	)
	if err != nil {
		return []MouseMovement{}, fmt.Errorf("failed to query mouse move - direction: %w", err)
	}
	defer directionRows.Close()

	for directionRows.Next() {
		rows := directionRows.Record()

		direction, ok := rows.Value().(string)
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

		outputMouseMove = append(
			outputMouseMove,
			MouseMovement{
				Direction:      direction,
				QuestionNumber: questionNumber,
				SessionID:      sessionID,
				Timestamp:      timestamp,
			},
		)
	}

	xPositionRows, err := queryAPI.Query(
		ctx,
		influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
			Measurement: "mousemove",
			SessionID:   sessionID.String(),
			Buckets:     d.BucketInputEvents,
			Field:       "x_position",
			SortByTime:  true,
		}),
	)
	if err != nil {
		return []MouseMovement{}, fmt.Errorf("failed to query mouse move - x_position: %w", err)
	}
	defer xPositionRows.Close()

	for xPositionRows.Next() {
		rows := xPositionRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		xPosition, ok := rows.Value().(int64)
		if !ok {
			// FIXME: add default value
		}

		outputMouseMove[table].XPosition = xPosition
	}

	yPositionRows, err := queryAPI.Query(
		ctx,
		influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
			Measurement: "mousemove",
			SessionID:   sessionID.String(),
			Buckets:     d.BucketInputEvents,
			Field:       "y_position",
			SortByTime:  true,
		}),
	)
	if err != nil {
		return []MouseMovement{}, fmt.Errorf("failed to query mouse move - y_position: %w", err)
	}
	defer yPositionRows.Close()

	for yPositionRows.Next() {
		rows := yPositionRows.Record()

		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		yPosition, ok := rows.Value().(int64)
		if !ok {
			// FIXME: add default value
		}

		outputMouseMove[table].YPosition = yPosition
	}

	windowWidthRows, err := queryAPI.Query(
		ctx,
		influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
			Measurement: "mousemove",
			SessionID:   sessionID.String(),
			Buckets:     d.BucketInputEvents,
			Field:       "window_width",
			SortByTime:  true,
		}),
	)
	if err != nil {
		return []MouseMovement{}, fmt.Errorf("failed to query mouse move - window_width: %w", err)
	}
	defer windowWidthRows.Close()

	for windowWidthRows.Next() {
		rows := windowWidthRows.Record()

		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		windowWidth, ok := rows.Value().(int64)
		if !ok {
			// FIXME: add default value
		}

		outputMouseMove[table].WindowWidth = windowWidth
	}

	windowHeightRows, err := queryAPI.Query(
		ctx,
		influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
			Measurement: "mousemove",
			SessionID:   sessionID.String(),
			Buckets:     d.BucketInputEvents,
			Field:       "window_height",
			SortByTime:  true,
		}),
	)
	if err != nil {
		return []MouseMovement{}, fmt.Errorf("failed to query mouse move - window_height: %w", err)
	}
	defer windowHeightRows.Close()

	for windowHeightRows.Next() {
		rows := windowHeightRows.Record()

		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		windowHeight, ok := rows.Value().(int64)
		if !ok {
			// FIXME: add default value
		}

		outputMouseMove[table].WindowHeight = windowHeight
	}

	return outputMouseMove, nil
}
