import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface State {
  isCollapsed: boolean;
  questionTabIndex: number;
}

const initialState: State = {
  isCollapsed: true,
  questionTabIndex: 0
};

export const codingTestSlice = createSlice({
  name: "codingTest",
  initialState,
  reducers: {
    toggleSideBar: (state) => {
      state.isCollapsed = !state.isCollapsed;
    },
    setQuestionTabIndex: (state, action: PayloadAction<"question" | "result">) => {
      if (action.payload === "question") {
        state.questionTabIndex = 0;
        return;
      }

      if (action.payload === "result") {
        state.questionTabIndex = 1;
        return;
      }

      console.error("KOK BISA SAMPE SINI????");
    }
  }
});

export const { toggleSideBar, setQuestionTabIndex } = codingTestSlice.actions;

export default codingTestSlice.reducer;
