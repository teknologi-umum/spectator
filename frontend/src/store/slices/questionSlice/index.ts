import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import type { InitialState } from "./types";

const initialState: InitialState = {
  currentQuestion: 0,
  results: []
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
    },
    setQuestion: (state, action: PayloadAction<number>) => {
      state.currentQuestion = action.payload;
    }
  }
});

export const { prevQuestion, nextQuestion, setQuestion } =
  questionSlice.actions;

export default questionSlice.reducer;
