package main

import (
	"context"
	"fmt"
)

func (d *Dependency) PrepareBucket(ctx context.Context) error {
	bucketsAPI := d.DB.BucketsAPI()
	_, err := bucketsAPI.FindBucketByName(ctx, "log")
	if err != nil && err.Error() != "bucket 'log' not found" {
		return fmt.Errorf("finding bucket: %v", err)
	}

	if err != nil && err.Error() == "bucket 'log' not found" {
		organizationAPI := d.DB.OrganizationsAPI()
		orgDomain, err := organizationAPI.FindOrganizationByName(ctx, d.Org)
		if err != nil {
			return fmt.Errorf("finding organization: %v", err)
		}

		_, err = bucketsAPI.CreateBucketWithName(ctx, orgDomain, "log")
		if err != nil {
			return fmt.Errorf("creating bucket: %v", err)
		}
	}

	return nil
}
