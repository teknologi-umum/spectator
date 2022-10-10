package file

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"worker/common"

	"github.com/google/uuid"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/rs/zerolog/log"
)

type PersonalInfo struct {
	Measurement       string    `json:"_measurement" csv:"_measurement"`
	SessionID         string    `json:"session_id" csv:"session_id"`
	Email             string    `json:"email" csv:"email"`
	Age               int64     `json:"age" csv:"age"`
	Gender            string    `json:"gender" csv:"gender"`
	Nationality       string    `json:"nationality" csv:"nationality"`
	StudentNumber     string    `json:"student_number" csv:"student_number"`
	HoursOfPractice   int64     `json:"hours_of_practice" csv:"hours_of_experience"`
	YearsOfExperience int64     `json:"years_of_experience" csv:"years_of_experience"`
	FamiliarLanguages string    `json:"familiar_languages" csv:"familliar_languages"`
	WalletNumber      string    `json:"wallet_number" csv:"wallet_number"`
	WalletType        string    `json:"wallet_type" csv:"wallet_type"`
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
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Err(err).Msg("closing personalInfoRows")
		}
	}()

	for rows.Next() {
		record := rows.Record()

		if record.Time().Year() != 2022 {
			log.Warn().
				Str("current time from record.Time() is not 2022, it's ", strconv.Itoa(record.Time().Year())).
				Msg("invalid date on QueryPersonalInfo")
		}

		sessionId, ok := record.ValueByKey("session_id").(string)
		if !ok {
			sessionId = ""
		}

		studentNumber, ok := record.ValueByKey("student_number").(string)
		if !ok {
			studentNumber = ""
		}

		email, ok := record.ValueByKey("email").(string)
		if !ok {
			email = ""
		}

		age, ok := record.ValueByKey("age").(int64)
		if !ok {
			age = 0
		}

		gender, ok := record.ValueByKey("gender").(string)
		if !ok {
			gender = ""
		}

		nationality, ok := record.ValueByKey("nationality").(string)
		if !ok {
			nationality = ""
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

		walletType, ok := record.ValueByKey("wallet_type").(string)
		if !ok {
			walletType = ""
		}

		personalInfo = PersonalInfo{
			Measurement:       common.MeasurementPersonalInfoSubmitted,
			SessionID:         sessionId,
			Email:             email,
			Age:               age,
			Gender:            gender,
			Nationality:       nationality,
			StudentNumber:     studentNumber,
			HoursOfPractice:   hoursOfPractice,
			YearsOfExperience: yearsOfExperience,
			FamiliarLanguages: familiarLanguages,
			WalletNumber:      walletNumber,
			WalletType:        walletType,
			Timestamp:         record.Time(),
		}
	}

	return &personalInfo, nil
}
