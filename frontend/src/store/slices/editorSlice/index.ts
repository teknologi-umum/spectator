import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { InitialState } from "./types";

const initialState: InitialState = {
  currentLanguage: "javascript",
  code: "",
  fontSize: 16,
  solutions: []
};

export const editorSlice = createSlice({
  name: "editor",
  initialState,
  reducers: {
    changeFontSize: (state, action: PayloadAction<number>) => {
      state = {
        ...state,
        fontSize: action.payload
      };
    }
  }
});

export const { changeFontSize } = editorSlice.actions;

export default editorSlice.reducer;
