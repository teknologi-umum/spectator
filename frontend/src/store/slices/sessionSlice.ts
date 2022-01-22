import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface State {
  accessToken: string | null;
}

const initialState: State = {
  accessToken: null
};

export const sessionSlice = createSlice({
  name: "session",
  initialState,
  reducers: {
    setAccessToken: (state, action: PayloadAction<string>) => {
      state.accessToken = action.payload;
    },
  }
});

export const { setAccessToken } = sessionSlice.actions;

export default sessionSlice.reducer;