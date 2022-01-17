package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type queries struct {
	Level     string
	SessionID string
	Buckets   string
	TimeFrom  time.Time
	TimeTo    time.Time
}

func (d *Dependency) IsDebug() bool {
	return d.Environment == "DEVELOPMENT"
}

func (d *Dependency) QueryKeystrokes(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]Keystroke, error) {
	keystrokeMouseRows, err := queryAPI.Query(
		ctx,
		reinaldysBuildQuery(queries{
			Level:     "coding_event_keystroke",
			SessionID: sessionID.String(),
			Buckets:   BucketInputEvents,
		}),
	)
	if err != nil {
		return []Keystroke{}, err
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

	// ? : this part ask Reynaldi's i had no ideas.
	if len(outputKeystroke) > 0 || tempKeystroke.SessionID != "" {
		outputKeystroke = append(outputKeystroke, tempKeystroke)
	}
	return outputKeystroke, nil
}

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

func (d *Dependency) QueryMouseMove(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]MouseMovement, error) {
	mouseMoveRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "coding_event_mousemove")
		`,
	)
	if err != nil {
		return []MouseMovement{}, err
	}

	outputMouseMove := []MouseMovement{}
	tempMouseMove := MouseMovement{}
	var tablePosition int64
	for mouseMoveRows.Next() {
		// TODO: remove this, just use normal stuffs instead of
		// reinventing the wheel. lol.

		rows := mouseMoveRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		switch rows.Field() {
		case "direction":
			tempMouseMove.Direction = rows.Value().(string)
		case "x_position":
			x, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
			}
			tempMouseMove.XPosition = x
		case "y_position":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
			}
			tempMouseMove.YPosition = y
		case "window_height":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
			}
			tempMouseMove.WindowHeight = y
		case "window_width":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
			}
			tempMouseMove.WindowWidth = y
		}

		if d.IsDebug() {
			log.Println(rows.String())
			log.Printf("table %d\n", rows.Table())
		}

		if table != 0 && table > tablePosition {
			outputMouseMove = append(outputMouseMove, tempMouseMove)
			tablePosition = table
		} else {
			var ok bool

			tempMouseMove.QuestionNumber, ok = rows.ValueByKey("question_number").(string)
			if !ok {
				tempMouseMove.QuestionNumber = ""
			}

			tempMouseMove.SessionID, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempMouseMove.SessionID = ""
			}
			tempMouseMove.Timestamp = rows.Time()
		}
	}

	// ? : this part ask Reynaldi's i had no ideas.
	if len(outputMouseMove) > 0 || tempMouseMove.SessionID != "" {
		outputMouseMove = append(outputMouseMove, tempMouseMove)
	}

	return outputMouseMove, nil
}

func (d *Dependency) QueryPersonalInfo(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]PersonalInfo, error) {
	personalInfoRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "personal_info")
		`,
	)
	if err != nil {
		return []PersonalInfo{}, err
	}

	outputPersonalInfo := []PersonalInfo{}
	tempPersonalInfo := PersonalInfo{}
	var tablePosition int64
	for personalInfoRows.Next() {
		// TODO: mabok
		rows := personalInfoRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		switch rows.Field() {
		case "student_number":
			tempPersonalInfo.StudentNumber, ok = rows.Value().(string)
			if !ok {
				tempPersonalInfo.StudentNumber = ""
			}
		case "hours_of_practice":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []PersonalInfo{}, err
			}
			tempPersonalInfo.HoursOfPractice = y
		case "years_of_experience":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []PersonalInfo{}, err
			}
			tempPersonalInfo.YearsOfExperience = y
		case "familiar_language":
			tempPersonalInfo.FamiliarLanguages, ok = rows.Value().(string)
			if !ok {
				tempPersonalInfo.FamiliarLanguages = ""
			}
		}

		if d.IsDebug() {
			log.Println(rows.String())
			log.Printf("table %d\n", rows.Table())
		}

		if table != 0 && table > tablePosition {
			outputPersonalInfo = append(outputPersonalInfo, tempPersonalInfo)
			tablePosition = table
		} else {
			var ok bool

			tempPersonalInfo.SessionID, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempPersonalInfo.SessionID = ""
			}
			tempPersonalInfo.Timestamp = rows.Time()
		}
	}

	// ? : this part ask Reynaldi's i had no ideas.
	if len(outputPersonalInfo) > 0 || tempPersonalInfo.SessionID != "" {
		outputPersonalInfo = append(outputPersonalInfo, tempPersonalInfo)
	}
	return outputPersonalInfo, nil
}

func (d *Dependency) QuerySAMTest(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]SamTest, error) {
	samTestRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "sam_test_before")
		`,
	)
	if err != nil {
		return []SamTest{}, err
	}

	outputSamTest := []SamTest{}
	tempSamTest := SamTest{}
	var tablePosition int64
	for samTestRows.Next() {
		rows := samTestRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		switch rows.Field() {
		case "aroused_level":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []SamTest{}, err
			}
			tempSamTest.ArousedLevel = y
		case "pleased_level":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []SamTest{}, err
			}
			tempSamTest.PleasedLevel = y
		}

		if d.IsDebug() {
			log.Println(rows.String())
			log.Printf("table %d\n", rows.Table())
		}

		if table != 0 && table > tablePosition {
			outputSamTest = append(outputSamTest, tempSamTest)
			tablePosition = table
		} else {
			var ok bool

			tempSamTest.SessionID, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempSamTest.SessionID = ""
			}
			tempSamTest.Timestamp = rows.Time()
		}
	}

	// ? : this part ask Reynaldi's i had no ideas.
	if len(outputSamTest) > 0 || tempSamTest.SessionID != "" {
		outputSamTest = append(outputSamTest, tempSamTest)
	}

	return outputSamTest, nil
}

func reinaldysBuildQuery(q queries) string {
	var str strings.Builder
	str.WriteString("from(bucket: \"" + q.Buckets + "\")\n")
	// range query
	str.WriteString("|> range(")
	if !q.TimeFrom.IsZero() {
		str.WriteString("start: " + strconv.FormatInt(q.TimeFrom.Unix(), 10))
	} else {
		str.WriteString("start: 0")
	}

	if !q.TimeTo.IsZero() {
		str.WriteString(", stop: " + strconv.FormatInt(q.TimeTo.Unix(), 10))
	}

	str.WriteString(")\n")

	str.WriteString("|> sort(columns: [\"_time\"])\n")

	if q.SessionID == "" {
		str.WriteString("|> group(columns: [\"session_id\", \"_time\"])\n")
	} else {
		str.WriteString("|> group(columns: [\"_time\"])\n")
	}

	if q.Level != "" {
		str.WriteString(`|> filter(fn: (r) => r["_measurement"] == "` + q.Level + `")` + "\n")
	}

	if q.SessionID != "" {
		str.WriteString(`|> filter(fn: (r) => r["session_id"] == "` + q.SessionID + `")` + "\n")
	}

	str.WriteString("|> yield()\n")

	return str.String()
}
