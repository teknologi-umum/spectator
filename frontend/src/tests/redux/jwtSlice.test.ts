import { expect, test } from "vitest";
import reducer, { setJwt, finishSession } from "@/store/slices/jwtSlice";

const FAKE_JWT =
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdHVkZW50SWQiOiIxMjM0NTY3ODkwIiwiaWF0IjoxNjQxMDMxMTYwODE3LCJleHAiOjYwMDAwfQ==.ySxmfXC7SSP4NR7Go2qitririWvL-vWMLZUDjY0w6U8";

const initialState = {
  jwt: "",
  jwtPayload: {
    studentId: "",
    iat: 0,
    exp: 0
  },
  hasFinished: false
};

test("should return the initial state", () => {
  expect(reducer(undefined, { type: null })).toEqual(initialState);
});

test("should be able to set jwt string and its parsed payload", () => {
  expect(reducer(initialState, setJwt(FAKE_JWT))).toEqual({
    jwt: FAKE_JWT,
    jwtPayload: {
      studentId: "1234567890",
      iat: 1641031160817,
      exp: 60000
    },
    hasFinished: false
  });
});

test("should be able to finish session", () => {
  expect(reducer(initialState, finishSession())).toEqual({
    ...initialState,
    hasFinished: true
  });
});
