import { expect, test } from "vitest";
import reducer, {
  nextQuestion,
  prevQuestion
} from "@/store/slices/questionSlice";

test("should return the initial state", () => {
  expect(reducer(undefined, { type: null })).toEqual({
    currentQuestion: 0,
    submissions: []
  });
});

test("should be able to go to the next question", () => {
  expect(
    reducer({ currentQuestion: 0, submissions: [] }, nextQuestion())
  ).toEqual({
    currentQuestion: 1,
    submissions: []
  });
});

test("should not be able to go past 6th question", () => {
  expect(
    reducer({ currentQuestion: 5, submissions: [] }, nextQuestion())
  ).toEqual({
    currentQuestion: 5,
    submissions: []
  });
});

test("should be able to go to the previous question", () => {
  expect(
    reducer({ currentQuestion: 1, submissions: [] }, prevQuestion())
  ).toEqual({
    currentQuestion: 0,
    submissions: []
  });
});

test("should not be able to go past 0th question", () => {
  expect(
    reducer({ currentQuestion: 0, submissions: [] }, prevQuestion())
  ).toEqual({
    currentQuestion: 0,
    submissions: []
  });
});
