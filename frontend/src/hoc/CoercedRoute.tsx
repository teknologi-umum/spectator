import React from "react";
import { Outlet, Navigate, useLocation } from "react-router-dom";
import { useAppSelector } from "@/store";
import { useMemo } from "react";

export default function CoercedRoute() {
  const {
    accessToken,
    firstSAMSubmitted,
    secondSAMSubmitted,
    hasPermission,
    deviceId
  } = useAppSelector((state) => state.session);
  const { studentNumber } = useAppSelector((state) => state.personalInfo);
  const { deadlineUtc, questions } = useAppSelector((state) => state.editor);
  const { examResult } = useAppSelector((state) => state.examResult);
  const location = useLocation();

  const validPath = useMemo(() => {
    // new user
    if (
      accessToken === null ||
      studentNumber === null ||
      studentNumber === ""
    ) {
      return "/";
    }

    // filled the first personal info but haven't done the coding test
    if (
      firstSAMSubmitted === false ||
      deadlineUtc === null ||
      questions === null
    ) {
      return "/sam-test";
    }

    // checking video capture device before continuing with the test
    if (!hasPermission || deviceId === null || deviceId === "") {
      return "/video-test";
    }

    // filled out the prerequisites and is currently doing the test
    if (examResult === null) return "/coding-test";

    // has finished the coding test but haven't done the last SAM test
    if (!secondSAMSubmitted) return "/sam-test";

    // has done everything
    return "/fun-fact";
  }, [
    accessToken,
    studentNumber,
    firstSAMSubmitted,
    deadlineUtc,
    questions,
    examResult,
    secondSAMSubmitted,
    hasPermission
  ]);

  if (location.pathname !== validPath) {
    return <Navigate to={validPath} />;
  }

  return <Outlet />;
}
