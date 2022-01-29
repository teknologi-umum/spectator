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
	examStartedRow, err := queryAPI.Query(
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
	defer examStartedRow.Close()

	var startTime int64
	var endTime int64

	if examStartedRow.Next() {
		startTime = examStartedRow.Record().Time().Unix()
	}

	examEndedRow, err := queryAPI.Query(
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
	defer examEndedRow.Close()

	if examEndedRow.Next() {
		endTime = examEndedRow.Record().Time().Unix()
	}

	// If the end time is 0, we check from the exam_forfeited measurement.
	if endTime == 0 {
		examForfeitedRow, err := queryAPI.Query(
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
		defer examForfeitedRow.Close()

		if examForfeitedRow.Next() {
			endTime = examForfeitedRow.Record().Time().Unix()
		}

		if endTime == 0 {
			return fmt.Errorf("session end time is not defined")
		}
	}

	// keystrokesIgnore contains the keys that might appear on the "key_char" that we don't
	// want to count into the resulting words per minute.
	var keystrokesIgnore = []string{"backspace", "delete", "insert", "pageup", "pagedown"}
	// wordsPerMinute contains the array of each minute's words per minute.
	// This can be used to calculate the average of all the words per minute.
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
			|> filter(fn: (r) => r["_measurement"] == "keystroke")
			|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
			|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
			|> sort(columns: ["_time"])`,
		)
		if err != nil {
			return fmt.Errorf("failed to query keystroke events: %w", err)
		}
		defer rows.Close()

		var currentWordCount uint32
		for rows.Next() {
			record := rows.Record()

			keyChar, ok := record.ValueByKey("key_char").(string)
			if !ok {
				return fmt.Errorf("failed to parse key_char data: %v", err)
			}

			if !contains(keystrokesIgnore, keyChar) {
				currentWordCount++
			}
		}

		wordsPerMinute = append(wordsPerMinute, currentWordCount)
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
