package file

import (
	"context"
	"fmt"
	"log"
	"time"

	"worker/influxhelpers"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type File struct {
	CSVFile       string
	JSONFile      string
	StudentNumber string
	SessionId     string
}

// TODO: add documentation on what this function does
func (d *Dependency) ListFiles(ctx context.Context, sessionID uuid.UUID) ([]File, error) {
	testFileRows, err := d.DB.QueryAPI(d.DBOrganization).Query(
		ctx,
		// TODO: remove this query builder
		influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
			Measurement: "test_file",
			SessionID:   sessionID.String(),
			Buckets:     d.BucketSessionEvents,
		}),
	)
	if err != nil {
		return []File{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	var outputFile []File
	var tempFile File
	var tablePosition int64
	for testFileRows.Next() {
		rows := testFileRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		switch rows.Field() {
		case "file_url_json":
			tempFile.JSONFile, ok = rows.Value().(string)
			if !ok {
				tempFile.JSONFile = ""
			}
		case "file_url_csv":
			tempFile.CSVFile, ok = rows.Value().(string)
			if !ok {
				tempFile.CSVFile = ""
			}
		}

		if d.IsDebug() {
			log.Println(rows.String())
			log.Printf("table %d\n", rows.Table())
		}

		if table != 0 && table > tablePosition {
			outputFile = append(outputFile, tempFile)
			tablePosition = table
		} else {
			var ok bool

			tempFile.StudentNumber, ok = rows.ValueByKey("student_number").(string)
			if !ok {
				tempFile.StudentNumber = ""
			}

			tempFile.SessionId, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempFile.SessionId = ""
			}
		}
	}

	if len(outputFile) > 0 || tempFile.SessionId != "" {
		outputFile = append(outputFile, tempFile)
	}

	newCtx, newCancel := context.WithTimeout(ctx, time.Second*15)
	defer newCancel()

	for _, i := range outputFile {
		_, err := d.Bucket.StatObject(newCtx, "spectator", i.JSONFile, minio.GetObjectOptions{})
		if err != nil {
			errCode := minio.ToErrorResponse(err)
			if errCode.Code == "NoSuchKey" {
				return []File{}, fmt.Errorf("no %s file: still processing", i.JSONFile)
			}
		}

		_, err = d.Bucket.StatObject(newCtx, "spectator", i.CSVFile, minio.GetObjectOptions{})
		if err != nil {
			errCode := minio.ToErrorResponse(err)
			if errCode.Code == "NoSuchKey" {
				return []File{}, fmt.Errorf("no %s file: still processing", i.CSVFile)
			}
		}
	}

	return outputFile, nil

}
