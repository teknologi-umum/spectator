import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface State {
  accessToken: string | null;
  firstSAMSubmitted: boolean;
  secondSAMSubmitted: boolean;
}

const initialState: State = {
  accessToken: null,
  firstSAMSubmitted: false,
  secondSAMSubmitted: false
};

export const sessionSlice = createSlice({
  name: "session",
  initialState,
  reducers: {
    setAccessToken: (state, action: PayloadAction<string>) => {
      state.accessToken = action.payload;
    },
    markFirstSAMSubmitted: (state) => {
      state.firstSAMSubmitted = true;
    },
    markSecondSAMSubmitted: (state) => {
      state.secondSAMSubmitted = true;
    }
  }
});

export const { setAccessToken, markFirstSAMSubmitted, markSecondSAMSubmitted } = sessionSlice.actions;

export default sessionSlice.reducer;