package file

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"

	"github.com/gocarina/gocsv"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"worker/logger"
)

// Dependency contains the dependency injection
// to be used on this package.
type Dependency struct {
	Environment         string
	DB                  influxdb2.Client
	Bucket              *minio.Client
	DBOrganization      string
	Logger              logger.LoggerClient
	LoggerToken         string
	BucketInputEvents   string
	BucketSessionEvents string
}

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
	Timestamp      time.Time `json:"timestamp" csv:"_timestamp"`
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

func (d *Dependency) CreateFile(requestID string, sessionID uuid.UUID) {
	// Defer a func that will recover from panic.
	// TODO: Send this data into the Logging service.

	defer func() {
		r := recover()
		if r != nil {
			log.Println(r.(error))
		}

		d.Log(
			r.(error).Error(),
			logger.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "recovering from panic",
			},
		)
	}()

	// Let's create a new context
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	// Now we fetch all the data with the _actor being sessionID.String()
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// TODO: might be better to refactor each of these into their own functions.
	// keystroke and mouse

	outputKeystroke, err := d.QueryKeystrokes(ctx, queryAPI, sessionID)
	if err != nil {
		// TODO: handle this error!
		return
	}
	//t.Log(outputKeystroke)

	outputMouseClick, err := d.QueryMouseClick(ctx, queryAPI, sessionID)
	if err != nil {
		// TODO: handle error
		return
	}

	outputMouseMove, err := d.QueryMouseMove(ctx, queryAPI, sessionID)
	if err != nil {
		// todoo: handle error
		return
	}

	outputPersonalInfo, err := d.QueryPersonalInfo(ctx, queryAPI, sessionID)
	if err != nil {
		// TODO: handle error
		return
	}

	outputSamTest, err := d.QuerySAMTest(ctx, queryAPI, sessionID)
	if err != nil {
		// TODO handle error
		return
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

	// FIXME: this should be like this
	studentNumber := outputPersonalInfo[0].StudentNumber

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

	// TODO: should use WriteAPIBlocking because InfluxDB has no locks
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

	log.Println(studentNumber)

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
