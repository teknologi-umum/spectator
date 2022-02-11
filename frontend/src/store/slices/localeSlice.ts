import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import type { Locale } from "@/models/Locale";

interface State {
  locale: Locale;
}

const initialState: State = {
  locale: "EN"
};

export const sessionSlice = createSlice({
  name: "session",
  initialState,
  reducers: {
    setLocale: (state, action: PayloadAction<Locale>) => {
      state.locale = action.payload;
    }
  }
});

export const { setLocale } = sessionSlice.actions;

export default sessionSlice.reducer;
