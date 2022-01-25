package file

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type PersonalInfo struct {
	Type              string    `json:"type" csv:"-"`
	SessionID         string    `json:"session_id" csv:"session_id"`
	StudentNumber     string    `json:"student_number" csv:"student_number"`
	HoursOfPractice   int64     `json:"hours_of_practice" csv:"hours_of_experience"`
	YearsOfExperience int64     `json:"years_of_experience" csv:"years_of_experience"`
	FamiliarLanguages string    `json:"familiar_languages" csv:"familliar_languages"`
	Timestamp         time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryPersonalInfo(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (PersonalInfo, error) {

	tempPersonalInfo := PersonalInfo{}
	for _, x := range []string{"student_number", "hours_of_practice", "years_of_experience", "familiar_language"} {
		personalInfoRows, err := queryAPI.Query(
			ctx,
			`from(bucket: "`+d.BucketSessionEvents+`")
			|> range(start: 0)
			|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+`")
			|> filter(fn : (r) => r["_measurement"] == "personal_info")
			`,
		)
		if err != nil {
			return PersonalInfo{}, fmt.Errorf("failed to query personal info: %w", err)
		}

		var ok bool
		for personalInfoRows.Next() {

			rows := personalInfoRows.Record()

			switch x {
			case "student_number":
				tempPersonalInfo.StudentNumber, ok = rows.Value().(string)
				if !ok {
					tempPersonalInfo.StudentNumber = ""
				}
			case "hours_of_practice":
				tempPersonalInfo.HoursOfPractice, ok = rows.Value().(int64)
				if !ok {
					tempPersonalInfo.HoursOfPractice = 0
					// return PersonalInfo{}, fmt.Errorf("failed to parse hours of practice type")
				}
			case "years_of_experience":
				tempPersonalInfo.YearsOfExperience, ok = rows.Value().(int64)
				if !ok {
					tempPersonalInfo.YearsOfExperience = 0
					// return PersonalInfo{}, fmt.Errorf("failed to parse years of experience type")
				}
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

			tempPersonalInfo.SessionID, ok = rows.ValueByKey("session_id").(string)
			if !ok {
				tempPersonalInfo.SessionID = ""
			}
			tempPersonalInfo.Timestamp = rows.Time()
		}
	}

	return tempPersonalInfo, nil
}
