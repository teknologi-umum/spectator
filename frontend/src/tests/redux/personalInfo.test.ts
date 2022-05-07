import { expect, test } from "vitest";
import reducer, { setPersonalInfo } from "@/store/slices/personalInfoSlice";
import { PersonalInfo } from "@/models/PersonalInfo";

const initialState: PersonalInfo = {
  studentNumber: "",
  yearsOfExperience: 0,
  hoursOfPractice: 0,
  familiarLanguages: "",
  walletNumber: "",
  walletType: "gopay"
};

test("should return the initial state", () => {
  expect(reducer(undefined, { type: null })).toEqual(initialState);
});

test("should be able to set personal info data", () => {
  const data: PersonalInfo = {
    studentNumber: "133017",
    yearsOfExperience: 3,
    hoursOfPractice: 2,
    familiarLanguages: "rust, haskell",
    walletNumber: "08472817817",
    walletType: "grabpay"
  };

  expect(reducer(initialState, setPersonalInfo(data))).toEqual(data);
});
