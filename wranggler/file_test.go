package main_test

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"

	main "worker"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	// BucketInputEvents is the bucket name for storing
	// keystroke events, window events, and mouse events.
	BucketInputEvents = "input_events"
	// BucketSessionEvents is the bucket name for storing
	// the session events, including their personal information.
	BucketSessionEvents = "session_events"
)

func TestConvertDataToJSON(t *testing.T) {

	influxToken, ok := os.LookupEnv("INFLUX_TOKEN")
	if !ok {
		t.Errorf("INFLUX_TOKEN envar missing")
	}

	influxHost, ok := os.LookupEnv("INFLUX_HOST")
	if !ok {
		t.Errorf("INFLUX_HOST envar missing")
	}

	influxOrg, ok := os.LookupEnv("INFLUX_ORG")
	if !ok {
		t.Errorf("INFLUX_ORG envar missing")
	}

	influxConn := influxdb2.NewClient(influxHost, influxToken)
	defer influxConn.Close()

	queryAPI := influxConn.QueryAPI(influxOrg)
	ctx := context.TODO()

	keystrokeMouseRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] != "")
		|> filter(fn : (r) => r["_measurement"] == "coding_event_keystroke")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		t.Error(err)
		return
	}

	//var lastTableIndex int = -1
	outputKeystroke := []main.Keystroke{}
	tempKeystroke := main.Keystroke{}
	for keystrokeMouseRows.Next() {
		unmarshaledRow, err := main.UnmarshalInfluxRow(keystrokeMouseRows.Record().String())
		if err != nil {
			return
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

	t.Log("ISI")
	//t.Log(outputKeystroke)

	mouseClickRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] != "")
		|> filter(fn : (r) => r["_measurement"] == "coding_event_mouseclick")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	outputMouseClick := []main.MouseClick{}
	tempMouseClick := main.MouseClick{}

	for mouseClickRows.Next() {
		unmarshaledRow, err := main.UnmarshalInfluxRow(mouseClickRows.Record().String())
		if err != nil {
			return
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

	mouseMoveRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] != "")
		|> filter(fn : (r) => r["_measurement"] == "coding_event_mousemove")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	outputMouseMove := []main.MouseMovement{}
	tempMouseMove := main.MouseMovement{}
	for mouseMoveRows.Next() {
		unmarshaledRow, err := main.UnmarshalInfluxRow(mouseMoveRows.Record().String())
		if err != nil {
			return
		}

		switch unmarshaledRow["_field"].(string) {
		case "direction":
			tempMouseMove.Direction = unmarshaledRow["_value"].(string)
		case "x_position":
			x, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return
			}
			tempMouseMove.XPosition = x
		case "y_position":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return
			}
			tempMouseMove.YPosition = y
		case "window_height":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return
			}
			tempMouseMove.WindowHeight = y
		case "window_width":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return
			}
			tempMouseMove.WindowWidth = y
		}

		tempMouseMove.QuestionNumber = unmarshaledRow["question_number"].(string)
		tempMouseMove.SessionID = unmarshaledRow["session_id"].(string)
		tempMouseMove.Timestamp = mouseMoveRows.Record().Time()

		outputMouseMove = append(outputMouseMove, tempMouseMove)
	}

	personalInfoRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] != "")
		|> filter(fn : (r) => r["_measurement"] == "personal_info")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	outputPersonalInfo := []main.PersonalInfo{}
	tempPersonalInfo := main.PersonalInfo{}

	for personalInfoRows.Next() {
		unmarshaledRow, err := main.UnmarshalInfluxRow(mouseMoveRows.Record().String())
		if err != nil {
			return
		}

		switch unmarshaledRow["_field"].(string) {
		case "student_number":
			tempPersonalInfo.StudentNumber = unmarshaledRow["_value"].(string)
		case "hours_of_practice":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return
			}
			tempPersonalInfo.HoursOfPractice = y
		case "years_of_experience":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return
			}
			tempPersonalInfo.YearsOfExperience = y
		case "familiar_language":
			tempPersonalInfo.FamiliarLanguages = unmarshaledRow["_value"].(string)
		}

		tempPersonalInfo.SessionID = unmarshaledRow["session_id"].(string)
		tempPersonalInfo.Timestamp = personalInfoRows.Record().Time()

		outputPersonalInfo = append(outputPersonalInfo, tempPersonalInfo)
	}

	samTestRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] != "")
		|> filter(fn : (r) => r["_measurement"] == "sam_test_before")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	outputSamTest := []main.SamTest{}
	tempSamTest := main.SamTest{}
	for samTestRows.Next() {
		unmarshaledRow, err := main.UnmarshalInfluxRow(samTestRows.Record().String())
		if err != nil {
			return
		}

		switch unmarshaledRow["_field"].(string) {
		case "aroused_level":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return
			}
			tempSamTest.ArousedLevel = y
		case "pleased_level":
			y, err := strconv.ParseInt(unmarshaledRow["_value"].(string), 10, 64)
			if err != nil {
				return
			}
			tempSamTest.PleasedLevel = y
		}

		tempSamTest.SessionID = unmarshaledRow["session_id"].(string)
		tempSamTest.Timestamp = samTestRows.Record().Time()

		outputSamTest = append(outputSamTest, tempSamTest)
	}

}

func TestConvertDataToCSV(t *testing.T) {
	// data := []main.SampleInput{
	// 	{
	// 		Time:  time.Now().Add(time.Second * 1),
	// 		Actor: "James",
	// 		X:     20,
	// 		Y:     13,
	// 	},
	// 	{
	// 		Time:  time.Now().Add(time.Second * 2),
	// 		Actor: "James, Riyadi",
	// 		X:     21,
	// 		Y:     13,
	// 	},
	// 	{
	// 		Time:  time.Now().Add(time.Second * 3),
	// 		Actor: "James Riyadi",
	// 		X:     22,
	// 		Y:     14,
	// 	},
	// }

	// res, err := main.ConvertDataToCSV(data)
	// if err != nil {
	// 	t.Errorf("an error was thrown: %v", err)
	// }

	// t.Log(string(res))
}
