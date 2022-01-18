package funfact

import (
	"context"

	"github.com/google/uuid"
)

func (d *Dependency) CalculateWordsPerMinute(ctx context.Context, sessionID uuid.UUID, result chan uint32) error {
	// Cara calculate WPM:
	// SELECT semua KeystrokeEvent, group by TIME, each TIME itu 1 menit
	// for every 1 minute, hitung total keystroke event itu,
	// terus dibagi dengan 5
	//
	// Itu baru dapet WPM per 1 menit itu.
	// Nanti mungkin bisa di store data nya jadi slice (per 1 menit,
	// ngga perlu specify menit keberapanya, karena slice pasti urut)
	// terus return ke channel hasil average dari semua menit yang ada

	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketInputEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "coding_event_keystroke")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> window(every: 1m)
		|> filter(fn: (r) => r["_field"] == "key_char")
		|> sort(columns: ["_time"])`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	var wordsPerMinute []uint32
	var tablePosition int64
	var temporaryWords uint32
	for rows.Next() {
		record := rows.Record()
		table, ok := record.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		temporaryWords += 1

		if table != 0 && table > tablePosition {
			wordsPerMinute = append(wordsPerMinute, temporaryWords)
			temporaryWords = 0
		}
	}

	// Append the last value
	if len(wordsPerMinute) > 0 && temporaryWords != 0 {
		wordsPerMinute = append(wordsPerMinute, temporaryWords)
	}

	var averageWpm uint32
	var wordsSum uint32
	for _, wpm := range wordsPerMinute {
		wordsSum += wpm / 5
	}

	averageWpm = wordsSum / uint32(len(wordsPerMinute))

	// Return the result here
	result <- averageWpm
	return nil
}
