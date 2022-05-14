import { describe, expect } from "vitest";
import reducer, { setFontSize, setLanguage } from "@/store/slices/editorSlice";
import type { EditorState } from "@/models/EditorState";

const initialState: EditorState = {
  deadlineUtc: null,
  questions: null,
  currentQuestionNumber: 1,
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
        1: {
          language: "cpp",
          questionNumber: 1,
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
});
