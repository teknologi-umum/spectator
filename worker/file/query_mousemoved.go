package file

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/rs/zerolog/log"
)

type MouseMovement struct {
	Measurement string    `json:"_measurement" csv:"_measurement"`
	SessionID   string    `json:"session_id" csv:"session_id"`
	Direction   string    `json:"direction" csv:"direction"`
	X           int64     `json:"x" csv:"x"`
	Y           int64     `json:"y" csv:"y"`
	Timestamp   time.Time `json:"timestamp" csv:"_timestamp"`
}

func (d *Dependency) QueryMouseMove(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseMovement, error) {
	var outputMouseMove []MouseMovement

	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementMouseMoved+`" and r["session_id"] == `+fmt.Sprintf("\"%s\"", sessionID.String())+`)
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"])`,
	)
	if err != nil {
		return []MouseMovement{}, fmt.Errorf("failed to query mouse move - direction: %w", err)
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Err(err).Msg("closing mouseMoveRows")
		}
	}()

	for rows.Next() {
		record := rows.Record()

		if record.Time().Year() != 2022 {
			log.Warn().
				Str("current time from record.Time() is not 2022, it's ", strconv.Itoa(record.Time().Year())).
				Msg("invalid date on QueryMouseMove")
		}

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
				SessionID:   sessionID.String(),
				Measurement: common.MeasurementMouseMoved,
				Direction:   direction,
				X:           x,
				Y:           y,
				Timestamp:   record.Time(),
			},
		)
	}

	return outputMouseMove, nil
}
