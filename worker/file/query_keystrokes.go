package file

import (
	"context"
	"fmt"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Keystroke struct {
	Measurement  string    `json:"_measurement" csv:"_measurement"`
	SessionID    string    `json:"session_id" csv:"session_id"`
	KeyChar      string    `json:"key_char" csv:"key_char"`
	KeyCode      string    `json:"key_code" csv:"key_code"`
	Shift        bool      `json:"shift" csv:"shift"`
	Alt          bool      `json:"alt" csv:"alt"`
	Control      bool      `json:"control" csv:"control"`
	Meta         bool      `json:"meta" csv:"meta"`
	UnrelatedKey bool      `json:"unrelated_key" csv:"unrelated_key"`
	Timestamp    time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryKeystrokes(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (*[]Keystroke, error) {
	keystrokeMouseRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementKeystroke+`" and r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return &[]Keystroke{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}
	defer keystrokeMouseRows.Close()

	var outputKeystroke []Keystroke

	for keystrokeMouseRows.Next() {
		record := keystrokeMouseRows.Record()

		keyChar, ok := record.ValueByKey("key_char").(string)
		if !ok {
			keyChar = ""
		}

		keyCode, ok := record.ValueByKey("key_code").(string)
		if !ok {
			keyCode = ""
		}

		shift, ok := record.ValueByKey("shift").(bool)
		if !ok {
			shift = false
		}

		alt, ok := record.ValueByKey("alt").(bool)
		if !ok {
			alt = false
		}

		control, ok := record.ValueByKey("control").(bool)
		if !ok {
			control = false
		}

		meta, ok := record.ValueByKey("meta").(bool)
		if !ok {
			meta = false
		}

		unrelatedKey, ok := record.ValueByKey("unrelated_key").(bool)
		if !ok {
			unrelatedKey = false
		}

		outputKeystroke = append(outputKeystroke, Keystroke{
			Measurement:  common.MeasurementKeystroke,
			SessionID:    sessionID.String(),
			KeyChar:      keyChar,
			KeyCode:      keyCode,
			Shift:        shift,
			Alt:          alt,
			Control:      control,
			Meta:         meta,
			UnrelatedKey: unrelatedKey,
			Timestamp:    record.Time(),
		})
	}

	return &outputKeystroke, nil
}
