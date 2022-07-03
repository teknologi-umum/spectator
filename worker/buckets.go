package main

import (
	"context"
	"fmt"
	"worker/common"
)

func (d *Dependency) prepareBuckets(ctx context.Context) error {
	bucketsAPI := d.DB.BucketsAPI()
	orgsAPI := d.DB.OrganizationsAPI()

	_, err := bucketsAPI.FindBucketByName(ctx, common.BucketInputStatisticEvents)
	if err != nil && err.Error() == "bucket '"+common.BucketInputStatisticEvents+"' not found" {
		orgDomain, err := orgsAPI.FindOrganizationByName(ctx, d.DBOrganization)
		if err != nil {
			return fmt.Errorf("failed to find organization: %w", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, common.BucketInputStatisticEvents)
		if err != nil {
			return fmt.Errorf("failed to create bucket: %w", err)
		}
	}

	_, err = bucketsAPI.FindBucketByName(ctx, common.BucketFileEvents)
	if err != nil && err.Error() == "bucket '"+common.BucketFileEvents+"' not found" {
		orgDomain, err := orgsAPI.FindOrganizationByName(ctx, d.DBOrganization)
		if err != nil {
			return fmt.Errorf("failed to find organization: %w", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, common.BucketFileEvents)
		if err != nil {
			return fmt.Errorf("failed to create bucket: %w", err)
		}
	}

	_, err = bucketsAPI.FindBucketByName(ctx, common.BucketWorkerStatus)
	if err != nil && err.Error() == "bucket '"+common.BucketWorkerStatus+"' not found" {
		orgDomain, err := orgsAPI.FindOrganizationByName(ctx, d.DBOrganization)
		if err != nil {
			return fmt.Errorf("failed to find organization: %w", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, common.BucketWorkerStatus)
		if err != nil {
			return fmt.Errorf("failed to create bucket: %w", err)
		}
	}

	return nil
}
