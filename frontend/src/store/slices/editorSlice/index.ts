import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { InitialState, Languages, Solution } from "./types";

const initialState: InitialState = {
  currentLanguage: "javascript",
  fontSize: 14,
  solutions: []
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
        (solution) => solution.questionNo === action.payload.questionNo
      );

      if (idx > -1) state.solutions[idx] = action.payload;
      else state.solutions.concat(action.payload);
    }
  }
});

export const { changeFontSize, changeCurrentLanguage, setSolution } =
  editorSlice.actions;

export default editorSlice.reducer;
