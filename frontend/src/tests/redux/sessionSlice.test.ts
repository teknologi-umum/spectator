import { describe, expect } from "vitest";
import reducer, {
  allowVideoPermission,
  markFirstSAMSubmitted,
  markSecondSAMSubmitted,
  markTourCompleted,
  SessionState,
  setAccessToken,
  setVideoDeviceId
} from "@/store/slices/sessionSlice";

const initialState: SessionState = {
  sessionId: null,
  accessToken: null,
  firstSAMSubmitted: false,
  secondSAMSubmitted: false,
  hasPermission: false,
  deviceId: null,
  tourCompleted: {
    personalInfo: false,
    samTest: false,
    codingTest: false
  }
};

describe("Initial State", (it) => {
  it("should return the initial state", () => {
    expect(reducer(undefined, { type: null })).toEqual(initialState);
  });
});

describe("Access Token", (it) => {
  it("should be able to set the access token", () => {
    expect(reducer(initialState, setAccessToken("foo"))).toEqual({
      ...initialState,
      accessToken: "foo"
    });
  });
});

describe("SAM Test", (it) => {
  it("should be able to submit the first sam test", () => {
    expect(reducer(initialState, markFirstSAMSubmitted())).toEqual({
      ...initialState,
      firstSAMSubmitted: true
    });
  });

  it("should be able to submit the second sam test", () => {
    expect(reducer(initialState, markSecondSAMSubmitted())).toEqual({
      ...initialState,
      secondSAMSubmitted: true
    });
  });
});

describe("Tour", (it) => {
  it("should be able to complete the personal info tour", () => {
    expect(reducer(initialState, markTourCompleted("personalInfo"))).toEqual({
      ...initialState,
      tourCompleted: {
        ...initialState.tourCompleted,
        personalInfo: true
      }
    });
  });

  it("should be able to complete the sam test tour", () => {
    expect(reducer(initialState, markTourCompleted("samTest"))).toEqual({
      ...initialState,
      tourCompleted: {
        ...initialState.tourCompleted,
        samTest: true
      }
    });
  });

  it("should be able to complete the coding test tour", () => {
    expect(reducer(initialState, markTourCompleted("codingTest"))).toEqual({
      ...initialState,
      tourCompleted: {
        ...initialState.tourCompleted,
        codingTest: true
      }
    });
  });
});

describe("Video", (it) => {
  it("should be able to complete the personal info tour", () => {
    expect(reducer(initialState, allowVideoPermission())).toEqual({
      ...initialState,
      hasPermission: true
    });
  });

  it("should be able to set the device ID", () => {
    expect(reducer(initialState, setVideoDeviceId("bruh"))).toEqual({
      ...initialState,
      deviceId: "bruh"
    });
  });
});

