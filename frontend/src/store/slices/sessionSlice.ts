import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface SessionState {
  hasFinished: boolean;
  hasGivenUp: boolean;
  isWorkingOnSAM: boolean;
  sessionId: string | null;
  accessToken: string | null;
  firstSAMSubmitted: boolean;
  secondSAMSubmitted: boolean;
  hasPermission: boolean;
  deviceId: string | null;
  tourCompleted: {
    personalInfo: boolean;
    samTest: boolean;
    codingTest: boolean;
  };
}

const initialState: SessionState = {
  hasFinished: false,
  hasGivenUp: false,
  isWorkingOnSAM: false,
  accessToken: null,
  sessionId: null,
  firstSAMSubmitted: false,
  secondSAMSubmitted: false,
  hasPermission: false,
  deviceId: null,
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
    removeAccessToken: (state) => {
      state.accessToken = null;
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
    },
    allowVideoPermission: (state) => {
      state.hasPermission = true;
    },
    setVideoDeviceId: (state, action: PayloadAction<string>) => {
      state.deviceId = action.payload;
    },
    startSAMTest: (state) => {
      state.isWorkingOnSAM = true;
    },
    finishSAMTest: (state) => {
      state.isWorkingOnSAM = false;
    }
  }
});

export const {
  setAccessToken,
  removeAccessToken,
  markFirstSAMSubmitted,
  markSecondSAMSubmitted,
  markTourCompleted,
  setSessionId,
  removeSessionId,
  allowVideoPermission,
  setVideoDeviceId,
  startSAMTest,
  finishSAMTest
} = sessionSlice.actions;

export default sessionSlice.reducer;
