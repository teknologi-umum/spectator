package file

import (
	"context"
	"fmt"
	"log"
	"time"
	"worker/influxhelpers"

	pb "worker/worker_proto"

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
// FIXME: this should not use the direct []*pb.File, should be wrapped
// to another type of struct that later can be processed by the list.go file
// on the main package.
func (d *Dependency) ListFiles(ctx context.Context, sessionID uuid.UUID) ([]*pb.File, error) {
	testFileRows, err := d.DB.QueryAPI(d.DBOrganization).Query(
		ctx,
		influxhelpers.ReinaldysBuildQuery(influxhelpers.Queries{
			Measurement: "test_file",
			SessionID:   sessionID.String(),
			Buckets:     d.BucketSessionEvents,
		}),
	)
	if err != nil {
		return []*pb.File{}, fmt.Errorf("failed to query keystrokes: %w", err)
	}

	outputFile := []*pb.File{}
	tempFile := pb.File{}
	var tablePosition int64
	for testFileRows.Next() {
		rows := testFileRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		switch rows.Field() {
		case "file_url_json":
			tempFile.FileUrlJson, ok = rows.Value().(string)
			if !ok {
				tempFile.FileUrlJson = ""
			}
		case "file_url_csv":
			tempFile.FileUrlCsv, ok = rows.Value().(string)
			if !ok {
				tempFile.FileUrlCsv = ""
			}
		}

		if d.IsDebug() {
			log.Println(rows.String())
			log.Printf("table %d\n", rows.Table())
		}

		if table != 0 && table > tablePosition {
			outputFile = append(outputFile, &tempFile)
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
		outputFile = append(outputFile, &tempFile)
	}

	newCtx, newCancel := context.WithTimeout(ctx, time.Second*15)
	defer newCancel()

	for _, i := range outputFile {
		_, err := d.Bucket.StatObject(newCtx, "spectator", i.FileUrlJson, minio.GetObjectOptions{})
		if err != nil {
			errCode := minio.ToErrorResponse(err)
			if errCode.Code == "NoSuchKey" {
				return []*pb.File{}, fmt.Errorf("no %s file: still processing", i.FileUrlJson)
			}
		}

		_, err = d.Bucket.StatObject(newCtx, "spectator", i.FileUrlCsv, minio.GetObjectOptions{})
		if err != nil {
			errCode := minio.ToErrorResponse(err)
			if errCode.Code == "NoSuchKey" {
				return []*pb.File{}, fmt.Errorf("no %s file: still processing", i.FileUrlCsv)
			}
		}

	}

	return outputFile, nil
}
