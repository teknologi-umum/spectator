import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { HubConnectionState } from "@microsoft/signalr";

interface State {
  connectionState: HubConnectionState;
}

const initialState: State = {
  connectionState: HubConnectionState.Disconnected
};

export const signalRSlice = createSlice({
  name: "session",
  initialState,
  reducers: {
    setConnectionState: (state, action: PayloadAction<HubConnectionState>) => {
      state.connectionState = action.payload;
    }
  }
});

export const { setConnectionState } = signalRSlice.actions;

export default signalRSlice.reducer;
