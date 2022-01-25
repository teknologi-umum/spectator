package file

import (
	"context"
	"fmt"
	"log"
	"time"
	"worker/influxhelpers"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Keystroke struct {
	SessionID      string    `json:"session_id" csv:"session_id"`
	Type           string    `json:"type" csv:"-"`
	QuestionNumber string    `json:"question_number" csv:"question_number"`
	KeyChar        string    `json:"key_char" csv:"key_char"`
	KeyCode        string    `json:"key_code" csv:"key_code"`
	Shift          bool      `json:"shift" csv:"shift"`
	Alt            bool      `json:"alt" csv:"alt"`
	Control        bool      `json:"control" csv:"control"`
	UnrelatedKey   bool      `json:"unrelated_key" csv:"control"`
	Modifier       string    `json:"meta" csv:"meta"`
	Timestamp      time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryKeystrokes(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]Keystroke, error) {
	outputKeystroke := []Keystroke{}
	for _, x := range []string{"key_char", "key_code", "shift", "alt", "control", "unrelated_key", "modifier"} {
		keystrokeMouseRows, err := queryAPI.Query(
			ctx,
			influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
				Measurement: "coding_event_keystroke",
				SessionID:   sessionID.String(),
				Buckets:     d.BucketInputEvents,
				Field:       x,
			}),
		)
		if err != nil {
			return []Keystroke{}, fmt.Errorf("failed to query keystrokes: %w", err)
		}

		//var lastTableIndex int = -1
		tempKeystroke := Keystroke{}
		var tablePosition int64
		for keystrokeMouseRows.Next() {
			rows := keystrokeMouseRows.Record()
			table, ok := rows.ValueByKey("table").(int64)
			if !ok {
				table = 0
			}

			switch x {
			case "key_char":
				tempKeystroke.KeyChar, ok = rows.Value().(string)
				if !ok {
					tempKeystroke.KeyChar = ""
				}
			case "key_code":
				tempKeystroke.KeyCode, ok = rows.Value().(string)
				if !ok {
					tempKeystroke.KeyCode = ""
				}
			case "shift":
				v, ok := rows.Value().(bool)
				if !ok {
					v = false
				}
				tempKeystroke.Shift = v
			case "alt":
				v, ok := rows.Value().(bool)
				if !ok {
					v = false
				}
				tempKeystroke.Alt = v
			case "control":
				v, ok := rows.Value().(bool)
				if !ok {
					v = false
				}
				tempKeystroke.Control = v
			case "unrelated_key":
				v, ok := rows.Value().(bool)
				if !ok {
					v = false
				}
				tempKeystroke.UnrelatedKey = v
			case "meta":
				tempKeystroke.Modifier, ok = rows.Value().(string)
				if !ok {
					tempKeystroke.Modifier = ""
				}
			}

			if d.IsDebug() {
				log.Println(rows.String())
				log.Printf("table %d\n", rows.Table())
			}

			if table != 0 && table > tablePosition {
				outputKeystroke = append(outputKeystroke, tempKeystroke)
				tablePosition = table
			} else {
				var ok bool

				tempKeystroke.QuestionNumber, ok = rows.ValueByKey("question_number").(string)
				if !ok {
					tempKeystroke.QuestionNumber = ""
				}

				tempKeystroke.SessionID, ok = rows.ValueByKey("session_id").(string)
				if !ok {
					tempKeystroke.SessionID = ""
				}
				tempKeystroke.Timestamp = rows.Time()
			}
		}

		if len(outputKeystroke) > 0 || tempKeystroke.SessionID != "" {
			outputKeystroke = append(outputKeystroke, tempKeystroke)
		}

	}

	return outputKeystroke, nil
}
