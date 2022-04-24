package file

import (
	"context"
	"fmt"
	"worker/common"

	"github.com/google/uuid"
)

func (d *Dependency) CheckIfSessionExists(ctx context.Context, sessionID uuid.UUID) (bool, error) {
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementPersonalInfoSubmitted+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"])`,
	)
	if err != nil {
		return false, fmt.Errorf("failed to query session exists: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		record := rows.Record()

		sessionId, ok := record.ValueByKey("session_id").(string)
		if !ok {
			return false, nil
		}

		if sessionID.String() == sessionId {
			return true, nil
		}
	}

	return false, nil
}
