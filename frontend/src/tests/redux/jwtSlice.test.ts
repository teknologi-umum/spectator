import { describe, expect } from "vitest";
import reducer, { setAccessToken } from "@/store/slices/sessionSlice";

const FAKE_JWT =
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdHVkZW50SWQiOiIxMjM0NTY3ODkwIiwiaWF0IjoxNjQxMDMxMTYwODE3LCJleHAiOjYwMDAwfQ==.ySxmfXC7SSP4NR7Go2qitririWvL-vWMLZUDjY0w6U8";

const initialState = {
  accessToken: null,
  sessionId: null,
  firstSAMSubmitted: false,
  secondSAMSubmitted: false,
  tourCompleted: {
    personalInfo: false,
    samTest: false,
    codingTest: false
  }
};

describe("jwt", (it) => {
  it("should return the initial state", () => {
    expect(reducer(undefined, { type: null })).toEqual(initialState);
  });

  it("should be able to set access token", () => {
    expect(reducer(initialState, setAccessToken(FAKE_JWT))).toEqual({
      ...initialState,
      accessToken: FAKE_JWT,
      firstSAMSubmitted: false,
      secondSAMSubmitted: false
    });
  });
});
