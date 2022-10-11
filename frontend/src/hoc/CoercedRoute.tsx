import React from "react";
import { Outlet, Navigate, useLocation } from "react-router-dom";
import { useAppSelector } from "@/store";
import { useMemo } from "react";

export default function CoercedRoute() {
  const { accessToken, firstSAMSubmitted, hasPermission, deviceId } =
    useAppSelector((state) => state.session);
  const { studentNumber } = useAppSelector((state) => state.personalInfo);
  const {
    deadlineUtc,
    questions,
    currentQuestionNumber,
    snapshotByQuestionNumber
  } = useAppSelector((state) => state.editor);
  const { examResult } = useAppSelector((state) => state.examResult);
  const location = useLocation();
  const isCurrentSubmissionSubmitted =
    snapshotByQuestionNumber[currentQuestionNumber]?.submissionSubmitted;
  const currentQuestionHasNoSAMResult =
    snapshotByQuestionNumber[currentQuestionNumber]?.samTestResult === null;

  const validPath = useMemo(() => {
    // new user
    if (
      accessToken === null ||
      studentNumber === null ||
      studentNumber === ""
    ) {
      return "/";
    }

    if (
      // haven't done the first SAM test
      !firstSAMSubmitted ||
      // haven't done the second SAM test
      deadlineUtc === null ||
      // has no questions
      questions === null ||
      // haven't done the question SAM test
      (isCurrentSubmissionSubmitted && currentQuestionHasNoSAMResult)
    ) {
      return "/sam-test";
    }

    // checking video capture device before continuing with the test
    if (!hasPermission || deviceId === null || deviceId === "") {
      return "/video-test";
    }

    // filled out the prerequisites and is currently doing the test
    if (examResult === null) return "/coding-test";

    // has done everything
    return "/fun-fact";
  }, [
    accessToken,
    studentNumber,
    firstSAMSubmitted,
    deadlineUtc,
    questions,
    isCurrentSubmissionSubmitted,
    currentQuestionHasNoSAMResult,
    hasPermission,
    deviceId,
    examResult
  ]);

  if (location.pathname !== validPath) {
    return <Navigate to={validPath} />;
  }

  return <Outlet />;
}
