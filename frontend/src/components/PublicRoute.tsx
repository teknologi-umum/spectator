import { Navigate } from "react-router-dom";
import { useAppSelector } from "@/store";
import type { InitialState as JwtState } from "@/store/slices/jwtSlice/types";

export default function PublicRoute({ children }: { children: JSX.Element }) {
  const { jwt, jwtPayload, hasFinished } = useAppSelector<JwtState>(
    (state) => state.jwt
  );
  const hasExpired = jwtPayload
    ? (jwtPayload.exp as number) - Date.now() < 0
    : true;

  if (jwt && !hasExpired) {
    return <Navigate to="/coding-test" />;
  }

  if (jwt && hasFinished) {
    return <Navigate to="/fun-fact" />;
  }

  return children;
}
