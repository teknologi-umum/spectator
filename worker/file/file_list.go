package file

import (
	"context"
	"fmt"
	"strings"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
)

// File contains the struct regarding the file object
// that is stored on the InfluxDB database.
type File struct {
	CSVFile       string
	JSONFile      string
	StudentNumber string
	SessionId     string
}

// ListFiles fetch all the files (most importantly, the URL to the MinIO)
// that was generated for the specific session ID.
func (d *Dependency) ListFiles(ctx context.Context, sessionID uuid.UUID) ([]File, error) {
	testFileRows, err := d.DB.QueryAPI(d.DBOrganization).Query(
		ctx,
		`from(bucket: "`+common.BucketFileEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementExportedData+`")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return []File{}, fmt.Errorf("failed to query exported_data: %w", err)
	}
	defer func() {
		err := testFileRows.Close()
		if err != nil {
			log.Err(err).Msg("closing testFileRows")
		}
	}()

	var outputFile []File
	for testFileRows.Next() {
		var ok bool
		rows := testFileRows.Record()
		var file File

		file.StudentNumber, ok = rows.ValueByKey("student_number").(string)
		if !ok {
			file.StudentNumber = ""
		}

		file.SessionId, ok = rows.ValueByKey("session_id").(string)
		if !ok {
			file.SessionId = ""
		}

		file.JSONFile, ok = rows.ValueByKey("file_json_url").(string)
		if !ok {
			file.JSONFile = ""
		}

		file.CSVFile, ok = rows.ValueByKey("file_csv_url").(string)
		if !ok {
			file.CSVFile = ""
		}

		outputFile = append(outputFile, file)
	}

	newCtx, newCancel := context.WithTimeout(ctx, time.Second*15)
	defer newCancel()

	for _, i := range outputFile {
		_, err := d.Bucket.StatObject(newCtx, "public", strings.Replace(i.JSONFile, "/public", "", 1), minio.GetObjectOptions{})
		if err != nil {
			errCode := minio.ToErrorResponse(err)
			if errCode.Code == "NoSuchKey" {
				return []File{}, fmt.Errorf("no %s file: still processing", i.JSONFile)
			}
		}

		_, err = d.Bucket.StatObject(newCtx, "public", strings.Replace(i.CSVFile, "/public", "", 1), minio.GetObjectOptions{})
		if err != nil {
			errCode := minio.ToErrorResponse(err)
			if errCode.Code == "NoSuchKey" {
				return []File{}, fmt.Errorf("no %s file: still processing", i.CSVFile)
			}
		}
	}

	return outputFile, nil
}
