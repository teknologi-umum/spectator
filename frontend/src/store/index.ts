import { configureStore } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import { personalInfoReducer, questionReducer } from "./slices";

const store = configureStore({
  devTools: true,
  reducer: {
    personalInfo: personalInfoReducer,
    question: questionReducer
  }
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

// nanti ini yang bakalan dipake di semua app, bukan `useDispatch` dan `useSelector`
export const useAppDispatch = (): AppDispatch => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;

export default store;
