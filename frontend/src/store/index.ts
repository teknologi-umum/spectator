import { configureStore } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import { personalInfoReducer, editorReducer, questionReducer } from "./slices";

// see: https://vitejs.dev/guide/env-and-mode.html#modes
const isDev = import.meta.env.MODE === "development";

const store = configureStore({
  devTools: isDev,
  reducer: {
    personalInfo: personalInfoReducer,
    editor: editorReducer,
    question: questionReducer
  }
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

// nanti ini yang bakalan dipake di semua app, bukan `useDispatch` dan `useSelector`
export const useAppDispatch = (): AppDispatch => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;

export default store;
