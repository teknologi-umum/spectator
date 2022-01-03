import { expect, test } from "vitest";
import reducer, {
  changeFontSize,
  changeCurrentLanguage
} from "@/store/slices/editorSlice";
import type { InitialState } from "@/store/slices/editorSlice/types";

const initialState: InitialState = {
  currentLanguage: "javascript",
  fontSize: 14,
  solutions: [],
  scratchPads: []
};

test("should return the initial state", () => {
  expect(reducer(undefined, { type: null })).toEqual(initialState);
});

test("should be able to set current editor language", () => {
  expect(reducer(initialState, changeCurrentLanguage("c++"))).toEqual({
    ...initialState,
    currentLanguage: "c++"
  });
});

test("should be able to set current editor font size", () => {
  expect(reducer(initialState, changeFontSize(18))).toEqual({
    ...initialState,
    fontSize: 18
  });
});
