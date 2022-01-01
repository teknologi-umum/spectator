import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { InitialState } from "./types";

const initialState: InitialState = {
  jwt: "",
  jwtPayload: {
    studentId: "",
    iat: 0,
    exp: 0
  },
  hasFinished: false
};

export const jwtSlice = createSlice({
  name: "jwt",
  initialState,
  reducers: {
    setJwt: (state, action: PayloadAction<string>) => {
      const jwtPayload = JSON.parse(window.atob(action.payload.split(".")[1]));
      state.jwt = action.payload;
      state.jwtPayload = jwtPayload;
      state.hasFinished = false;
    },
    finishSession: (state) => {
      state.hasFinished = true;
    }
  }
});

export const { setJwt, finishSession } = jwtSlice.actions;

export default jwtSlice.reducer;
