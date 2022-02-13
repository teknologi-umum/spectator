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

	// MeasurementPersonalInfoSubmitted provides the name for
	// the measurement that holds the initial personal information
	// about a certain user.
	MeasurementPersonalInfoSubmitted = "personal_info_submitted"

	// MeasurementSessionStarted provides the name of the measurement
	// that is emitted when a session is started.
	MeasurementSessionStarted = "session_started"

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

	// MeasurementDeadlinePassed provides the name for which the user
	// has exceeded the coding test deadline.
	MeasurementDeadlinePassed = "deadline_passed"

	// MeasurementCodeTestAttempt provides the name for the measurement
	// that is emitted when a user attempts to solve the coding test.
	// This measurement doesn't acknowledge whether the user has a correct
	// answer. That is done by MeasurementSolutionRejected and
	// MeasurementSolutionAccepted.
	MeasurementCodeTestAttempt = "code_test_attempt"

	// MeasurementSolution provides the name for the solution measurement
	// type that is used for storing rejected or accepted test result.
	MeasurementSolutionRejected = "solution_rejected"
	MeasurementSolutionAccepted = "solution_accepted"

	// MeasurementExportedData provides the name for the measurement
	// that is meant to keep the resulting data link that is kept on MinIO
	MeasurementExportedData = "exported_data"

	// MeasurementLocaleSet provides the name for the measurement
	// that is meant to keep the locale set by the user.
	// For each changes, there will be a new event that is logged into
	// the InfluxDB database.
	MeasurementLocaleSet = "locale_set"

	// MeasurementFunfactProjection provides the projection
	// generated by the funfact domain. It contains information about
	// words per minute, deletion rate, and coding test attempt.
	MeasurementFunfactProjection = "funfact_projection"
)

// MouseButton is the mouse button that was clicked by the user such as left,
// right, and middle.
type MouseButton int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonRight
	MouseButtonMiddle
)