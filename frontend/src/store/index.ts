import { configureStore, combineReducers } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import {
  persistStore,
  persistReducer,
  FLUSH,
  REHYDRATE,
  PAUSE,
  PERSIST,
  PURGE,
  REGISTER
} from "redux-persist";
import storage from "redux-persist/lib/storage";
import {
  personalInfoReducer,
  editorReducer,
  questionReducer,
  jwtReducer,
  themeReducer,
  sideBarReducer,
  loginReducer
} from "./slices";

// see: https://vitejs.dev/guide/env-and-mode.html#modes
const isDev = import.meta.env.MODE === "development";

const persistConfig = {
  key: "root",
  version: 1,
  whitelist: ["editor", "question", "jwt"],
  storage
};

const rootReducer = combineReducers({
  personalInfo: personalInfoReducer,
  editor: editorReducer,
  question: questionReducer,
  jwt: jwtReducer,
  app: themeReducer,
  sideBar: sideBarReducer,
  login: loginReducer
});

const store = configureStore({
  devTools: isDev,
  reducer: persistReducer(persistConfig, rootReducer),
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER]
      }
    })
});

const persistor = persistStore(store);

export type RootState = ReturnType<typeof rootReducer>;
export type AppDispatch = typeof store.dispatch;

// nanti ini yang bakalan dipake di semua app, bukan `useDispatch` dan `useSelector`
export const useAppDispatch = (): AppDispatch => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;

export { store, persistor };
