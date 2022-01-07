import { expect, test } from "vitest";
import reducer, { recordPersonalInfo } from "@/store/slices/personalInfoSlice";

const initialState = {
  studentId: "",
  programmingExp: 0,
  programmingExercise: 0,
  programmingLanguage: ""
};

test("should return the initial state", () => {
  expect(reducer(undefined, { type: null })).toEqual(initialState);
});

test("should be able to set personal info data", () => {
  const data = {
    studentId: "133017",
    programmingExp: 3,
    programmingExercise: 2,
    programmingLanguage: "rust, haskell"
  };

  expect(reducer(initialState, recordPersonalInfo(data))).toEqual(data);
});
