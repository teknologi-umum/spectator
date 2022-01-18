package file

import (
	"context"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func (d *Dependency) QuerySAMTest(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]SamTest, error) {
	samTestRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "sam_test_before")
		`,
	)
	if err != nil {
		return []SamTest{}, err
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

		switch rows.Field() {
		case "aroused_level":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []SamTest{}, err
			}
			tempSamTest.ArousedLevel = y
		case "pleased_level":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []SamTest{}, err
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

	// ? : this part ask Reynaldi's i had no ideas.
	if len(outputSamTest) > 0 || tempSamTest.SessionID != "" {
		outputSamTest = append(outputSamTest, tempSamTest)
	}

	return outputSamTest, nil
}
