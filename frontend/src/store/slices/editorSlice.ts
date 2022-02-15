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
  snapshotByQuestionNumber: {}
};

export const editorSlice = createSlice({
  name: "editor",
  initialState,
  reducers: {
    setDeadlineAndQuestions: (state, action: PayloadAction<{ deadlineUtc: number, questions: Question[] }>) => {
      state.deadlineUtc = action.payload.deadlineUtc;
      state.questions = action.payload.questions;
    },
    setCurrentQuestion: (state: EditorState, action: PayloadAction<Question>) => {
      if (!(action.payload.questionNumber in state.snapshotByQuestionNumber)) {
        state.snapshotByQuestionNumber[action.payload.questionNumber] = {
          questionNumber: action.payload.questionNumber,
          language: state.currentLanguage,
          solutionByLanguage: {
            ...action.payload.templateByLanguage,
            [state.currentLanguage]: action.payload.templateByLanguage[state.currentLanguage]
          },
          scratchPad: "",
          submissionSubmitted: false,
          submissionAccepted: false,
          submissionRefactored: false,
          testResults: null
        };
      }
      state.currentQuestionNumber = action.payload.questionNumber;
      state.currentLanguage = state.snapshotByQuestionNumber[action.payload.questionNumber].language;
    },
    setLanguage: (state, action: PayloadAction<Language>) => {
      state.snapshotByQuestionNumber[state.currentQuestionNumber!].language = action.payload;
      state.currentLanguage = action.payload;
    },
    setFontSize: (state, action: PayloadAction<number>) => {
      state.fontSize = action.payload;
    },
    setSolution: (state, action: PayloadAction<string>) => {
      const solutionByLanguage = state.snapshotByQuestionNumber[state.currentQuestionNumber!]?.solutionByLanguage;
      if (solutionByLanguage !== undefined) {
        state.snapshotByQuestionNumber[state.currentQuestionNumber!].solutionByLanguage[state.currentLanguage] = action.payload;
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
      const solutionByLanguage = state.snapshotByQuestionNumber[state.currentQuestionNumber!]?.scratchPad;
      if (solutionByLanguage !== undefined) {
        state.snapshotByQuestionNumber[state.currentQuestionNumber!].scratchPad = action.payload;
      }
    }
  }
});

export const {
  setCurrentQuestion,
  setLanguage,
  setFontSize,
  setSolution,
  setScratchPad,
  setDeadlineAndQuestions,
  setSnapshot
} = editorSlice.actions;

export default editorSlice.reducer;
