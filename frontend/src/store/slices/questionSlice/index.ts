import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { InitialState } from "./types";

const questionInit = {
  mark: 0,
  mouseEngagementScore: 0,
  timeUsed: 0,
};

const initialState: InitialState = {
  nextQuestion: 0,
  question1: questionInit,
  panas1: 0,
  question2: questionInit,
  panas2: 0,
  question3: questionInit,
  panas3: 0,
  question4: questionInit,
  panas4: 0,
  question5: questionInit,
  panas5: 0,
};

export const questionSlice = createSlice({
  name: "question",
  initialState,
  reducers: {
    setNextQuestion: (state, action: PayloadAction<number>) => {
      state.nextQuestion = action.payload;
    },
    submitQuestion: (
      state,
      action: PayloadAction<{
        question:
          | "question1"
          | "question2"
          | "question3"
          | "question4"
          | "question5";
        submission?: Partial<typeof questionInit>;
      }>
    ) => {
      const { question, submission } = action.payload;

      if (submission) {
        state[question] = { ...questionInit, ...submission };
      } else {
        state[question] = questionInit;
      }
    },
    submitPanas: (
      state,
      action: PayloadAction<{
        panas: "panas1" | "panas2" | "panas3" | "panas4" | "panas5";
        score?: number;
      }>
    ) => {
      const { panas, score } = action.payload;
      state[panas] = score ?? 0;
    },
  },
});

export const { setNextQuestion, submitQuestion, submitPanas } =
  questionSlice.actions;

export default questionSlice.reducer;
