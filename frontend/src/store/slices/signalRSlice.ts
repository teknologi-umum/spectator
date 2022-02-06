import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { HubConnectionState } from "@microsoft/signalr";

interface State {
  state: HubConnectionState;
}

const initialState: State = {
  state: HubConnectionState.Disconnected
};

export const signalRSlice = createSlice({
  name: "session",
  initialState,
  reducers: {
    setState: (state, action: PayloadAction<HubConnectionState>) => {
      state.state = action.payload;
    }
  }
});

export const { setState } = signalRSlice.actions;

export default signalRSlice.reducer;
