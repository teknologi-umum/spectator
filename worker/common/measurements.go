package common

// Measurement file provides the common measurement variable names
// to be used against the InfluxDB query.

const (
	// MeasurementKeystroke provides the measurement name for the keystroke
	// event. This is the event that is emitted when a user types a key.
	MeasurementKeystroke = "keystroke"

	// MeasurementMouse provides the name for the mouse measurement
	// type that is used for any mouse related events.
	MeasurementMouseDown     = "mouse_down"
	MeasurementMouseUp       = "mouse_up"
	MeasurementMouseMoved    = "mouse_moved"
	MeasurementMouseScrolled = "mouse_scrolled"

	MeasurementWindowSized = "window_sized"

	// MeasurementSAM provides the name for the sam measurement
	// type that is used for any sam related events.
	MeasurementBeforeExamSAMSubmitted = "before_exam_sam_submitted"
	MeasurementAfterExamSAMSubmitted  = "after_exam_sam_submitted"

	// MeasurementExam provides the name for the exam measurement
	// type that is used for any exam related events.
	MeasurementExamStarted     = "exam_started"
	MeasurementExamEnded       = "exam_ended"
	MeasurementExamForfeited   = "exam_forfeited"
	MeasurementExamIDEReloaded = "exam_ide_reloaded"
	MeasurementExamPassed      = "exam_passed"

	// MeasurementSolution provides the name for the solution measurement
	// type that is used for storing rejected or accepted test result.
	MeasurementSolutionRejected = "solution_rejected"
	MeasurementSolutionAccepted = "solution_accepted"

	// MeasurementExportedData provides the name for the measurement
	// that is meant to keep the resulting data link that is kept on MinIO
	MeasurementExportedData = "exported_data"

	MeasurementDeadlinePassed = "deadline_passed"

	MeasurementLocaleSet = "locale_set"

	MeasurementPersonalInfoSubmitted = "personal_info_submitted"

	MeasurementSessionStarted = "session_started"

	MeasurementFunfactProjection = "funfact_projection"

	MeasurementCodeTestAttempt = "code_test_attempt"
)

// MouseButton is the mouse button that was clicked by the user such as left,
// right, and middle.
type MouseButton int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonRight
	MouseButtonMiddle
)
