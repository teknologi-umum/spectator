import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { InitialState, Theme } from "./types";

const initialState: InitialState = {
  currentTheme: "light"
};

const appSlice = createSlice({
  name: "app",
  initialState,
  reducers: {
    setColorMode: (state, action: PayloadAction<Theme>) => {
      state.currentTheme = action.payload;
    }
  }
});

export const { setColorMode } = appSlice.actions;

export default appSlice.reducer;