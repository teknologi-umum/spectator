package file

// UserEvents provides the struct for collecting data regarding
// a specific user.
type UserEvents struct {
	PersonalInfo                    *PersonalInfo          `json:"personal_info" csv:"personal_info"`
	SelfAssessmentManekinBeforeTest *SelfAssessmentManekin `json:"self_assessment_manekin_before_test" csv:"self_assessment_manekin_before_test"`
	SelfAssessmentManekinAfterTest  *SelfAssessmentManekin `json:"self_assessment_manekin_after_test" csv:"self_assessment_manekin_after_test"`
	ExamStarted                     *ExamStarted           `json:"exam_started" csv:"exam_started"`
	ExamEnded                       *ExamEvent             `json:"exam_ended" csv:"exam_ended"`
	ExamForfeited                   *ExamEvent             `json:"exam_forfeited" csv:"exam_forfeited"`
	ExamIDEReloaded                 *[]ExamEvent           `json:"exam_ide_reloaded" csv:"exam_ide_reloaded"`
	Funfact                         *Funfact               `json:"funfact" csv:"funfact"`
	DeadlinePassed                  *DeadlinePassed        `json:"deadline_passed" csv:"deadline_passed"`
}

type KeystrokeEvents struct {
	Keystroke *[]Keystroke `json:"keystroke" csv:"keystroke"`
}

type MouseEvents struct {
	MouseClick            *[]MouseClick            `json:"mouse_click" csv:"mouse_click"`
	MouseMoved            *[]MouseMovement         `json:"mouse_moved" csv:"mouse_moved"`
	MouseScrolled         *[]MouseScrolled         `json:"mouse_scrolled" csv:"mouse_scrolled"`
	MouseDistanceTraveled *[]MouseDistanceTraveled `json:"mouse_distance_Traveled" csv:"mouse_distance_Traveled"`
}

type SolutionEvents struct {
	SolutionAccepted *[]Solution `json:"solution_accepted" csv:"solution_accepted"`
	SolutionRejected *[]Solution `json:"solution_rejected" csv:"solution_rejected"`
}
