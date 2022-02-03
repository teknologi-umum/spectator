package funfact

import (
	"context"
	"fmt"
	"worker/common"

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
		`from(bucket: "`+common.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementKeystroke+`")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> filter(fn: (r) => r["key_char"] == "Backspace" or r["key_char"] == "Delete")`,
	)
	if err != nil {
		return fmt.Errorf("failed to query deletion rate: %w", err)
	}
	defer deletionRows.Close()

	for deletionRows.Next() {
		totalDeletion += 1
	}

	keystrokeTotalRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementKeystroke+`")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`,
	)
	if err != nil {
		return fmt.Errorf("failed to query keystroke total: %w", err)
	}
	defer keystrokeTotalRows.Close()

	for keystrokeTotalRows.Next() {
		totalKeystrokes += 1
	}

	// Avoiding NaN output
	if totalDeletion == 0 {
		result <- 0
		return nil
	}

	result <- totalDeletion / totalKeystrokes

	return nil
}
