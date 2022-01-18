package file

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func (d *Dependency) QueryMouseClick(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseClick, error) {
	mouseClickRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "coding_event_mouseclick")
		`,
	)
	if err != nil {
		return []MouseClick{}, err
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

	// ? : this part ask Reynaldi's i had no ideas.
	if len(outputMouseClick) > 0 || tempMouseClick.SessionID != "" {
		outputMouseClick = append(outputMouseClick, tempMouseClick)
	}

	return outputMouseClick, nil
}
