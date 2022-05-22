import reducer, {
  ExamResultState,
  setExamResult
} from "@/store/slices/examResultSlice";
import { ExamResult } from "@/stub/session";
import { describe, expect } from "vitest";

const initialState: ExamResultState = {
  examResult: null
};

const examResult: ExamResult = {
  answeredQuestionNumbers: [1, 2],
  duration: BigInt(123),
  funFact: {
    deletionRate: 0.2,
    submissionAttempts: BigInt(123),
    wordsPerMinute: BigInt(123)
  }
};

describe("Exam Result", (it) => {
  it("should be able to return the initial state", () => {
    expect(reducer(undefined, { type: null })).toEqual(initialState);
  });

  it("should be able to set the exam result data", () => {
    expect(reducer(initialState, setExamResult(examResult))).toEqual({
      ...initialState,
      examResult: examResult
    });
  });
});
