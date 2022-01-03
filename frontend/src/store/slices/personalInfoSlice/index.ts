import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { InitialState } from "./types";

const initialState: InitialState = {
  studentId: "",
  programmingExp: 0,
  programmingExercise: 0,
  programmingLanguage: ""
};

export const personalInfoSlice = createSlice({
  name: "personalInfo",
  initialState,
  reducers: {
    recordPersonalInfo: (state, action: PayloadAction<InitialState>) => {
      const {
        studentId,
        programmingExp,
        programmingExercise,
        programmingLanguage
      } = action.payload;

      state.studentId = studentId;
      state.programmingExp = programmingExp;
      state.programmingExercise = programmingExercise;
      state.programmingLanguage = programmingLanguage;
    }
  }
});

export const { recordPersonalInfo } = personalInfoSlice.actions;

export default personalInfoSlice.reducer;
