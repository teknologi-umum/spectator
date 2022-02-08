import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { Theme } from "@/models/Theme";

interface State {
  currentTheme: Theme;
}

const initialState: State = {
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