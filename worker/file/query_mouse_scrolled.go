package file

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type MouseScrolled struct {
	SessionID      string      `json:"session_id" csv:"session_id"`
	Type           string      `json:"type" csv:"-"`
	X              string      `json:"x" csv:"x"`
	Y              string      `json:"y" csv:"y"`
	Button         MouseButton `json:"button" csv:"button"`
	Timestamp      time.Time   `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryMouseScrolled(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseDown, error) {
	outputMouseClick := []MouseDown{}
	mouseClickRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "mouse_scrolled" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return []MouseDown{}, fmt.Errorf("failed to query mouse down: %w", err)
	}

	tempMouseClick := MouseDown{}
	var tablePosition int64
	for mouseClickRows.Next() {
		rows := mouseClickRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		v, ok := rows.Value().(MouseButton)
		if !ok {
			v = 0
		}
		tempMouseClick.Button = v

		if d.IsDebug() {
			log.Println(rows.String())
			log.Printf("table %d\n", rows.Table())
		}

		if table != 0 && table > tablePosition {
			outputMouseClick = append(outputMouseClick, tempMouseClick)
			tablePosition = table
		} else {
			var ok bool

			tempMouseClick.SessionID, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempMouseClick.SessionID = ""
			}
			tempMouseClick.Timestamp = rows.Time()
		}
	}

	if len(outputMouseClick) > 0 || tempMouseClick.SessionID != "" {
		outputMouseClick = append(outputMouseClick, tempMouseClick)
	}

	return outputMouseClick, nil
}
