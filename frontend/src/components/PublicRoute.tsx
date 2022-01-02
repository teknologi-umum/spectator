import { Navigate } from "react-router-dom";
import { useAppSelector } from "@/store";
import type { InitialState as JwtState } from "@/store/slices/jwtSlice/types";

export default function PublicRoute({ children }: { children: JSX.Element }) {
  const { jwt, jwtPayload, hasFinished } = useAppSelector<JwtState>(
    (state) => state.jwt
  );
  const timeLeft = jwtPayload.exp + jwtPayload.iat - Date.now();
  const hasExpired = timeLeft < 0;

  if (jwt !== "" && !hasExpired) {
    return <Navigate to="/coding-test" />;
  }

  if (jwt !== "" && hasFinished) {
    return <Navigate to="/fun-fact" />;
  }

  return children;
}
