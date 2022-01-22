import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { PersonalInfo } from "@/models/PersonalInfo";

interface State {
  personalInfo: PersonalInfo | null;
}

const initialState: State = {
  personalInfo: null
};

export const personalInfoSlice = createSlice({
  name: "personalInfo",
  initialState,
  reducers: {
    setPersonalInfo: (state, action: PayloadAction<PersonalInfo>) => {
      state.personalInfo = action.payload;
    }
  }
});

export const { setPersonalInfo } = personalInfoSlice.actions;

export default personalInfoSlice.reducer;
