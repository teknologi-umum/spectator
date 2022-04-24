import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { PersonalInfo } from "@/models/PersonalInfo";

const initialState: PersonalInfo = {
  studentNumber: "",
  yearsOfExperience: 0,
  hoursOfPractice: 0,
  familiarLanguages: "",
  walletNumber: "",
  walletType: "grabpay"
};

export const personalInfoSlice = createSlice({
  name: "personalInfo",
  initialState,
  reducers: {
    setPersonalInfo: (state, action: PayloadAction<PersonalInfo>) => {
      state.studentNumber = action.payload.studentNumber;
      state.yearsOfExperience = action.payload.yearsOfExperience;
      state.hoursOfPractice = action.payload.hoursOfPractice;
      state.familiarLanguages = action.payload.familiarLanguages;
      state.walletNumber = action.payload.walletNumber;
      state.walletType = action.payload.walletType;
    }
  }
});

export const { setPersonalInfo } = personalInfoSlice.actions;

export default personalInfoSlice.reducer;
