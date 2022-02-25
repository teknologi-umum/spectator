package file

import (
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

// mkFileAndUpload creates a file and uploads it to MinIO.
// Bear in mind that it doesn't write to the InfluxDB. The path of the file is specified
// on the `path` parameter.
func (d *Dependency) mkFileAndUpload(ctx context.Context, data []byte, path string) (*minio.UploadInfo, error) {
	file := bytes.NewReader(data)

	uploadInformation, err := d.Bucket.PutObject(
		ctx,
		"spectator",
		path,
		file,
		file.Size(),
		minio.PutObjectOptions{ContentType: "application/octet-stream"},
	)
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("uploading a file: %v", err)
	}

	return &uploadInformation, nil
}
