import type { EditorState } from "@/models/EditorState";

export function defaultEditorSnapshot<StateDraft extends EditorState>(
  state: StateDraft
) {
  return {
    language: state.currentLanguage,
    questionNumber: state.currentQuestionNumber,
    scratchPad: "",
    solutionByLanguage: {
      javascript: "",
      php: "",
      java: "",
      python: "",
      c: "",
      cpp: ""
    },
    submissionAccepted: false,
    submissionSubmitted: false,
    submissionRefactored: false,
    testResults: null,
    samTestResult: null
  };
}
