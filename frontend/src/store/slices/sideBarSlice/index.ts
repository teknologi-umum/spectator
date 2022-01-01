import { createSlice } from "@reduxjs/toolkit";
import type { InitialState } from "./types";

const initialState: InitialState = {
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
