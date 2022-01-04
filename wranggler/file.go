package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	pb "worker/proto"
)

type MouseMovement struct {
	SessionID      string    `json:"session_id" csv:"session_id"`
	Type           string    `json:"type" csv:"type"`
	QuestionNumber string    `json:"question_number" csv:"question_number"`
	Direction      string    `json:"direction" csv:"direction"`
	XPosition      int64     `json:"x_position" csv:"x_position"`
	YPosition      int64     `json:"y_position" csv:"y_position"`
	WindowWidth    int64     `json:"window_width" csv:"window_width"`
	WindowHeight   int64     `json:"window_height" csv:"window_height"`
	Timestamp      time.Time `json:"_timestap" csv:"_timestamp"`
}

type Keystroke struct {
	SessionID      string    `json:"session_id" csv:"session_id"`
	Type           string    `json:"type" csv:"type"`
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

type MouseClick struct {
	SessionID      string    `json:"session_id" csv:"session_id"`
	Type           string    `json:"type" csv:"type"`
	QuestionNumber string    `json:"question_number" csv:"question_number"`
	RightClick     bool      `json:"right_click" csv:"right_click"`
	LeftClick      bool      `json:"left_click" csv:"left_click"`
	MiddleClick    bool      `json:"middle_click" csv:"middle_click"`
	Timestamp      time.Time `json:"timestamp" csv:"timestamp"`
}

type PersonalInfo struct {
	Type              string    `json:"type" csv:"type"`
	SessionID         string    `json:"session_id" csv:"session_id"`
	StudentNumber     string    `json:"student_number" csv:"student_number"`
	HoursOfPractice   int64     `json:"hours_of_practice" csv:"hours_of_experience"`
	YearsOfExperience int64     `json:"years_of_experience" csv:"years_of_experience"`
	FamiliarLanguages string    `json:"familiar_languages" csv:"familliar_languages"`
	Timestamp         time.Time `json:"timestamp" csv:"timestamp"`
}

type SamTest struct {
	SessionID    string    `json:"session_id" csv:"session_id"`
	Type         string    `json:"type" csv:"type"`
	ArousedLevel int64     `json:"aroused_level" csv:"aroused_level"`
	PleasedLevel int64     `json:"pleased_level" csv:"pleased_level"`
	Timestamp    time.Time `json:"timestamp" csv:"timestamp"`
}

// GenerateFile is the handler for generating file into CSV and JSON based on
// the input data (which only contains the Session ID).
func (d *Dependency) GenerateFile(ctx context.Context, in *pb.Member) (*pb.EmptyResponse, error) {
	sessionID, err := uuid.Parse(in.GetSessionId())
	if err != nil {
		return &pb.EmptyResponse{}, fmt.Errorf("parsing uuid: %v", err)
	}

	go d.CreateFile(sessionID)

	return &pb.EmptyResponse{}, nil
}

func (d *Dependency) CreateFile(sessionID uuid.UUID) {
	// Defer a func that will recover from panic.
	// TODO: Send this data into the Logging service.
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r.(error))
		}
	}()

	// Let's create a new context
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	// Now we fetch all the data with the _actor being sessionID.String()
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// keystroke and mouse
	keystrokeMouseRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`
		|> filter(fn : (r) => r["_measurement"] == "coding_event_keystroke")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	//var lastTableIndex int = -1
	output := []Keystroke{}
	temp := Keystroke{}
	for keystrokeMouseRows.Next() {
		unmarshaledRow, err := UnmarshalInfluxRow(keystrokeMouseRows.Record().String())
		if err != nil {
			return
		}

		// tableStr, ok := unmarshaledRow["table"].(string)
		// if !ok {
		// 	continue
		// }

		// table, err := strconv.Atoi(tableStr)
		// if err != nil {
		// 	return
		// }
		// if table == lastTableIndex {
		switch unmarshaledRow["_field"].(string) {
		case "key_char":
			temp.KeyChar = unmarshaledRow["_value"].(string)
		case "key_code":
			temp.KeyCode = unmarshaledRow["_value"].(string)
		case "shift":
			temp.Shift = unmarshaledRow["_value"].(bool)
		case "alt":
			temp.Alt = unmarshaledRow["_value"].(bool)
		case "control":
			temp.Control = unmarshaledRow["_value"].(bool)
		case "unrelated_key":
			temp.UnrelatedKey = unmarshaledRow["_value"].(bool)
		case "meta":
			temp.Modifier = unmarshaledRow["_value"].(string)
		}
		// } else {
		// clear the last temp, but check if its less than zero
		// if lastTableIndex >= 0 {
		// 	output = append(output, temp)
		// }
		// create a new one
		temp.QuestionNumber = unmarshaledRow["question_number"].(string)
		temp.SessionID = unmarshaledRow["session_id"].(string)
		temp.Timestamp = keystrokeMouseRows.Record().Time()
		// lastTableIndex = table
		// }
		output = append(output, temp)
	}

	//keystrokeMouseClickRows
	_, err = queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`
		|> filter(fn : (r) => r["_measurement"] == "coding_event_mouseclick")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	// keystrokeMouseMoveRows
	_, err = queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`
		|> filter(fn : (r) => r["_measurement"] == "coding_event_mousemove")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	// coding test result
	// codeSubmissionRows,
	_, err = queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketSessionEvents+`")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> fliter(fn: (r) => r["_measurement"] == "code_submission")`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	// user
	_, err = queryAPI.Query(ctx, `
	from(bucket: "`+BucketSessionEvents+`")
	|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
	|> filter(fn: (r) => (r["event"] == "sam_test_before") or
		(r["event"] == "personal_info"))
	`)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	// Then, we'll write to 2 different files with 2 different formats.
	// Do this repeatedly for each event.
	//
	// So in the end, we have multiple files,
	// one is about the keystroke & mouse events
	// one is about coding test results
	// one is all about the user (personal info, sam test)
	//
	// After that, store the file into MinIO
	// then, put the MinIO link on the influxdb database
	// in a different bucket. You might want to check and do a
	// create if not exists on the bucket.
	// So you'd make sure you're not inserting data into a
	// nil bucket.
}

