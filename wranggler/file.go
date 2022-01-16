package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"

	"github.com/gocarina/gocsv"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	pb "worker/proto"
)

type DataAnu interface {
	Anu()
}

type MouseMovement struct {
	SessionID      string    `json:"session_id" csv:"session_id"`
	Type           string    `json:"type" csv:"-"`
	QuestionNumber string    `json:"question_number" csv:"question_number"`
	Direction      string    `json:"direction" csv:"direction"`
	XPosition      int64     `json:"x_position" csv:"x_position"`
	YPosition      int64     `json:"y_position" csv:"y_position"`
	WindowWidth    int64     `json:"window_width" csv:"window_width"`
	WindowHeight   int64     `json:"window_height" csv:"window_height"`
	Timestamp      time.Time `json:"_timestap" csv:"_timestamp"`
}

func (MouseMovement) Anu() {}

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

func (Keystroke) Anu() {}

type MouseClick struct {
	SessionID      string    `json:"session_id" csv:"session_id"`
	Type           string    `json:"type" csv:"-"`
	QuestionNumber string    `json:"question_number" csv:"question_number"`
	RightClick     bool      `json:"right_click" csv:"right_click"`
	LeftClick      bool      `json:"left_click" csv:"left_click"`
	MiddleClick    bool      `json:"middle_click" csv:"middle_click"`
	Timestamp      time.Time `json:"timestamp" csv:"timestamp"`
}

func (MouseClick) Anu() {}

type PersonalInfo struct {
	Type              string    `json:"type" csv:"-"`
	SessionID         string    `json:"session_id" csv:"session_id"`
	StudentNumber     string    `json:"student_number" csv:"student_number"`
	HoursOfPractice   int64     `json:"hours_of_practice" csv:"hours_of_experience"`
	YearsOfExperience int64     `json:"years_of_experience" csv:"years_of_experience"`
	FamiliarLanguages string    `json:"familiar_languages" csv:"familliar_languages"`
	Timestamp         time.Time `json:"timestamp" csv:"timestamp"`
}

func (PersonalInfo) Anu() {}

type SamTest struct {
	SessionID    string    `json:"session_id" csv:"session_id"`
	Type         string    `json:"type" csv:"-"`
	ArousedLevel int64     `json:"aroused_level" csv:"aroused_level"`
	PleasedLevel int64     `json:"pleased_level" csv:"pleased_level"`
	Timestamp    time.Time `json:"timestamp" csv:"timestamp"`
}

func (SamTest) Anu() {}

