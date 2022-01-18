package file

import (
	"context"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func (d *Dependency) QueryPersonalInfo(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) ([]PersonalInfo, error) {
	personalInfoRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn : (r) => r["_measurement"] == "personal_info")
		`,
	)
	if err != nil {
		return []PersonalInfo{}, err
	}

	outputPersonalInfo := []PersonalInfo{}
	tempPersonalInfo := PersonalInfo{}
	var tablePosition int64
	for personalInfoRows.Next() {
		// TODO: mabok
		rows := personalInfoRows.Record()
		table, ok := rows.ValueByKey("table").(int64)
		if !ok {
			table = 0
		}

		switch rows.Field() {
		case "student_number":
			tempPersonalInfo.StudentNumber, ok = rows.Value().(string)
			if !ok {
				tempPersonalInfo.StudentNumber = ""
			}
		case "hours_of_practice":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []PersonalInfo{}, err
			}
			tempPersonalInfo.HoursOfPractice = y
		case "years_of_experience":
			y, err := strconv.ParseInt(rows.Value().(string), 10, 64)
			if err != nil {
				return []PersonalInfo{}, err
			}
			tempPersonalInfo.YearsOfExperience = y
		case "familiar_language":
			tempPersonalInfo.FamiliarLanguages, ok = rows.Value().(string)
			if !ok {
				tempPersonalInfo.FamiliarLanguages = ""
			}
		}

		if d.IsDebug() {
			log.Println(rows.String())
			log.Printf("table %d\n", rows.Table())
		}

		if table != 0 && table > tablePosition {
			outputPersonalInfo = append(outputPersonalInfo, tempPersonalInfo)
			tablePosition = table
		} else {
			var ok bool

			tempPersonalInfo.SessionID, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempPersonalInfo.SessionID = ""
			}
			tempPersonalInfo.Timestamp = rows.Time()
		}
	}

	// ? : this part ask Reynaldi's i had no ideas.
	if len(outputPersonalInfo) > 0 || tempPersonalInfo.SessionID != "" {
		outputPersonalInfo = append(outputPersonalInfo, tempPersonalInfo)
	}
	return outputPersonalInfo, nil
}
