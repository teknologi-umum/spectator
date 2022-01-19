package file

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type MouseClick struct {
	SessionID      string    `json:"session_id" csv:"session_id"`
	Type           string    `json:"type" csv:"-"`
	QuestionNumber string    `json:"question_number" csv:"question_number"`
	RightClick     bool      `json:"right_click" csv:"right_click"`
	LeftClick      bool      `json:"left_click" csv:"left_click"`
	MiddleClick    bool      `json:"middle_click" csv:"middle_click"`
	Timestamp      time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryMouseClick(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseClick, error) {
	mouseClickRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "coding_event_mouseclick")
		`,
	)
	if err != nil {
		return []MouseClick{}, fmt.Errorf("failed to query mouse clicks: %w", err)
	}

	outputMouseClick := []MouseClick{}
	tempMouseClick := MouseClick{}
	var tablePosition int64
	for mouseClickRows.Next() {
		rows := mouseClickRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		switch rows.Field() {
		case "left_click":
			tempBool := false
			if rows.Value().(string) == "true" {
				tempBool = true
			}
			tempMouseClick.LeftClick = tempBool
		case "right_click":
			tempBool := false
			if rows.Value().(string) == "true" {
				tempBool = true
			}
			tempMouseClick.RightClick = tempBool
		case "middle_click":
			tempBool := false
			if rows.Value().(string) == "true" {
				tempBool = true
			}
			tempMouseClick.MiddleClick = tempBool
		}

		if d.IsDebug() {
			log.Println(rows.String())
			log.Printf("table %d\n", rows.Table())
		}

		if table != 0 && table > tablePosition {
			outputMouseClick = append(outputMouseClick, tempMouseClick)
			tablePosition = table
		} else {
			var ok bool

			tempMouseClick.QuestionNumber, ok = rows.ValueByKey("question_number").(string)
			if !ok {
				tempMouseClick.QuestionNumber = ""
			}

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
