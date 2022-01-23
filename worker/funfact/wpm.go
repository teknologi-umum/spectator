package funfact

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (d *Dependency) CalculateWordsPerMinute(ctx context.Context, sessionID uuid.UUID, result chan uint32) error {
	// The formula to calculate words per minute is as follows:
	// SELECT all KeystrokeEvent, group by TIME, each TIME is windowed by 1 minute
	// for every 1 minute, calculate the total keystroke event and divide by 5.
	//
	// A quick note, that we can't use al the KeystrokeEvent input, because
	// calculating words per minute doesn't includes the key of backspaces,
	// delete, insert, pageup, pagedown, etc. So we'll have to filter it on our app.
	//
	// Now you've got the words per minute on that specific minute.
	// Then, move to the next minute and repeat the same process.
	//
	// Then, return a chanel that took the average of all the words per minute.

	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// Get the value of the time that the user started and ended the session.
	row, err := queryAPI.Query(
		ctx,
		`from (bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => 
			(r["_measurement"] == "exam_started" and r["session_id"] == "`+sessionID.String()+`"))
		|> yield()`,
	)
	if err != nil {
		return fmt.Errorf("failed to query session start time: %w", err)
	}
	defer row.Close()

	var startTime int64
	var endTime int64

	if row.Next() {
		startTime = row.Record().Time().Unix()
	}

	row, err = queryAPI.Query(
		ctx,
		`from (bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) =>
			(r["_measurement"] == "exam_ended" and r["session_id"] == "`+sessionID.String()+`"))
		|> yield()`,
	)
	if err != nil {
		return fmt.Errorf("failed to query session end time: %w", err)
	}
	defer row.Close()

	if row.Next() {
		endTime = row.Record().Time().Unix()
	}

	// If the end time is 0, we check from the exam_forfeited measurement.
	if endTime == 0 {
		row, err = queryAPI.Query(
			ctx,
			`from (bucket: "`+d.BucketSessionEvents+`")
			|> range(start: 0)
			|> filter(fn: (r) =>
				(r["_measurement"] == "exam_forfeited" and r["session_id"] == "`+sessionID.String()+`"))
			|> yield()`,
		)
		if err != nil {
			return fmt.Errorf("failed to query session forfeited time: %w", err)
		}
		defer row.Close()

		if row.Next() {
			endTime = row.Record().Time().Unix()
		}

		if endTime == 0 {
			return fmt.Errorf("session end time is not defined")
		}
	}

	var keystrokesIgnore = []string{"backspace", "delete", "insert", "pageup", "pagedown"}
	var wordsPerMinute []uint32

	// Find the delta between endTime and startTime in minute.
	delta := (endTime - startTime) / 60

	// Now we loop over the delta and calculate the words per minute.
	var i int64
	for i = 0; i < delta; i++ {
		rows, err := queryAPI.Query(
			ctx,
			`from(bucket: "`+d.BucketInputEvents+`")
			|> range(start: `+fmt.Sprintf("%d", startTime+int64(i)*60)+`, 
				stop: `+fmt.Sprintf("%d", startTime+int64(i+1)*60)+`)
			|> filter(fn: (r) => r["_measurement"] == "coding_event_keystroke")
			|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
			|> sort(columns: ["_time"])`,
		)
		if err != nil {
			return fmt.Errorf("failed to query keystroke events: %w", err)
		}
		defer rows.Close()

		var tablePosition int64
		var temporaryWords uint32
		var keystrokeData KeystrokeInput
		for rows.Next() {
			record := rows.Record()
			table, ok := record.ValueByKey("table").(int64)
			if !ok {
				table = 0
			}

			temporaryWords += 1
			switch record.Field() {
			case "key_char":
				keystrokeData.KeyChar, ok = record.Value().(string)
				if !ok {
					return fmt.Errorf("failed to parse keystroke data: %w", err)
				}
			case "unrelated_key":
				keystrokeData.UnrelatedKey, ok = record.Value().(bool)
				if !ok {
					return fmt.Errorf("failed to parse unrelated_key: %w", err)
				}
			default:
				continue
			}

			if table != 0 && table > tablePosition {
				// Only append data if the current keystroke data is not unrelated key
				if !keystrokeData.UnrelatedKey && !contains(keystrokesIgnore, keystrokeData.KeyChar) {
					wordsPerMinute = append(wordsPerMinute, temporaryWords)
					temporaryWords = 0
				}

				// Always update the table position
				tablePosition = table
			}
		}

		// Append the last value
		if len(wordsPerMinute) > 0 && temporaryWords != 0 {
			wordsPerMinute = append(wordsPerMinute, temporaryWords)
		}
	}

	// Check the wordsPerMinute length, if it's zero, we return an error
	// because it shouldn't be zero.
	if len(wordsPerMinute) == 0 {
		return fmt.Errorf("no keystroke events found")
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

// contains checks whether a string is in a slice of strings.
// It's case insensitive, meaning it will convert the string value
// into lowercase, then compare it to the corresponding
// string input.
func contains(s []string, e string) bool {
	eLower := strings.ToLower(e)
	for _, a := range s {
		if strings.ToLower(a) == eLower {
			return true
		}
	}

	return false
}
