import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface State {
  sessionId: string | null;
  accessToken: string | null;
  firstSAMSubmitted: boolean;
  secondSAMSubmitted: boolean;
  tourCompleted: {
    personalInfo: boolean;
    samTest: boolean;
    codingTest: boolean;
  };
}

const initialState: State = {
  accessToken: null,
  sessionId: null,
  firstSAMSubmitted: false,
  secondSAMSubmitted: false,
  tourCompleted: {
    personalInfo: false,
    samTest: false,
    codingTest: false
  }
};

export const sessionSlice = createSlice({
  name: "session",
  initialState,
  reducers: {
    setSessionId: (state, action: PayloadAction<string>) => {
      state.sessionId = action.payload;
    },
    removeSessionId: (state) => {
      state.sessionId = null;
    },
    setAccessToken: (state, action: PayloadAction<string>) => {
      state.accessToken = action.payload;
    },
    markFirstSAMSubmitted: (state) => {
      state.firstSAMSubmitted = true;
    },
    markSecondSAMSubmitted: (state) => {
      state.secondSAMSubmitted = true;
    },
    markTourCompleted: (
      state,
      action: PayloadAction<"personalInfo" | "samTest" | "codingTest">
    ) => {
      state.tourCompleted[action.payload] = true;
    }
  }
});

export const {
  setAccessToken,
  markFirstSAMSubmitted,
  markSecondSAMSubmitted,
  markTourCompleted,
  setSessionId,
  removeSessionId
} = sessionSlice.actions;

export default sessionSlice.reducer;
