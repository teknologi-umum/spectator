package file

import (
	"context"
	"fmt"
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
	var personalInfo PersonalInfo

	studentNumberRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => r["_measurement"] == "personal_info_submitted")
		|> filter(fn: (r) => r["_field"] == "student_number")
		|> sort(columns: ["_time"])
		|> group(columns: ["_time"])`,
	)
	if err != nil {
		return PersonalInfo{}, fmt.Errorf("failed to query personal info - student number: %w", err)
	}
	defer studentNumberRows.Close()

	var ok bool
	for studentNumberRows.Next() {
		record := studentNumberRows.Record()

		personalInfo.SessionID, ok = record.ValueByKey("session_id").(string)
		if !ok {
			personalInfo.SessionID = ""
		}
		personalInfo.Timestamp = record.Time()

		personalInfo.StudentNumber, ok = record.Value().(string)
		if !ok {
			// FIXME: add default value
		}
	}

	hoursOfPracticeRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => r["_measurement"] == "personal_info")
		|> filter(fn: (r) => r["_field"] == "hours_of_practice")
		|> sort(columns: ["_time"])
		|> group(columns: ["_time"])`,
	)
	if err != nil {
		return PersonalInfo{}, fmt.Errorf("failed to query personal info - hours of practice: %w", err)
	}
	defer hoursOfPracticeRows.Close()

	for hoursOfPracticeRows.Next() {
		rows := hoursOfPracticeRows.Record()

		personalInfo.HoursOfPractice, ok = rows.Value().(int64)
		if !ok {
			// FIXME: add default value
		}
	}

	yearsOfExperienceRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => r["_measurement"] == "personal_info")
		|> filter(fn: (r) => r["_field"] == "years_of_experience")
		|> sort(columns: ["_time"])
		|> group(columns: ["_time"])`,
	)
	if err != nil {
		return PersonalInfo{}, fmt.Errorf("failed to query personal info - years of experience: %w", err)
	}
	defer yearsOfExperienceRows.Close()

	for yearsOfExperienceRows.Next() {
		rows := yearsOfExperienceRows.Record()

		personalInfo.YearsOfExperience, ok = rows.Value().(int64)
		if !ok {
			// FIXME: add default value
		}
	}

	familiarLanguagesRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+d.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => r["_measurement"] == "personal_info")
		|> filter(fn: (r) => r["_field"] == "familiar_languages")
		|> sort(columns: ["_time"])
		|> group(columns: ["_time"])`,
	)
	if err != nil {
		return PersonalInfo{}, fmt.Errorf("failed to query personal info - familiar languages: %w", err)
	}
	defer familiarLanguagesRows.Close()

	for familiarLanguagesRows.Next() {
		rows := familiarLanguagesRows.Record()

		personalInfo.FamiliarLanguages, ok = rows.Value().(string)
		if !ok {
			// FIXME: add default value
		}
	}

	return personalInfo, nil
}
