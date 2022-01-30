package file

import (
	"worker/logger"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/minio/minio-go/v7"
)

// MeasurementMouse provides the name for the mouse measurement
// type that is used for any mouse related events.
type MeasurementMouse string

const (
	MeasurementMouseDown     MeasurementMouse = "mouse_down"
	MeasurementMouseUp       MeasurementMouse = "mouse_up"
	MeasurementMouseMoved    MeasurementMouse = "mouse_moved"
	MeasurementMouseScrolled MeasurementMouse = "mouse_scrolled"
)

// MeasurementSAM provides the name for the sam measurement
// type that is used for any sam related events.
type MeasurementSAM string

const (
	MeasurementBeforeExamSAMSubmitted MeasurementSAM = "before_exam_sam_submitted"
	MeasurementAfterExamSAMSubmitted  MeasurementSAM = "after_exam_sam_submitted"
)

// MeasurementExam provides the name for the exam measurement
// type that is used for any exam related events.
type MeasurementExam string

const (
	MeasurementExamStarted     MeasurementExam = "exam_started"
	MeasurementExamEnded       MeasurementExam = "exam_ended"
	MeasurementExamForfeited   MeasurementExam = "exam_forfeited"
	MeasurementExamIDEReloaded MeasurementExam = "exam_ide_reloaded"
)

// MouseButton is the mouse button that was clicked by the user such as left,
// right, and middle.
type MouseButton int

const (
	Left MouseButton = iota
	Right
	Middle
)

// Dependency contains the dependency injection
// to be used on this package.
type Dependency struct {
	Environment         string
	DB                  influxdb2.Client
	Bucket              *minio.Client
	DBOrganization      string
	Logger              *logger.Logger
	LoggerToken         string
	BucketInputEvents   string
	BucketSessionEvents string
	BucketFileEvents    string
}
