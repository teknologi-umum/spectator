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
		// TODO: handle this error!
		return
	}

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

	_, err = mkFileAndUpload(ctx, []byte(keystrokeCSV), studentNumber+"_keystroke.csv", d.Bucket)
	if err != nil {
		return
	}
	_, err = mkFileAndUpload(ctx, keystrokeJSON, studentNumber+"_keystroke.json", d.Bucket)
	if err != nil {
		return
	}
	_, err = mkFileAndUpload(ctx, []byte(mousclickCSV), studentNumber+"_mouse_click.csv", d.Bucket)
	if err != nil {
		return
	}
	_, err = mkFileAndUpload(ctx, mousclickJSON, studentNumber+"_mouse_click.json", d.Bucket)
	if err != nil {
		return
	}
	_, err = mkFileAndUpload(ctx, []byte(mousmoveCSV), studentNumber+"_mouse_move.csv", d.Bucket)
	if err != nil {
		return
	}
	_, err = mkFileAndUpload(ctx, mousmoveJSON, studentNumber+"_mouse_move.json", d.Bucket)
	if err != nil {
		return
	}
	_, err = mkFileAndUpload(ctx, []byte(personalCSV), studentNumber+"_personal.csv", d.Bucket)
	if err != nil {
		return
	}
	_, err = mkFileAndUpload(ctx, personalJSON, studentNumber+"_personal.json", d.Bucket)
	if err != nil {
		return
	}
	_, err = mkFileAndUpload(ctx, []byte(samtestCSV), studentNumber+"_sam_test.csv", d.Bucket)
	if err != nil {
		return
	}
	_, err = mkFileAndUpload(ctx, samtestJSON, studentNumber+"_sam_test.json", d.Bucket)
	if err != nil {
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
