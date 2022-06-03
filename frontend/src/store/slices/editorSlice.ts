import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { Language } from "@/models/Language";
import type { Question } from "@/models/Question";
import type { EditorState } from "@/models/EditorState";
import type { EditorSnapshot } from "@/models/EditorSnapshot";
import { defaultEditorSnapshot } from "@/store/entity/EditorSnapshot";

const initialState: EditorState = {
  deadlineUtc: null,
  questions: null,
  currentQuestionNumber: 1,
  lockedToCurrentQuestion: false,
  currentLanguage: "javascript",
  fontSize: 14,
  snapshotByQuestionNumber: {}
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
    setLockedToCurrentQuestion: (
      status: EditorState,
      action: PayloadAction<boolean>
    ) => {
      status.lockedToCurrentQuestion = action.payload;
    },
    setCurrentQuestion: (
      state: EditorState,
      action: PayloadAction<Question>
    ) => {
      if (!(action.payload.questionNumber in state.snapshotByQuestionNumber)) {
        state.snapshotByQuestionNumber[action.payload.questionNumber] =
          defaultEditorSnapshot(state);
      }

      state.currentQuestionNumber = action.payload.questionNumber;
      state.currentLanguage =
        state.snapshotByQuestionNumber[action.payload.questionNumber].language;
    },
    setLanguage: (state, action: PayloadAction<Language>) => {
      if (!(state.currentQuestionNumber in state.snapshotByQuestionNumber)) {
        state.snapshotByQuestionNumber[state.currentQuestionNumber] =
          defaultEditorSnapshot(state);
      }
      state.snapshotByQuestionNumber[state.currentQuestionNumber].language =
        action.payload;
      state.currentLanguage = action.payload;
    },
    setFontSize: (state, action: PayloadAction<number>) => {
      state.fontSize = action.payload;
    },
    setSolution: (state, action: PayloadAction<string>) => {
      if (!(state.currentQuestionNumber in state.snapshotByQuestionNumber)) {
        state.snapshotByQuestionNumber[state.currentQuestionNumber] =
          defaultEditorSnapshot(state);
      }

      state.snapshotByQuestionNumber[
        state.currentQuestionNumber
      ].solutionByLanguage[state.currentLanguage] = action.payload;
    },
    setSnapshot: (state, action: PayloadAction<EditorSnapshot>) => {
      state.snapshotByQuestionNumber[state.currentQuestionNumber] = {
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
      if (!(state.currentQuestionNumber in state.snapshotByQuestionNumber)) {
        state.snapshotByQuestionNumber[state.currentQuestionNumber] =
          defaultEditorSnapshot(state);
      }

      state.snapshotByQuestionNumber[state.currentQuestionNumber].scratchPad =
        action.payload;
    }
  }
});

export const {
  setCurrentQuestionNumber,
  setCurrentQuestion,
  setLockedToCurrentQuestion,
  setLanguage,
  setFontSize,
  setSolution,
  setScratchPad,
  setDeadlineAndQuestions,
  setSnapshot
} = editorSlice.actions;

export default editorSlice.reducer;