// GenerateFile is the handler for generating file into CSV and JSON based on
// the input data (which only contains the Session ID).
func (d *Dependency) GenerateFiles(ctx context.Context, in *pb.Member) (*pb.EmptyResponse, error) {
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	// Now we fetch all the data with the _actor being sessionID.String()
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// keystroke and mouse
	keystrokeMouseRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "coding_event_keystroke")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Fatalln(err)
		return
	}

	//var lastTableIndex int = -1
	outputKeystroke := []Keystroke{}
	tempKeystroke := Keystroke{}
	for keystrokeMouseRows.Next() {
		unmarshaledRow, err := UnmarshalInfluxRow(keystrokeMouseRows.Record().String())
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

	//t.Log(outputKeystroke)

	mouseClickRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "coding_event_mouseclick")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	log.Println("Pas here 207")
	outputMouseClick := []MouseClick{}
	tempMouseClick := MouseClick{}

	for mouseClickRows.Next() {
		unmarshaledRow, err := UnmarshalInfluxRow(mouseClickRows.Record().String())
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
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "coding_event_mousemove")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	outputMouseMove := []MouseMovement{}
	tempMouseMove := MouseMovement{}
	for mouseMoveRows.Next() {
		// TODO: remove this, just use normal stuffs instead of
		// reinventing the wheel. lol.
		unmarshaledRow, err := UnmarshalInfluxRow(mouseMoveRows.Record().String())
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
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "personal_info")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	outputPersonalInfo := []PersonalInfo{}
	tempPersonalInfo := PersonalInfo{}

	for personalInfoRows.Next() {
		unmarshaledRow, err := UnmarshalInfluxRow(mouseMoveRows.Record().String())
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
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "sam_test_before")
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	outputSamTest := []SamTest{}
	tempSamTest := SamTest{}
	for samTestRows.Next() {
		unmarshaledRow, err := UnmarshalInfluxRow(samTestRows.Record().String())
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

	keystrokeJSON, err := ConvertDataToJSON(outputKeystroke)
	if err != nil {
		return
	}
	keystrokeCSV, err := gocsv.MarshalString(outputKeystroke)
	if err != nil {
		return
	}
	mousmoveCSV, err := gocsv.MarshalString(outputMouseMove)
	if err != nil {
		return
	}
	mousmoveJSON, err := ConvertDataToJSON(outputMouseMove)
	if err != nil {
		return
	}
	mousclickCSV, err := gocsv.MarshalString(outputMouseClick)
	if err != nil {
		return
	}
	mousclickJSON, err := ConvertDataToJSON(outputMouseClick)
	if err != nil {
		return
	}
	personalCSV, err := gocsv.MarshalString(outputPersonalInfo)
	if err != nil {
		return
	}
	personalJSON, err := ConvertDataToJSON(outputPersonalInfo)
	if err != nil {
		return
	}
	samtestCSV, err := gocsv.MarshalString(outputSamTest)
	if err != nil {
		return
	}
	samtestJSON, err := ConvertDataToJSON(outputSamTest)
	if err != nil {
		return
	}

	studentNumber := tempPersonalInfo.StudentNumber

	mkFileAndUpload(ctx, []byte(keystrokeCSV), studentNumber+"_keystroke.csv", d.Bucket)
	mkFileAndUpload(ctx, keystrokeJSON, studentNumber+"_keystroke.json", d.Bucket)
	mkFileAndUpload(ctx, []byte(mousclickCSV), studentNumber+"_mouse_click.csv", d.Bucket)
	mkFileAndUpload(ctx, mousclickJSON, studentNumber+"_mouse_click.json", d.Bucket)
	mkFileAndUpload(ctx, []byte(mousmoveCSV), studentNumber+"_mouse_move.csv", d.Bucket)
	mkFileAndUpload(ctx, mousmoveJSON, studentNumber+"_mouse_move.json", d.Bucket)
	mkFileAndUpload(ctx, []byte(personalCSV), studentNumber+"_personal.csv", d.Bucket)
	mkFileAndUpload(ctx, personalJSON, studentNumber+"_personal.json", d.Bucket)
	mkFileAndUpload(ctx, []byte(samtestCSV), studentNumber+"_sam_test.csv", d.Bucket)
	mkFileAndUpload(ctx, samtestJSON, studentNumber+"_sam_test.json", d.Bucket)

	// TODO
	writeAPI := d.DB.WriteAPI(d.DBOrganization, BucketSessionEvents)

	lazyDir := []string{
		studentNumber + "_keystroke",
		studentNumber + "_mouse_click",
		studentNumber + "_mouse_move",
		studentNumber + "_personal",
		studentNumber + "_sam_test",
	}

	for _, item := range lazyDir {
		e := influxdb2.NewPointWithMeasurement("test_result")
		e.AddTag("session_id", sessionID.String())
		e.AddTag("student_number", studentNumber)
		e.AddField("file_csv_url", "/public/"+item+".csv")
		e.AddField("file_json_url", "/public/"+item+".json")
		e.SetTime(time.Now())
		writeAPI.WritePoint(e)
	}

	writeAPI.Flush()

	log.Println(tempPersonalInfo.StudentNumber)

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

// UnmarshalInfluxRow converts a row from InfluxDB into a map[string]interface{}
//
// Deprecated: use regular row parsing provided by InfluxDB client library
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

func ConvertDataToJSON(input interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		return []byte{}, err
	}

	return data, err
}

func ConvertDataToCSV(inputp interface{}) ([]byte, error) {
	input, ok := inputp.([]interface{})
	if !ok {
		return []byte{}, errors.New("failed to infer data type to array of interfaces")
	}

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

func mkFileAndUpload(ctx context.Context, b []byte, path string, m *minio.Client) (*minio.UploadInfo, error) {
	f, err := os.Create("./" + path)
	if err != nil {
		return &minio.UploadInfo{}, err
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		return &minio.UploadInfo{}, err
	}

	f.Sync()

	fileStat, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return &minio.UploadInfo{}, err
	}

	f, err = os.Open("./" + path)
	if err != nil {
		return &minio.UploadInfo{}, err
	}
	defer f.Close()

	upInfo, err := m.PutObject(
		ctx,
		"spectator",
		path,
		f,
		fileStat.Size(),
		minio.PutObjectOptions{ContentType: "application/octet-stream"},
	)
	if err != nil {
		fmt.Println(err)
		return &minio.UploadInfo{}, err
	}
	fmt.Println("Successfully uploaded bytes: ", upInfo)

	err = os.Remove("./" + path)
	if err != nil {
		return &minio.UploadInfo{}, err
	}
	return &upInfo, nil
}
