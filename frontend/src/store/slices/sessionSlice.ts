import type { SessionState } from "@/models/Session";
import { createSlice, type PayloadAction } from "@reduxjs/toolkit";

const initialState: SessionState = {
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
  setVideoDeviceId
} = sessionSlice.actions;

export default sessionSlice.reducer;
