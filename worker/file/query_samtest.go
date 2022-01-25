package file

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type SamTest struct {
	SessionID    string    `json:"session_id" csv:"session_id"`
	Type         string    `json:"type" csv:"-"`
	ArousedLevel int64     `json:"aroused_level" csv:"aroused_level"`
	PleasedLevel int64     `json:"pleased_level" csv:"pleased_level"`
	Timestamp    time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QuerySAMTest(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]SamTest, error) {
	samTestRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "sam_test_before")
		`,
	)
	if err != nil {
		return []SamTest{}, fmt.Errorf("failed to query sam test: %w", err)
	}

	outputSamTest := []SamTest{}
	tempSamTest := SamTest{}
	var tablePosition int64
	for samTestRows.Next() {
		rows := samTestRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		if rows.Field() == "aroused_level" {
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []SamTest{}, fmt.Errorf("failed to parse aroused level: %w", err)
			}
			tempSamTest.ArousedLevel = y
		}

		if rows.Field() == "pleased_level" {
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []SamTest{}, fmt.Errorf("failed to parse pleased level: %w", err)
			}
			tempSamTest.PleasedLevel = y
		}

		if d.IsDebug() {
			log.Println(rows.String())
			log.Printf("table %d\n", rows.Table())
		}

		if table != 0 && table > tablePosition {
			outputSamTest = append(outputSamTest, tempSamTest)
			tablePosition = table
		} else {
			var ok bool

			tempSamTest.SessionID, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempSamTest.SessionID = ""
			}
			tempSamTest.Timestamp = rows.Time()
		}
	}

	if len(outputSamTest) > 0 || tempSamTest.SessionID != "" {
		outputSamTest = append(outputSamTest, tempSamTest)
	}

	return outputSamTest, nil
}
