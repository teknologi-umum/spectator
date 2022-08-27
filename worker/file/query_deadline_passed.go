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

// measurement: deadline_passed
type DeadlinePassed struct {
	Measurement string    `json:"_measurement" csv:"_measurement"`
	SessionId   string    `json:"session_id" csv:"session_id"`
	Timestamp   time.Time `json:"timestamp" csv:"timestamp"`
}

func (d *Dependency) QueryDeadlinePassed(ctx context.Context, queryAPI api.QueryAPI, sessionID uuid.UUID) (*DeadlinePassed, error) {
	deadlinePassedRows, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+common.BucketSessionEvents+`")
		|> range(start: 0)
		|> filter(fn: (r) => r["_measurement"] == "`+common.MeasurementDeadlinePassed+`" and r["session_id"] == "`+sessionID.String()+`")`,
	)
	if err != nil {
		return &DeadlinePassed{}, fmt.Errorf("failed to query deadline_passed: %w", err)
	}
	defer func() {
		err := deadlinePassedRows.Close()
		if err != nil {
			log.Err(err).Msg("closing deadlinePassedRows")
		}
	}()

	var outputDeadlinePassed DeadlinePassed

	for deadlinePassedRows.Next() {
		record := deadlinePassedRows.Record()

		if record.Time().Year() != 2022 {
			log.Warn().
				Str("current time from record.Time() is not 2022, it's ", strconv.Itoa(record.Time().Year())).
				Msg("invalid date on QueryDeadlinePassed")
		}

		outputDeadlinePassed = DeadlinePassed{
			Measurement: common.MeasurementDeadlinePassed,
			SessionId:   sessionID.String(),
			Timestamp:   record.Time(),
		}
	}

	return &outputDeadlinePassed, nil
}
