package funfact

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (d *Dependency) CalculateDeletionRate(ctx context.Context, sessionID uuid.UUID, result chan float32) error {
	// Formula to calculate deletion rate:
	//
	// SELECT all KeystrokeEvent WHERE value = delete OR value = backspace
	// then sum it. That's it.

	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	var totalDeletion float32
	var totalKeystrokes float32

	deletionRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "keystroke")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => (
			(r["_field"] == "key_char" and r["_value"] == "Backspace") 
			or 
			(r["_field"] == "key_char" and r["_value"] == "Delete")))
		|> group(columns: ["_time"])
		|> yield()`,
	)
	if err != nil {
		return fmt.Errorf("failed to query deletion rate: %w", err)
	}
	defer deletionRows.Close()

	var tablePosition int64
	for deletionRows.Next() {
		table, ok := deletionRows.Record().ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		if tablePosition == 0 || table > tablePosition {
			totalDeletion++
			tablePosition = table
		}
	}

	keystrokeTotalRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "keystroke")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => r["_field"] == "key_char")
		|> group(columns: ["_time"])
		|> yield()`,
	)
	if err != nil {
		return fmt.Errorf("failed to query keystroke total: %w", err)
	}
	defer keystrokeTotalRows.Close()

	tablePosition = 0
	for keystrokeTotalRows.Next() {
		table, ok := keystrokeTotalRows.Record().ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		if tablePosition == 0 || table > tablePosition {
			totalKeystrokes++
			tablePosition = table
		}
	}

	// Avoiding NaN output
	if totalDeletion == 0 {
		result <- 0
		return nil
	}

	result <- totalDeletion / totalKeystrokes

	// Return the result here
	return nil
}
