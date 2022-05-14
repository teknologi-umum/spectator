import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { Theme } from "@/models/Theme";

interface State {
  currentTheme: Theme;
}

const initialState: State = {
  currentTheme: "light"
};

const themeSlice = createSlice({
  name: "app",
  initialState,
  reducers: {
    setColorMode: (state, action: PayloadAction<Theme>) => {
      state.currentTheme = action.payload;
    }
  }
});

export const { setColorMode } = themeSlice.actions;

export default themeSlice.reducer;
