package file

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"worker/common"

	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

// convertAndUpload converts the data into both JSON and CSV format,
// then upload it into the MinIO bucket. It also writes the link to the
// InfluxDB database.
//
// The data parameter MUST BE a pointer to a struct, not the direct
// copy of the struct. Otherwise, it will panic.
func (d *Dependency) convertAndUpload(ctx context.Context, writeAPI api.WriteAPIBlocking, data interface{}, fileName string, studentNumber string, requestID string, sessionID uuid.UUID) error {
	// points will store the points that we'll write to the database
	// by batching it.
	var points []*write.Point

	dataJSON, err := json.MarshalIndent(&data, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal json %s data: %w", fileName, err)
	}

	_, err = d.mkFileAndUpload(ctx, dataJSON, studentNumber+"_"+fileName+".json")
	if err != nil {
		return fmt.Errorf("failed to upload json %s file: %w", fileName, err)
	}

	points = append(points, influxdb2.NewPoint(
		common.MeasurementExportedData,
		map[string]string{
			"session_id":     sessionID.String(),
			"student_number": studentNumber,
		},
		map[string]interface{}{
			"file_json_url": "/public/" + studentNumber + "_" + fileName + ".json",
		},
		time.Now(),
	))

	// Now, for the CSV, it's a bit tricky.
	// I'm going to do it with the dumb approach because I don't want
	// to risk the performance by using the reflect stdlib.
	// It's shorter, I know, but the performance is ugh.
	switch s := data.(type) {
	case *KeystrokeEvents:
		dataCSV, err := gocsv.MarshalBytes(*(s.Keystroke))
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %w", fileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+fileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %w", fileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + fileName + ".csv",
			},
			time.Now(),
		))
	case *MouseEvents:
		// MouseDown, MouseUp, MouseMoved, MouseScrolled, MouseDistanceTraveled
		currentFileName := fileName + "_" + "mouse_click"
		dataCSV, err := gocsv.MarshalBytes(*(s.MouseClick))
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %w", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %w", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "mouse_moved"
		dataCSV, err = gocsv.MarshalBytes(*(s.MouseMoved))
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %w", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %w", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "mousescrolled"
		dataCSV, err = gocsv.MarshalBytes(*(s.MouseScrolled))
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %w", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %w", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "mouse_distance_traveled"
		dataCSV, err = gocsv.MarshalBytes(*(s.MouseDistanceTraveled))
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %w", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %w", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))
	case *SolutionEvents:
		// SolutionAcepted, SolutionRejected
		currentFileName := fileName + "_" + "solutionaccepted"
		dataCSV, err := gocsv.MarshalBytes(*(s.SolutionAccepted))
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %w", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %w", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "solutionrejected"
		dataCSV, err = gocsv.MarshalBytes(*(s.SolutionRejected))
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %w", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %w", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))
	case *UserEvents:
		// PersonalInfo, SAMBefore, SAMAfter, ExamStarted, ExamEnded, ExamForfeited, ExamIDEReloaded, DeadlinePassed, Funfact
		currentFileName := fileName + "_" + "personalinfo"
		dataCSV, err := gocsv.MarshalBytes([]PersonalInfo{*(s.PersonalInfo)})
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %w", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %w", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "sambefore"
		dataCSV, err = gocsv.MarshalBytes([]SelfAssessmentManekin{*(s.SelfAssessmentManekinBeforeTest)})
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %w", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %w", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "samafter"
		dataCSV, err = gocsv.MarshalBytes([]SelfAssessmentManekin{*(s.SelfAssessmentManekinAfterTest)})
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %w", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %w", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "examstarted"
		dataCSV, err = gocsv.MarshalBytes([]ExamStarted{*(s.ExamStarted)})
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %w", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %w", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "examended"
		dataCSV, err = gocsv.MarshalBytes([]ExamEvent{*(s.ExamEnded)})
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %v", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %v", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "examforfeited"
		dataCSV, err = gocsv.MarshalBytes([]ExamEvent{*(s.ExamForfeited)})
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %v", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %v", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "examidereloaded"
		dataCSV, err = gocsv.MarshalBytes(*(s.ExamIDEReloaded))
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %v", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %v", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "deadlinepassed"
		dataCSV, err = gocsv.MarshalBytes([]DeadlinePassed{*(s.DeadlinePassed)})
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %v", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %v", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))

		currentFileName = fileName + "_" + "funfact"
		dataCSV, err = gocsv.MarshalBytes([]Funfact{*(s.Funfact)})
		if err != nil {
			return fmt.Errorf("failed to marshal csv %s data: %v", currentFileName, err)
		}

		_, err = d.mkFileAndUpload(ctx, dataCSV, studentNumber+"_"+currentFileName+".csv")
		if err != nil {
			return fmt.Errorf("failed to upload csv %s file: %v", currentFileName, err)
		}

		points = append(points, influxdb2.NewPoint(
			common.MeasurementExportedData,
			map[string]string{
				"session_id":     sessionID.String(),
				"student_number": studentNumber,
			},
			map[string]interface{}{
				"file_csv_url": "/public/" + studentNumber + "_" + currentFileName + ".csv",
			},
			time.Now(),
		))
	default:
		return fmt.Errorf("unknown data type %T", s)
	}

	err = d.DB.WriteAPIBlocking(d.DBOrganization, common.BucketFileEvents).WritePoint(ctx, points...)
	if err != nil {
		return fmt.Errorf("failed to write %s test result: %w", fileName, err)
	}

	return nil
}
