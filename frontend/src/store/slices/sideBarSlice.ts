import { createSlice } from "@reduxjs/toolkit";

interface State {
  isCollapsed: boolean;
}

const initialState: State = {
  isCollapsed: true
};

export const sideBarSlice = createSlice({
  name: "sideBar",
  initialState,
  reducers: {
    toggleSideBar: (state) => {
      state.isCollapsed = !state.isCollapsed;
    }
  }
});

export const { toggleSideBar } = sideBarSlice.actions;

export default sideBarSlice.reducer;
