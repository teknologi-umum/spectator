package main

import (
	"context"
	"fmt"

	"github.com/influxdata/influxdb-client-go/v2/domain"
)

func (d *Dependency) PrepareBucket(ctx context.Context) error {
	bucketsAPI := d.DB.BucketsAPI()
	bucket, err := bucketsAPI.FindBucketByName(ctx, "log")
	if err != nil {
		return fmt.Errorf("finding bucket: %v", err)
	}

	if !bucket.CreatedAt.IsZero() {
		return nil
	}

	_, err = bucketsAPI.CreateBucketWithName(ctx, &domain.Organization{Name: d.Org}, "log")
	if err != nil {
		return fmt.Errorf("creating bucket: %v", err)
	}
	return nil
}
