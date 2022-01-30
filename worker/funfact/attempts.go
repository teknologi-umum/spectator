package funfact

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (d *Dependency) CalculateSubmissionAttempts(ctx context.Context, sessionID uuid.UUID, result chan uint32) error {
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// number of question submission attempts
	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "code_test_attempt")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> count()
		|> yield()`,
	)
	if err != nil {
		return fmt.Errorf("failed to query submission attempts: %w", err)
	}
	defer rows.Close()

	var output uint32
	for rows.Next() {
		count, ok := rows.Record().Value().(int64)
		if !ok {
			count = 0
		}

		output += uint32(count)
	}

	result <- output

	return nil
}
