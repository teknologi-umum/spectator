package file

import (
	"context"
	"log"
	"time"

	loggerpb "worker/logger_proto"

	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// This function should be called as a goroutine.
//
// It will not panic, instead the panic will be caught
func (d *Dependency) CreateFile(requestID string, sessionID uuid.UUID) {
	// Defer a func that will recover from panic.
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r.(error))
		}

		d.Logger.Log(
			r.(error).Error(),
			loggerpb.Level_ERROR.Enum(),
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

	outputKeystroke, err := d.QueryKeystrokes(ctx, queryAPI, sessionID)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "proceed keystroke query",
			},
		)
		return
	}

	outputMouseClick, err := d.QueryMouseClick(ctx, queryAPI, sessionID)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "proceed mouse click query",
			},
		)
		return
	}

	outputMouseMove, err := d.QueryMouseMove(ctx, queryAPI, sessionID)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "proceed mouse move query",
			},
		)
		return
	}

	outputPersonalInfo, err := d.QueryPersonalInfo(ctx, queryAPI, sessionID)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "proceed personal info query",
			},
		)
		return
	}

	outputSamTest, err := d.QuerySAMTest(ctx, queryAPI, sessionID)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "proceed sam test query",
			},
		)
		return
	}

	keystrokeJSON, err := ConvertDataToJSON(outputKeystroke)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "keystroke json convertion",
			},
		)
		return
	}
	keystrokeCSV, err := gocsv.MarshalString(outputKeystroke)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "keystroke csv convertion",
			},
		)
		return
	}
	mousmoveCSV, err := gocsv.MarshalString(outputMouseMove)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "mouse move CSV conversion",
			},
		)
		return
	}
	mousmoveJSON, err := ConvertDataToJSON(outputMouseMove)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "mouse move JSON conversion",
			},
		)
		return
	}
	mousclickCSV, err := gocsv.MarshalString(outputMouseClick)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "mouse click CSV conversion",
			},
		)
		return
	}
	mousclickJSON, err := ConvertDataToJSON(outputMouseClick)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "mouse click JSON conversion",
			},
		)
		return
	}
	personalCSV, err := gocsv.MarshalString(outputPersonalInfo)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "personal info CSV conversion",
			},
		)
		return
	}
	personalJSON, err := ConvertDataToJSON(outputPersonalInfo)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "personal info CSV conversion",
			},
		)
		return
	}
	samtestCSV, err := gocsv.MarshalString(outputSamTest)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "SAM Test CSV conversion",
			},
		)
		return
	}
	samtestJSON, err := ConvertDataToJSON(outputSamTest)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "SAM Test CSV conversion",
			},
		)
		return
	}

	// TODO: refactor. if the output personal info is an array containing only
	// a single element, then the function output should not be an array
	// it would be more cost effective that way.
	studentNumber := outputPersonalInfo[0].StudentNumber

	_, err = mkFileAndUpload(ctx, []byte(keystrokeCSV), studentNumber+"_keystroke.csv", d.Bucket)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "fail to upload keystroke csv",
			},
		)
		return
	}
	_, err = mkFileAndUpload(ctx, keystrokeJSON, studentNumber+"_keystroke.json", d.Bucket)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "fail to upload keystroke json",
			},
		)
		return
	}
	_, err = mkFileAndUpload(ctx, []byte(mousclickCSV), studentNumber+"_mouse_click.csv", d.Bucket)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "fail to upload mouse click csv",
			},
		)
		return
	}
	_, err = mkFileAndUpload(ctx, mousclickJSON, studentNumber+"_mouse_click.json", d.Bucket)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "fail to upload mouse click json",
			},
		)
		return
	}
	_, err = mkFileAndUpload(ctx, []byte(mousmoveCSV), studentNumber+"_mouse_move.csv", d.Bucket)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "fail to upload mouse move csv",
			},
		)
		return
	}
	_, err = mkFileAndUpload(ctx, mousmoveJSON, studentNumber+"_mouse_move.json", d.Bucket)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "fail to upload mouse move json",
			},
		)
		return
	}
	_, err = mkFileAndUpload(ctx, []byte(personalCSV), studentNumber+"_personal.csv", d.Bucket)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "fail to upload personal info csv",
			},
		)
		return
	}
	_, err = mkFileAndUpload(ctx, personalJSON, studentNumber+"_personal.json", d.Bucket)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "fail to upload personal info json",
			},
		)
		return
	}
	_, err = mkFileAndUpload(ctx, []byte(samtestCSV), studentNumber+"_sam_test.csv", d.Bucket)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "fail to upload sam test csv",
			},
		)
		return
	}
	_, err = mkFileAndUpload(ctx, samtestJSON, studentNumber+"_sam_test.json", d.Bucket)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "fail to upload sam test json",
			},
		)
		return
	}

	// TODO: should use WriteAPIBlocking because InfluxDB has no locks
	writeAPI := d.DB.WriteAPI(d.DBOrganization, d.BucketSessionEvents)

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