func UnmarshalInfluxRow(row string) (map[string]interface{}, error) {
	// because csv.NewReader() accepts io.Reader, we'll create one from strings pkg
	input := strings.NewReader(row)
	reader := csv.NewReader(input)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	records, err := reader.Read()
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("reading row value to csv: %v", err)
	}

	// find records length
	// because it's a jagged array, we'll do a nested one
	var recordsLength = len(records)

	output := make(map[string]interface{}, recordsLength)
	for _, rec := range records {
		kv := strings.Split(rec, ":")
		output[kv[0]] = kv[1]
	}

	return output, nil
}

func ConvertDataToJSON(input []interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		return []byte{}, err
	}

	return data, err
}

func ConvertDataToCSV(input []interface{}) ([]byte, error) {
	w := &bytes.Buffer{}
	writer := csv.NewWriter(w)
	// Because csv package does not have something like
	// json.Marshal, we'll gonna do what Thanos did.
	//
	// "Fine. I'll do it myself."

	// Create the CSV headers first
	structType := reflect.TypeOf(input[0])
	headers := make([]string, structType.NumField())
	for i := 0; i < structType.NumField(); i++ {
		headers = append(headers, structType.Field(i).Tag.Get("csv"))
	}

	err := writer.Write(headers)
	if err != nil {
		return []byte{}, err
	}

	for _, inputItem := range input {
		// Struct are always in-order, so it's easy to
		// put it into the temporary
		structValue := reflect.ValueOf(inputItem)
		data := make([]string, structValue.NumField())

		for k := 0; k < structValue.NumField(); k++ {
			currentValue := structValue.Field(k)

			switch currentValue.Interface().(type) {
			case bool:
				data = append(data, strconv.FormatBool(currentValue.Bool()))
				continue
			case string:
				data = append(data, currentValue.String())
				continue
			case uint:
				data = append(data, strconv.FormatUint(currentValue.Uint(), 10))
			case int64:
				data = append(data, strconv.FormatInt(currentValue.Int(), 10))
				continue
			case int:
				data = append(data, strconv.FormatInt(currentValue.Int(), 10))
				continue
			case time.Time:
				t, ok := currentValue.Interface().(time.Time)
				if !ok {
					return []byte{}, fmt.Errorf("struct name of %s has a type of time.Time yet cannot be parsed", currentValue.Type().Name())
				}
				data = append(data, t.Format(time.RFC3339Nano))
				continue
			default:
				return []byte{}, fmt.Errorf("struct name of %s has a weird and unsupported type", currentValue.Type().Name())
			}
		}

		err := writer.Write(data)
		if err != nil {
			return []byte{}, err
		}
	}

	writer.Flush()
	if writer.Error() != nil {
		return []byte{}, fmt.Errorf("last csv write error: %v", err)
	}

	return w.Bytes(), nil
}
