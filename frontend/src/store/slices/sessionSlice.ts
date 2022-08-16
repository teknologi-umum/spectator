import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface SessionState {
  sessionId: string | null;
  accessToken: string | null;
  firstSAMSubmitted: boolean;
  secondSAMSubmitted: boolean;
  hasPermission: boolean;
  deviceId: string | null;
  videoStream: MediaStream | null;
  tourCompleted: {
    personalInfo: boolean;
    samTest: boolean;
    codingTest: boolean;
  };
}

const initialState: SessionState = {
  accessToken: null,
  sessionId: null,
  firstSAMSubmitted: false,
  secondSAMSubmitted: false,
  hasPermission: false,
  deviceId: null,
  videoStream: null,
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
    setVideoStream: (state, action: PayloadAction<MediaStream | null>) => {
      state.videoStream = action.payload;
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
  setVideoStream
} = sessionSlice.actions;

export default sessionSlice.reducer;
