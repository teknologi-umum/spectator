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
    if (
      accessToken === null ||
      studentNumber === null ||
      studentNumber === ""
    ) {
      return "/";
    }

    if (
      firstSAMSubmitted === false ||
      deadlineUtc === null ||
      questions === null
    ) {
      return "/sam-test";
    }

    if (!hasPermission || deviceId === null || deviceId === "") {
      return "/video-test";
    }

    if (examResult === null) return "/coding-test";

    if (!secondSAMSubmitted) return "/sam-test";

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
