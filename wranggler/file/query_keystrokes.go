package file

import (
	"context"
	"fmt"
	"log"
	"time"

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
	keystrokeMouseRows, err := queryAPI.Query(
		ctx,
		reinaldysBuildQuery(queries{
			Level:     "coding_event_keystroke",
			SessionID: sessionID.String(),
			Buckets:   d.BucketInputEvents,
		}),
	)
	if err != nil {
		return []Keystroke{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	//var lastTableIndex int = -1
	outputKeystroke := []Keystroke{}
	tempKeystroke := Keystroke{}
	var tablePosition int64
	for keystrokeMouseRows.Next() {
		rows := keystrokeMouseRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		switch rows.Field() {
		case "key_char":
			tempKeystroke.KeyChar, ok = rows.Value().(string)
			if !ok {
				tempKeystroke.KeyChar = ""
			}
		case "key_code":
			tempKeystroke.KeyCode = rows.Value().(string)
		case "shift":
			tempBool := false
			if rows.Value().(string) == "true" {
				tempBool = true
			}
			tempKeystroke.Shift = tempBool
		case "alt":
			tempBool := false
			if rows.Value().(string) == "true" {
				tempBool = true
			}
			tempKeystroke.Alt = tempBool
		case "control":
			tempBool := false
			if rows.Value().(string) == "true" {
				tempBool = true
			}
			tempKeystroke.Control = tempBool
		case "unrelated_key":
			tempBool := false
			if rows.Value().(string) == "true" {
				tempBool = true
			}
			tempKeystroke.UnrelatedKey = tempBool
		case "meta":
			tempKeystroke.Modifier = rows.Value().(string)
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

	return outputKeystroke, nil
}
