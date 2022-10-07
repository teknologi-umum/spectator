import { describe, expect } from "vitest";
import reducer, {
  setFontSize,
  setLanguage,
  setLockedToCurrentQuestion
} from "@/store/slices/editorSlice";
import type { EditorState } from "@/models/EditorState";

const initialState: EditorState = {
  deadlineUtc: null,
  questions: null,
  currentQuestionNumber: 0,
  lockedToCurrentQuestion: false,
  currentLanguage: "javascript",
  fontSize: 14,
  snapshotByQuestionNumber: {}
};

describe("Editor related state", (it) => {
  it("should return the initial state", () => {
    expect(reducer(undefined, { type: null })).toEqual(initialState);
  });

  it("should be able to set current editor language", () => {
    expect(reducer(initialState, setLanguage("cpp"))).toEqual({
      ...initialState,
      currentLanguage: "cpp",
      snapshotByQuestionNumber: {
        0: {
          language: "cpp",
          samTestResult: null,
          questionNumber: 0,
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
          testResults: null
        }
      }
    });
  });

  it("should be able to set current editor font size", () => {
    expect(reducer(initialState, setFontSize(18))).toEqual({
      ...initialState,
      fontSize: 18
    });
  });

  it("should be able to lock to the current question", () => {
    expect(reducer(initialState, setLockedToCurrentQuestion(true))).toEqual({
      ...initialState,
      lockedToCurrentQuestion: true
    });
  });
});
