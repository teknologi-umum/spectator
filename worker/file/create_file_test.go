package file_test

import (
	"context"
	"testing"
	"time"

	"github.com/minio/minio-go/v7"
)

func TestCreateFile(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	
	// Check for MinIO bucket existence
	bucketFound, err := deps.Bucket.BucketExists(ctx, "spectator")
	if err != nil {
		t.Errorf("error checking MinIO bucket: %s\n", err)
	}

	if !bucketFound {
		err = deps.Bucket.MakeBucket(ctx, "spectator", minio.MakeBucketOptions{})
		if err != nil {
			t.Errorf("error creating MinIObucket: %s\n", err)
		}
	}

	deps.CreateFile("TESTING", globalID)
}
