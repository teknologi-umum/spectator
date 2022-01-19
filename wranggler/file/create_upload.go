package file

import (
	"context"
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
)

func mkFileAndUpload(ctx context.Context, b []byte, path string, m *minio.Client) (*minio.UploadInfo, error) {
	f, err := os.Create("./" + path)
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("creating a file: %v", err)
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("writing to a file: %v", err)
	}

	err = f.Sync()
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("syncing a file: %v", err)
	}

	fileStat, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return &minio.UploadInfo{}, fmt.Errorf("getting file stat: %v", err)
	}

	f, err = os.Open("./" + path)
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("opening a file: %v", err)
	}
	defer f.Close()

	upInfo, err := m.PutObject(
		ctx,
		"spectator",
		path,
		f,
		fileStat.Size(),
		minio.PutObjectOptions{ContentType: "application/octet-stream"},
	)
	if err != nil {
		fmt.Println(err)
		return &minio.UploadInfo{}, fmt.Errorf("uploading a file: %v", err)
	}
	fmt.Println("Successfully uploaded bytes: ", upInfo)

	err = os.Remove("./" + path)
	if err != nil {
		return &minio.UploadInfo{}, fmt.Errorf("removing a file: %v", err)
	}
	return &upInfo, nil
}
