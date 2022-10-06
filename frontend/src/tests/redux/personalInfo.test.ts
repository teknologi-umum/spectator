import { describe, expect } from "vitest";
import reducer, { setPersonalInfo } from "@/store/slices/personalInfoSlice";
import { PersonalInfo } from "@/models/PersonalInfo";

const initialState: PersonalInfo = {
  email: "",
  age: 0,
  gender: "M",
  nationality: "indonesia",
  studentNumber: "",
  yearsOfExperience: 0,
  hoursOfPractice: 0,
  familiarLanguages: "",
  walletNumber: "",
  walletType: "grabpay"
};

describe("Personal Info", (it) => {
  it("should return the initial state", () => {
    expect(reducer(undefined, { type: null })).toEqual(initialState);
  });

  it("should be able to set personal info data", () => {
    const data: PersonalInfo = {
      email: "something@gmail.com",
      age: 18,
      gender: "M",
      nationality: "indonesia",
      studentNumber: "133017",
      yearsOfExperience: 3,
      hoursOfPractice: 2,
      familiarLanguages: "rust, haskell",
      walletNumber: "08472817817",
      walletType: "grabpay"
    };

    expect(reducer(initialState, setPersonalInfo(data))).toEqual(data);
  });
});
