package funfact

import (
	"context"
	"fmt"
	"strings"
	"worker/common"
	"worker/status"

	"github.com/google/uuid"
)

func (d *Dependency) CalculateWordsPerMinute(ctx context.Context, sessionID uuid.UUID, result chan int64) error {
	d.Status.AppendState(ctx, sessionID, "calculate_wpm", status.StatePending)

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
	// Then, return a channel that took the average of all the words per minute.

	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// Get the value of the time that the user started and ended the session.
	examStartedRow, err := queryAPI.Query(
		ctx,
		`from (bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => (r["_measurement"] == "`+common.MeasurementExamStarted+`" and
			                  r["session_id"] == "`+sessionID.String()+`"))`,
	)
	if err != nil {
		d.Status.AppendState(ctx, sessionID, "calculate_wpm", status.StateFailed)
		return fmt.Errorf("failed to query session start time: %w", err)
	}
	defer examStartedRow.Close()

	var startTime int64
	if examStartedRow.Next() {
		startTime = examStartedRow.Record().Time().Unix()
	}

	// whitelist contains the keys that might appear on the "key_char" that we
	// want to count into the resulting words per minute.
	whitelist := []string{
		// Letters
		"KeyA", "KeyB", "KeyC", "KeyD", "KeyE", "KeyF", "KeyG", "KeyH", "KeyI", "KeyJ", "KeyK", "KeyL", "KeyM",
		"KeyN", "KeyO", "KeyP", "KeyQ", "KeyR", "KeyS", "KeyT", "KeyU", "KeyV", "KeyW", "KeyX", "KeyY", "KeyZ",
		// Numbers
		"Digit0", "Digit1", "Digit2", "Digit3", "Digit4", "Digit5", "Digit6", "Digit7", "Digit8", "Digit9",
		"Numpad0", "Numpad1", "Numpad2", "Numpad3", "Numpad4", "Numpad5", "Numpad6", "Numpad7", "Numpad8", "Numpad9",
		// Punctuation
		"Comma", "Period", "Semicolon", "Slash", "Backslash", "BracketLeft", "BracketRight", "Quote", "Backquote",
		"Minus", "Equal", "Subtract", "Add", "Multiply", "Divide", "Space",
		// Numpad Punctuation
		"NumpadAdd", "NumpadSubtract", "NumpadDecimal",
	}

	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketInputEvents+`")
			|> range(start: `+fmt.Sprintf("%d", startTime)+`)
			|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementKeystroke+`")
			|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
			|> pivot(columnKey: ["_field"], rowKey: ["_time"], valueColumn: "_value")
			|> filter(fn: (r) => r["unrelated_key"] == false)
			|> keep(columns: ["_time", "_value", "_measurement", "key_char"])
			|> filter(fn: (r) => contains(value: r["key_char"],
										  set: ["`+strings.Join(whitelist, `", "`)+`"]))
			|> sort(columns: ["_time"])
			|> elapsed(unit: 1ms, timeColumn: "_time", columnName: "elapsed")`,
	)
	if err != nil {
		d.Status.AppendState(ctx, sessionID, "calculate_wpm", status.StateFailed)
		return fmt.Errorf("failed to query keystroke events: %w", err)
	}
	defer rows.Close()

	// calculate wpm every short burst
	var totalKeystrokes []int64
	var currentFrame []int64
	for rows.Next() {
		record := rows.Record()
		// elapsed is the time delta between current keypress and previous keypress
		elapsed := record.ValueByKey("elapsed").(int64)

		// go to the next timeframe if the distance between keypress is more than 330ms
		if elapsed >= 330 && len(currentFrame) != 0 {
			// only accept at least 10 keystrokes per burst
			if len(currentFrame) > 10 {
				var characters int64
				start := currentFrame[0]
				end := currentFrame[len(currentFrame)-1]
				duration := (end - start) / 1000 // burst duration in seconds
				// less than 1 second means 1 second
				if duration == 0 {
					duration = 1
				}

				for i := 0; i < len(currentFrame); i++ {
					characters++
				}

				words := characters / 5

				totalKeystrokes = append(totalKeystrokes, (words/duration)*60)
			}
			// reset current frame when the distance between 2 keystrokes is too long
			currentFrame = []int64{}
			continue
		}

		currentFrame = append(currentFrame, record.Time().UnixMilli())
	}

	if len(totalKeystrokes) == 0 {
		// just send 0 which means the user didn't type enough
		result <- 0
		d.Status.AppendState(ctx, sessionID, "calculate_wpm", status.StateSuccess)
		return nil
	}

	var totalWpm int64
	var wpmDivisor int64
	for _, wpm := range totalKeystrokes {
		if wpm >= 20 {
			totalWpm += wpm
			wpmDivisor++
		}
	}

	// if the wordsSum is 0, just send it back.
	// the reason why is when we divide 0 with 0, it became NaN for some reason
	// and the final result will became a weird number like -9223372036854775808
	if totalWpm == 0 {
		result <- 0
		d.Status.AppendState(ctx, sessionID, "calculate_wpm", status.StateSuccess)
		return nil
	}

	totalWpm = totalWpm / wpmDivisor

	// Return the result here
	result <- totalWpm

	d.Status.AppendState(ctx, sessionID, "calculate_wpm", status.StateSuccess)
	return nil
}
