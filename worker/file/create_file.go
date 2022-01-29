package file

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	loggerpb "worker/logger_proto"

	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"golang.org/x/sync/errgroup"
)

// CreateFile creates a file and concurently uploads it to the MinIO bucket.
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

	outputSamBeforeTest, err := d.QueryBeforeExamSam(ctx, queryAPI, sessionID)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "proceed before exam sam test query",
			},
		)
		return
	}

	outputSamAfterTest, err := d.QueryAfterExamSam(ctx, queryAPI, sessionID)
	if err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "proceed after exam sam test query",
			},
		)
		return
	}

	writeAPI := d.DB.WriteAPIBlocking(d.DBOrganization, d.BucketSessionEvents)

	studentNumber := outputPersonalInfo.StudentNumber

	g, gctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return d.convertAndUpload(gctx, writeAPI, outputKeystroke, "keystroke", studentNumber, requestID, sessionID)
	})
	g.Go(func() error {
		return d.convertAndUpload(gctx, writeAPI, outputMouseClick, "mouse_click", studentNumber, requestID, sessionID)
	})
	g.Go(func() error {
		return d.convertAndUpload(gctx, writeAPI, outputMouseMove, "mouse_move", studentNumber, requestID, sessionID)
	})
	g.Go(func() error {
		return d.convertAndUpload(gctx, writeAPI, outputPersonalInfo, "personal_info", studentNumber, requestID, sessionID)
	})
	g.Go(func() error {
		return d.convertAndUpload(gctx, writeAPI, outputSamBeforeTest, "before_exam_sam_test", studentNumber, requestID, sessionID)
	})
	g.Go(func() error {
		return d.convertAndUpload(gctx, writeAPI, outputSamAfterTest, "after_exam_sam_test", studentNumber, requestID, sessionID)
	})

	if err := g.Wait(); err != nil {
		d.Logger.Log(
			err.Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "proceed convert and upload",
			},
		)
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

// convertAndUpload converts the data into both JSON and CSV format,
// then upload it into the MinIO bucket. It also writes the link to the
// InfluxDB database.
func (d *Dependency) convertAndUpload(ctx context.Context, writeAPI api.WriteAPIBlocking, data interface{}, fileName string, studentNumber string, requestID string, sessionID uuid.UUID) error {
	dataJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal json %s data: %v", fileName, err)
	}

	dataCSV, err := gocsv.MarshalBytes(data)
	if err != nil {
		return fmt.Errorf("failed to marshal csv %s data: %v", fileName, err)
	}

	_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+fileName+".csv")
	if err != nil {
		return fmt.Errorf("failed to upload csv %s file: %v", fileName, err)
	}

	_, err = d.mkFileAndUpload(ctx, dataJSON, studentNumber+"_"+fileName+".json")
	if err != nil {
		return fmt.Errorf("failed to upload json %s file: %v", fileName, err)
	}

	point := influxdb2.NewPoint(
		"exported_data",
		map[string]string{
			"session_id":     sessionID.String(),
			"student_number": studentNumber,
		},
		map[string]interface{}{
			"file_csv_url":  "/public/" + studentNumber + "_" + fileName + ".csv",
			"file_json_url": "/public/" + studentNumber + "_" + fileName + ".json",
		},
		time.Now(),
	)

	err = d.DB.WriteAPIBlocking(d.DBOrganization, d.BucketFileEvents).WritePoint(ctx, point)
	if err != nil {
		return fmt.Errorf("failed to write %s test result: %v", fileName, err)
	}

	return nil
}
