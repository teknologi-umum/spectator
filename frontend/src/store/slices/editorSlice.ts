import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { Language } from "@/models/Language";
import type { Question } from "@/models/Question";
import type { EditorState } from "@/models/EditorState";
import { EditorSnapshot } from "@/models/EditorSnapshot";

const initialState: EditorState = {
  deadlineUtc: null,
  questions: null,
  currentQuestionNumber: 0,
  currentLanguage: "javascript",
  fontSize: 14,
  // TODO(elianiva): replace this with a data coming from redux
  snapshotByQuestionNumber: {
    1: {
      questionNumber: 1,
      language: "javascript" as Language,
      scratchPad: "",
      submissionAccepted: false,
      submissionRefactored: false,
      submissionSubmitted: true,
      solutionByLanguage: {
        javascript: "function foo() { return 1; }",
        java: "",
        c: "",
        cpp: "",
        python: "",
        php: ""
      },
      testResults: [
        { testNumber: 1, status: "Passing" },
        {
          testNumber: 2,
          status: "RuntimeError",
          stderr: "Trying to access a non-existent variable"
        },
        {
          testNumber: 3,
          status: "CompileError",
          stderr:
            "Failed to compile: error: 'yeet' was not declared in this scope"
        },
        {
          testNumber: 4,
          status: "Failing",
          expectedStdout: "2",
          actualStdout: "1"
        },
        {
          testNumber: 5,
          status: "Failing",
          expectedStdout: "{ \"foo\": \"bar\" }",
          actualStdout: "1"
        },
        { testNumber: 6, status: "Passing" }
      ]
    }
  }
};

export const editorSlice = createSlice({
  name: "editor",
  initialState,
  reducers: {
    setDeadlineAndQuestions: (
      state,
      action: PayloadAction<{ deadlineUtc: number; questions: Question[] }>
    ) => {
      state.deadlineUtc = action.payload.deadlineUtc;
      state.questions = action.payload.questions;
    },
    setCurrentQuestionNumber: (
      state: EditorState,
      action: PayloadAction<number>
    ) => {
      state.currentQuestionNumber = action.payload;
    },
    setCurrentQuestion: (
      state: EditorState,
      action: PayloadAction<Question>
    ) => {
      if (!(action.payload.questionNumber in state.snapshotByQuestionNumber)) {
        state.snapshotByQuestionNumber[action.payload.questionNumber] = {
          questionNumber: action.payload.questionNumber,
          language: state.currentLanguage,
          solutionByLanguage: {
            ...action.payload.templateByLanguage,
            [state.currentLanguage]:
              action.payload.templateByLanguage[state.currentLanguage]
          },
          scratchPad: "",
          submissionSubmitted: false,
          submissionAccepted: false,
          submissionRefactored: false,
          testResults: null
        };
      }
      state.currentQuestionNumber = action.payload.questionNumber;
      state.currentLanguage =
        state.snapshotByQuestionNumber[action.payload.questionNumber].language;
    },
    setLanguage: (state, action: PayloadAction<Language>) => {
      state.snapshotByQuestionNumber[state.currentQuestionNumber!].language =
        action.payload;
      state.currentLanguage = action.payload;
    },
    setFontSize: (state, action: PayloadAction<number>) => {
      state.fontSize = action.payload;
    },
    setSolution: (state, action: PayloadAction<string>) => {
      const solutionByLanguage =
        state.snapshotByQuestionNumber[state.currentQuestionNumber!]
          ?.solutionByLanguage;
      if (solutionByLanguage !== undefined) {
        state.snapshotByQuestionNumber[
          state.currentQuestionNumber!
        ].solutionByLanguage[state.currentLanguage] = action.payload;
      }
    },
    setSnapshot: (state, action: PayloadAction<EditorSnapshot>) => {
      state.snapshotByQuestionNumber[state.currentQuestionNumber!] = {
        language: action.payload.language,
        questionNumber: action.payload.questionNumber,
        scratchPad: action.payload.scratchPad,
        solutionByLanguage: action.payload.solutionByLanguage,
        submissionAccepted: action.payload.submissionAccepted,
        submissionRefactored: action.payload.submissionRefactored,
        submissionSubmitted: action.payload.submissionSubmitted,
        testResults: action.payload.testResults
      };
    },
    setScratchPad: (state, action: PayloadAction<string>) => {
      const solutionByLanguage =
        state.snapshotByQuestionNumber[state.currentQuestionNumber!]
          ?.scratchPad;
      if (solutionByLanguage !== undefined) {
        state.snapshotByQuestionNumber[
          state.currentQuestionNumber!
        ].scratchPad = action.payload;
      }
    }
  }
});

export const {
  setCurrentQuestionNumber,
  setCurrentQuestion,
  setLanguage,
  setFontSize,
  setSolution,
  setScratchPad,
  setDeadlineAndQuestions,
  setSnapshot
} = editorSlice.actions;

export default editorSlice.reducer;
