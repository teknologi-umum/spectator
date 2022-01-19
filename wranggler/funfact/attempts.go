package funfact

import (
	"context"
	"fmt"
	"log"

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
		|> group(columns: ["question_id"])
		|> count()`,
	)
	if err != nil {
		return fmt.Errorf("failed to query submission attempts: %w", err)
	}

	// terus langsung return hasilnya
	// tapi bisa juga di group per question, jadi
	// misalnya untuk question #1, dia ada 5 attempt, question #2 ada 10 attempt
	// and so on so forth.

	// Return the result here
	if rows.Record() == nil {
		result <- uint32(0)
		return nil
	}

	value, ok := rows.Record().Value().(int64)
	if !ok {
		log.Println("[ERROR] casting value to int64")
		result <- uint32(0)
		return nil
	}

	result <- uint32(value)

	return nil
}
