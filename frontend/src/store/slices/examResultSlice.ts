import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { ExamResult } from "@/stub/session";

export interface ExamResultState {
  examResult: ExamResult | null;
}

const initialState: ExamResultState = {
  examResult: null
};

export const examResultSlice = createSlice({
  name: "examResult",
  initialState,
  reducers: {
    setExamResult: (state, action: PayloadAction<ExamResult>) => {
      state.examResult = action.payload;
    }
  }
});

export const {
  setExamResult
} = examResultSlice.actions;

export default examResultSlice.reducer;
