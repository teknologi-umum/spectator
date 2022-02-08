import { expect, test } from "vitest";
import reducer, {
  setFontSize
} from "@/store/slices/editorSlice";
import type { EditorState } from "@/models/EditorState";

const initialState: EditorState = {
  deadlineUtc: null,
  questions: null,
  currentQuestionNumber: 0,
  currentLanguage: "javascript",
  fontSize: 14,
  snapshotByQuestionNumber: {}
};

test("should return the initial state", () => {
  expect(reducer(undefined, { type: null })).toEqual(initialState);
});

//test("should be able to set current editor language", () => {
//  expect(reducer(initialState, changeCurrentLanguage("c++"))).toEqual({
//    ...initialState,
//    currentLanguage: "c++"
//  });
//});

test("should be able to set current editor font size", () => {
  expect(reducer(initialState, setFontSize(18))).toEqual({
    ...initialState,
    fontSize: 18
  });
});
