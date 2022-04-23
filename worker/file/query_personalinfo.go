package file

import (
	"context"
	"fmt"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type PersonalInfo struct {
	Measurement       string    `json:"_measurement" csv:"_measurement"`
	SessionID         string    `json:"session_id" csv:"session_id"`
	StudentNumber     string    `json:"student_number" csv:"student_number"`
	HoursOfPractice   int64     `json:"hours_of_practice" csv:"hours_of_experience"`
	YearsOfExperience int64     `json:"years_of_experience" csv:"years_of_experience"`
	FamiliarLanguages string    `json:"familiar_languages" csv:"familliar_languages"`
	WalletNumber      string    `json:"wallet_number" csv:"wallet_number"`
	Timestamp         time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryPersonalInfo(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (*PersonalInfo, error) {
	var personalInfo PersonalInfo

	rows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementPersonalInfoSubmitted+`")
		|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> sort(columns: ["_time"])`,
	)
	if err != nil {
		return &PersonalInfo{}, fmt.Errorf("failed to query personal info - student number: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		record := rows.Record()

		sessionId, ok := record.ValueByKey("session_id").(string)
		if !ok {
			sessionId = ""
		}

		studentNumber, ok := record.ValueByKey("student_number").(string)
		if !ok {
			studentNumber = ""
		}

		hoursOfPractice, ok := record.ValueByKey("hours_of_practice").(int64)
		if !ok {
			hoursOfPractice = 0
		}

		yearsOfExperience, ok := record.ValueByKey("years_of_experience").(int64)
		if !ok {
			yearsOfExperience = 0
		}

		familiarLanguages, ok := record.ValueByKey("familiar_languages").(string)
		if !ok {
			familiarLanguages = ""
		}

		walletNumber, ok := record.ValueByKey("wallet_number").(string)
		if !ok {
			walletNumber = ""
		}

		personalInfo = PersonalInfo{
			Measurement:       common.MeasurementPersonalInfoSubmitted,
			SessionID:         sessionId,
			StudentNumber:     studentNumber,
			HoursOfPractice:   hoursOfPractice,
			YearsOfExperience: yearsOfExperience,
			FamiliarLanguages: familiarLanguages,
			WalletNumber:      walletNumber,
			Timestamp:         time.Now(),
		}
	}

	return &personalInfo, nil
}
