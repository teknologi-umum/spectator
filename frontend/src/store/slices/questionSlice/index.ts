import { createSlice } from "@reduxjs/toolkit";
import type { InitialState } from "./types";

const initialState: InitialState = {
  currentQuestion: 0
};

export const questionSlice = createSlice({
  name: "editor",
  initialState,
  reducers: {
    // TODO(elianiva): this is temporary, remove it when we got the final logic
    //                 down (submit code to backend, if correct then next,
    //                 disable previous question, etc)
    prevQuestion: (state) => {
      if (state.currentQuestion <= 0) return;
      state.currentQuestion -= 1;
    },
    nextQuestion: (state) => {
      if (state.currentQuestion >= 5) return;
      state.currentQuestion += 1;
    }
  }
});

export const { prevQuestion, nextQuestion } = questionSlice.actions;

export default questionSlice.reducer;
