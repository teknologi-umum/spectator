import React from "react";
import { Outlet, Navigate, useLocation } from "react-router-dom";
import { useAppSelector } from "@/store";
import { useMemo } from "react";

export const CoercedRoute = () => {
  const { accessToken, firstSAMSubmitted, secondSAMSubmitted } = useAppSelector((state) => state.session);
  const { studentNumber } = useAppSelector((state) => state.personalInfo);
  const { deadlineUtc, questions } = useAppSelector((state) => state.editor);
  const { examResult } = useAppSelector((state) => state.examResult);
  const location = useLocation();

  const validPath = useMemo(
    () => accessToken === null || studentNumber === null || studentNumber === ""
      ? "/"
      : firstSAMSubmitted !== true || deadlineUtc === null || questions === null
        ? "/sam-test"
        : examResult === null
          ? "/coding-test"
          : secondSAMSubmitted === null
            ? "/sam-test"
            : "/fun-fact",
    [accessToken, studentNumber, firstSAMSubmitted, deadlineUtc, questions, examResult, secondSAMSubmitted]
  );

  if (location.pathname !== validPath) {
    return <Navigate to={validPath} />;
  }

  return <Outlet />;
};
