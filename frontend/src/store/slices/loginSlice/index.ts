import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { InitialState } from "./types";

const initialState: InitialState = {
  password: ""
}

export const LoginSlice = createSlice({
  name: "login",
  initialState,
  reducers: {
    recordPassword: (state, action: PayloadAction<string>) => {
      state.password = action.payload;
    }
  }
});

export const { recordPassword } = LoginSlice.actions;

export default LoginSlice.reducer;