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
		|> group(columns: ["_time"])
		|> yield()`,
	)
	if err != nil {
		return fmt.Errorf("failed to query submission attempts: %w", err)
	}
	defer rows.Close()

	// terus langsung return hasilnya
	// tapi bisa juga di group per question, jadi
	// misalnya untuk question #1, dia ada 5 attempt, question #2 ada 10 attempt
	// and so on so forth.

	// Return the result here
	var output uint32
	var tablePosition int64
	for rows.Next() {
		table, ok := rows.Record().ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		if tablePosition == 0 || table > tablePosition {
			output++
			tablePosition = table
		}
	}

	result <- output

	return nil
}
