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
	outputMouseClick := []MouseClick{}
	for _, x := range []string{"right_click", "left_click", "middle_click"} {
		mouseClickRows, err := queryAPI.Query(
			ctx,
			influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
				Measurement: "coding_event_mouseclick",
				SessionID:   sessionID.String(),
				Buckets:     d.BucketInputEvents,
				Field:       x,
			}),
		)
		if err != nil {
			return []MouseClick{}, fmt.Errorf("failed to query mouse clicks: %w", err)
		}

		tempMouseClick := MouseClick{}
		var tablePosition int64
		for mouseClickRows.Next() {
			rows := mouseClickRows.Record()
			table, ok := rows.ValueByKey("table").(int64)
			if !ok {
				table = 0
			}

			switch x {
			case "left_click":
				v, ok := rows.Value().(bool)
				if !ok {
					v = false
				}
				tempMouseClick.LeftClick = v
			case "right_click":
				v, ok := rows.Value().(bool)
				if !ok {
					v = false
				}
				tempMouseClick.RightClick = v
			case "middle_click":
				v, ok := rows.Value().(bool)
				if !ok {
					v = false
				}
				tempMouseClick.MiddleClick = v
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
	}

	return outputMouseClick, nil
}
