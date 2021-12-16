import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { InitialState } from "./types";

const initialState: InitialState = {
  stdNo: "",
  name: "",
  degreeYear: 0,
  gender: 0,
  major: 0,
  race: "",
  programmingGrade: 0,
  programmingExp: 0,
  programmingExercise: 0,
  programmingLanguage: ""
};

export const personalInfoSlice = createSlice({
  name: "personalInfo",
  initialState,
  reducers: {
    recordPersonalInfo: (
      state,
      action: PayloadAction<Partial<InitialState>>
    ) => {
      state = {
        ...state,
        ...action.payload
      };
    }
  }
});

export const { recordPersonalInfo } = personalInfoSlice.actions;

export default personalInfoSlice.reducer;
