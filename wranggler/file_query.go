package main

import (
	"context"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func (d *Dependency) QueryKeystrokes(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]Keystroke, error) {
	keystrokeMouseRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "coding_event_keystroke")`,
	)
	if err != nil {
		return []Keystroke{}, err
	}

	//var lastTableIndex int = -1
	outputKeystroke := []Keystroke{}
	tempKeystroke := Keystroke{}
	for keystrokeMouseRows.Next() {
		// TODO: no need to use UnmarshalInfluxRow
		// See the implementation on the logger service
		unmarshaledRow, err := UnmarshalInfluxRow(keystrokeMouseRows.Record().String())
		if err != nil {
			return []Keystroke{}, err
		}

		switch unmarshaledRow["_field"].(string) {
		case "key_char":
			tempKeystroke.KeyChar = unmarshaledRow["_value"].(string)
		case "key_code":
			tempKeystroke.KeyCode = unmarshaledRow["_value"].(string)
		case "shift":
			tempBool := false
			if unmarshaledRow["_value"].(string) == "true" {
				tempBool = true
			}
			tempKeystroke.Shift = tempBool
		case "alt":
			tempBool := false
			if unmarshaledRow["_value"].(string) == "true" {
				tempBool = true
			}
			tempKeystroke.Alt = tempBool
		case "control":
			tempBool := false
			if unmarshaledRow["_value"].(string) == "true" {
				tempBool = true
			}
			tempKeystroke.Control = tempBool
		case "unrelated_key":
			tempBool := false
			if unmarshaledRow["_value"].(string) == "true" {
				tempBool = true
			}
			tempKeystroke.UnrelatedKey = tempBool
		case "meta":
			tempKeystroke.Modifier = unmarshaledRow["_value"].(string)
		}

		// create a new one
		tempKeystroke.QuestionNumber = unmarshaledRow["question_number"].(string)
		tempKeystroke.SessionID = unmarshaledRow["session_id"].(string)
		tempKeystroke.Timestamp = keystrokeMouseRows.Record().Time()

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

	log.Println("Pas here 207")
	outputMouseClick := []MouseClick{}
	tempMouseClick := MouseClick{}

	for mouseClickRows.Next() {
		unmarshaledRow, err := UnmarshalInfluxRow(mouseClickRows.Record().String())
		if err != nil {
			return []MouseClick{}, err
		}

		switch unmarshaledRow["_field"].(string) {
		case "left_click":
			tempBool := false
			if unmarshaledRow["_value"].(string) == "true" {
				tempBool = true
			}
			tempMouseClick.LeftClick = tempBool
		case "right_click":
			tempBool := false
			if unmarshaledRow["_value"].(string) == "true" {
				tempBool = true
			}
			tempMouseClick.RightClick = tempBool
		case "middle_click":
			tempBool := false
			if unmarshaledRow["_value"].(string) == "true" {
				tempBool = true
			}
			tempMouseClick.MiddleClick = tempBool
		}

		// create a new one
		tempMouseClick.QuestionNumber = unmarshaledRow["question_number"].(string)
		tempMouseClick.SessionID = unmarshaledRow["session_id"].(string)
		tempMouseClick.Timestamp = mouseClickRows.Record().Time()

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
	for mouseMoveRows.Next() {
		// TODO: remove this, just use normal stuffs instead of
		// reinventing the wheel. lol.
		unmarshaledRow, err := UnmarshalInfluxRow(mouseMoveRows.Record().String())
		if err != nil {
			return []MouseMovement{}, err
		}

		switch unmarshaledRow["_field"].(string) {
		case "direction":
			tempMouseMove.Direction = unmarshaledRow["_value"].(string)
		case "x_position":
			x, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
			}
			tempMouseMove.XPosition = x
		case "y_position":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
			}
			tempMouseMove.YPosition = y
		case "window_height":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
			}
			tempMouseMove.WindowHeight = y
		case "window_width":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return []MouseMovement{}, err
			}
			tempMouseMove.WindowWidth = y
		}

		tempMouseMove.QuestionNumber = unmarshaledRow["question_number"].(string)
		tempMouseMove.SessionID = unmarshaledRow["session_id"].(string)
		tempMouseMove.Timestamp = mouseMoveRows.Record().Time()

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

	for personalInfoRows.Next() {
		// TODO: mabok
		unmarshaledRow, err := UnmarshalInfluxRow(personalInfoRows.Record().String())
		if err != nil {
			return []PersonalInfo{}, err
		}

		switch unmarshaledRow["_field"].(string) {
		case "student_number":
			tempPersonalInfo.StudentNumber = unmarshaledRow["_value"].(string)
		case "hours_of_practice":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return []PersonalInfo{}, err
			}
			tempPersonalInfo.HoursOfPractice = y
		case "years_of_experience":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return []PersonalInfo{}, err
			}
			tempPersonalInfo.YearsOfExperience = y
		case "familiar_language":
			tempPersonalInfo.FamiliarLanguages = unmarshaledRow["_value"].(string)
		}

		tempPersonalInfo.SessionID = unmarshaledRow["session_id"].(string)
		tempPersonalInfo.Timestamp = personalInfoRows.Record().Time()

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
	for samTestRows.Next() {
		unmarshaledRow, err := UnmarshalInfluxRow(samTestRows.Record().String())
		if err != nil {
			return []SamTest{}, err
		}

		switch unmarshaledRow["_field"].(string) {
		case "aroused_level":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return []SamTest{}, err
			}
			tempSamTest.ArousedLevel = y
		case "pleased_level":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return []SamTest{}, err
			}
			tempSamTest.PleasedLevel = y
		}

		tempSamTest.SessionID = unmarshaledRow["session_id"].(string)
		tempSamTest.Timestamp = samTestRows.Record().Time()

		outputSamTest = append(outputSamTest, tempSamTest)
	}

	return outputSamTest, nil
}
