import { expect, test } from "vitest";
import reducer, {
  removeSessionId,
  setAccessToken,
  setSessionId
} from "@/store/slices/sessionSlice";

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

test("should return the initial state", () => {
  expect(reducer(undefined, { type: null })).toEqual(initialState);
});

test("should be able to set access token", () => {
  expect(reducer(initialState, setAccessToken(FAKE_JWT))).toEqual({
    ...initialState,
    accessToken: FAKE_JWT
  });
});

test("should be able to set session ID", () => {
  expect(reducer(initialState, setSessionId("fake-session-uuid"))).toEqual({
    ...initialState,
    sessionId: "fake-session-uuid"
  });
});

test("should be able to remove session ID", () => {
  expect(
    reducer(
      { ...initialState, sessionId: "fake-session-id" },
      removeSessionId()
    )
  ).toEqual({
    ...initialState,
    sessionId: null
  });
});
