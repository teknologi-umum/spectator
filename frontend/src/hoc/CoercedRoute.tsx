import { PathRouteProps, Navigate, Route, IndexRouteProps } from "react-router-dom";
import { useAppSelector } from "@/store";
import { useMemo } from "react";

export const CoercedRoute = (props: PathRouteProps | IndexRouteProps) => {
  const { accessToken, firstSAMSubmitted, secondSAMSubmitted } = useAppSelector((state) => state.session);
  const { personalInfo } = useAppSelector((state) => state.personalInfo);
  const { deadlineUtc, questions } = useAppSelector((state) => state.editor);
  const { examResult } = useAppSelector((state) => state.examResult);

  const validPath = useMemo(
    () => accessToken === null || personalInfo === null
      ? "/"
      : firstSAMSubmitted !== true || deadlineUtc === null || questions === null
        ? "/sam-test"
        : examResult === null
          ? "/coding-test"
          : secondSAMSubmitted === null
            ? "/sam-test"
            : "/fun-fact",
    [accessToken, personalInfo, firstSAMSubmitted, deadlineUtc, questions, examResult, secondSAMSubmitted]
  );

  if (props.index && validPath !== "/") {
    return <Route {...props} element={<Navigate to="/" />} />;
  }

  if ("path" in props && props.path !== validPath) {
    return <Route {...props} element={<Navigate to={validPath} />} />;
  }

  return <Route {...props} />;
};