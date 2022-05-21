import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { SecretPage } from "../../models/SecretPage";

const initialState: SecretPage = {
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