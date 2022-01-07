import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { InitialState, Languages, Solution, ScratchPad } from "./types";

const initialState: InitialState = {
  currentLanguage: "javascript",
  fontSize: 14,
  solutions: [],
  scratchPads: []
};

export const editorSlice = createSlice({
  name: "editor",
  initialState,
  reducers: {
    changeFontSize: (state, action: PayloadAction<number>) => {
      state.fontSize = action.payload;
    },
    changeCurrentLanguage: (state, action: PayloadAction<string>) => {
      state.currentLanguage = action.payload as Languages;
    },
    setSolution: (state, action: PayloadAction<Solution>) => {
      const idx = state.solutions.findIndex(
        (solution) =>
          solution.questionNo === action.payload.questionNo &&
          solution.language === action.payload.language
      );

      if (idx > -1) {
        state.solutions[idx] = action.payload;
      } else {
        state.solutions = state.solutions.concat(action.payload);
      }
    },
    setScratchPad: (state, action: PayloadAction<ScratchPad>) => {
      const idx = state.scratchPads.findIndex(
        (solution) => solution.questionNo === action.payload.questionNo
      );

      if (idx > -1) {
        state.scratchPads[idx] = action.payload;
      } else {
        state.scratchPads = state.scratchPads.concat(action.payload);
      }
    }
  }
});

export const {
  changeFontSize,
  changeCurrentLanguage,
  setSolution,
  setScratchPad
} = editorSlice.actions;

export default editorSlice.reducer;
