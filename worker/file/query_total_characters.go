package file

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/rs/zerolog/log"
)

type TotalCharacters struct {
	SessionID string `json:"session_id" csv:"session_id"`
	Total     int64  `json:"total_characters" csv:"total_characters"`
}

func (d *Dependency) QueryTotalCharacters(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (TotalCharacters, error) {
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
		// Deletion
		"Delete", "Backspace",
	}

	keystrokeRows, err := queryAPI.Query(
		ctx,
		`from(bucket: `+strconv.Quote(common.BucketInputEvents)+`)
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == `+strconv.Quote(common.MeasurementKeystroke)+`)
		|> filter(fn: (r) => r["session_id"] == `+strconv.Quote(sessionID.String())+`)
		|> pivot(columnKey: ["_field"], rowKey: ["_time"], valueColumn: "_value")
		|> filter(fn: (r) => r["unrelated_key"] == false)
		|> keep(columns: ["_time", "_value", "_measurement", "key_char"])
		|> filter(fn: (r) => contains(value: r["key_char"],
										set: ["`+strings.Join(whitelist, `", "`)+`"]))
		|> sort(columns: ["_time"])`,
	)
	if err != nil {
		return TotalCharacters{}, fmt.Errorf("failed to query total characters: %w", err)
	}
	defer func() {
		err := keystrokeRows.Close()
		if err != nil {
			log.Err(err).Msg("closing keystrokeRows")
		}
	}()

	var totalCharacters int64

	for keystrokeRows.Next() {
		record := keystrokeRows.Record()

		keyChar, ok := record.ValueByKey("key_char").(string)
		if !ok {
			keyChar = ""
		}

		if keyChar == "Delete" || keyChar == "Backspace" {
			totalCharacters = totalCharacters - 1
			continue
		}

		totalCharacters = totalCharacters + 1
	}

	return TotalCharacters{
		SessionID: sessionID.String(),
		Total:     totalCharacters,
	}, nil
}
