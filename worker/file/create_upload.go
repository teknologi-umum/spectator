package file

import (
	"context"
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
)

// TODO: add documentation of that this function does
func (d *Dependency) mkFileAndUpload(ctx context.Context, data []byte, path string) (*minio.UploadInfo, error) {
	createFile, err := os.Create("./" + path)
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("creating a file: %v", err)
	}
	defer createFile.Close()

	_, err = createFile.Write(data)
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("writing to a file: %v", err)
	}

	err = createFile.Sync()
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("syncing a file: %v", err)
	}

	fileStat, err := createFile.Stat()
	if err != nil {
		fmt.Println(err)
		return &minio.UploadInfo{}, fmt.Errorf("getting file stat: %v", err)
	}

	readFile, err := os.Open("./" + path)
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("opening a file: %v", err)
	}
	defer readFile.Close()

	uploadInformation, err := d.Bucket.PutObject(
		ctx,
		"spectator",
		path,
		readFile,
		fileStat.Size(),
		minio.PutObjectOptions{ContentType: "application/octet-stream"},
	)
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("uploading a file: %v", err)
	}

	err = os.Remove("./" + path)
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("removing a file: %v", err)
	}

	return &uploadInformation, nil
}
