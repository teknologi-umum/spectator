package funfact

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (d *Dependency) CalculateDeletionRate(ctx context.Context, sessionID uuid.UUID, result chan float32) error {
	var deletionTotal int64 = 0
	var totalKeystrokes int64 = 0
	var ok bool

	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	deletionRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "coding_event_keystroke")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => (r["key_char"] == "backspace" or r["key_char"] == "delete"))
		|> count()
		|> yield(name: "count")`,
	)
	if err != nil {
		return fmt.Errorf("failed to query deletion rate: %w", err)
	}
	defer deletionRows.Close()

	for deletionRows.Next() {
		deletionTotal, ok = deletionRows.Record().Value().(int64)
		if !ok {
			return fmt.Errorf("failed to cast deletionTotal to int64")
		}
	}

	keystrokeTotalRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`")
		|> range(start: -1d)
		|> filter(fn: (r) => r["_measurement"] == "coding_event_keystroke")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => (r._field == "key_char" and r._value != ""))
		|> count()
		|> yield(name: "count")`,
	)
	if err != nil {
		return (err)
	}
	defer keystrokeTotalRows.Close()

	for keystrokeTotalRows.Next() {
		value, ok := keystrokeTotalRows.Record().Value().(int64)
		if !ok {
			return errors.New("fail to infer keystroke Total")
		}

		totalKeystrokes = value
	}

	result <- (float32(deletionTotal) / float32(totalKeystrokes))

	// SELECT semua KeystrokeEvent WHERE value = delete OR value = backspace
	// terus jumlahin
	// dah gitu doang.

	// Return the result here
	return nil
}
