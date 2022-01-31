import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { Language } from "@/models/Language";
import type { Question } from "@/models/Question";
import type { EditorState } from "@/models/EditorState";

const initialState: EditorState = {
  deadlineUtc: null,
  questions: null,
  currentQuestionNumber: null,
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
          submissionAccepted: null,
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
      state.snapshotByQuestionNumber[state.currentQuestionNumber!].solutionByLanguage[state.currentLanguage] = action.payload;
    },
    setScratchPad: (state, action: PayloadAction<string>) => {
      state.snapshotByQuestionNumber[state.currentQuestionNumber!].scratchPad = action.payload;
    }
  }
});

export const {
  setCurrentQuestion,
  setLanguage,
  setFontSize,
  setSolution,
  setScratchPad,
  setDeadlineAndQuestions
} = editorSlice.actions;

export default editorSlice.reducer;
