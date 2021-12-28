import { Navigate } from "react-router-dom";
import { useAppSelector } from "@/store";
import type { InitialState as JwtState } from "@/store/slices/jwtSlice/types";

export default function FinalRoute({ children }: { children: JSX.Element }) {
  const { jwt, hasFinished } = useAppSelector<JwtState>((state) => state.jwt);

  if (!jwt) {
    return <Navigate to="/" />;
  }

  if (jwt && !hasFinished) {
    return <Navigate to="/coding-test" />;
  }

  return children;
}
