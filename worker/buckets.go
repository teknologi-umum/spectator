package main

import (
	"context"
	"fmt"
)

func (d *Dependency) prepareBuckets(ctx context.Context) error {
	bucketsAPI := d.DB.BucketsAPI()
	orgsAPI := d.DB.OrganizationsAPI()

	_, err := bucketsAPI.FindBucketByName(ctx, BucketInputStatisticEvents)
	if err != nil && err.Error() != "bucket '"+BucketInputStatisticEvents+"' not found" {
		orgDomain, err := orgsAPI.FindOrganizationByName(ctx, d.DBOrganization)
		if err != nil {
			return fmt.Errorf("failed to find organization: %v", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, BucketInputStatisticEvents)
		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}
	}

	_, err = bucketsAPI.FindBucketByName(ctx, BucketFileEvents)
	if err != nil && err.Error() != "bucket '"+BucketFileEvents+"' not found" {
		orgDomain, err := orgsAPI.FindOrganizationByName(ctx, d.DBOrganization)
		if err != nil {
			return fmt.Errorf("failed to find organization: %v", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, BucketFileEvents)
		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}
	}

	return nil
}
