import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { InitialState, Languages } from "./types";

const initialState: InitialState = {
  currentLanguage: "javascript",
  code: "",
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
    }
  }
});

export const { changeFontSize, changeCurrentLanguage } = editorSlice.actions;

export default editorSlice.reducer;
