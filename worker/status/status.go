// status provides a common functionality to declare the status
// of a certain sessionID
package status

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"worker/common"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// Dependency contains the dependency injection
// struct to be used on this package
type Dependency struct {
	DB             influxdb2.Client
	DBOrganization string
}

type State string

const (
	StateFailed  State = "FAILED"
	StateSuccess State = "SUCCESS"
	StatePending State = "PENDING"
)

var ErrEmptyFieldParameter = errors.New("empty field parameter")

func (d *Dependency) AppendState(ctx context.Context, sessionID uuid.UUID, what string, state State) error {
	if what == "" {
		return fmt.Errorf("%w: what", ErrEmptyFieldParameter)
	}
	writeAPI := d.DB.WriteAPI(d.DBOrganization, common.BucketWorkerStatus)

	point := influxdb2.NewPoint(
		strings.ToLower(what),
		map[string]string{
			"session_id": sessionID.String(),
		},
		map[string]interface{}{
			"state": string(state),
		},
		time.Now(),
	)

	writeAPI.WritePoint(point)

	return nil
}
