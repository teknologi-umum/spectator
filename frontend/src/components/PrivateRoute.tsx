import { useEffect } from "react";
import { useNavigate, Navigate } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "@/store";
import { finishSession } from "@/store/slices/jwtSlice";
import type { InitialState as JwtState } from "@/store/slices/jwtSlice/types";

export default function PrivateRoute({ children }: { children: JSX.Element }) {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const { jwt, jwtPayload, hasFinished } = useAppSelector<JwtState>(
    (state) => state.jwt
  );

  const timeLeft = jwtPayload.exp + jwtPayload.iat - Date.now();
  const hasExpired = timeLeft < 0;

  useEffect(() => {
    setTimeout(() => {
      dispatch(finishSession());
      navigate("/fun-fact");
    }, timeLeft);
  }, []);

  if (!jwt || hasExpired) {
    return <Navigate to="/" />;
  }

  if (jwt && hasFinished) {
    return <Navigate to="/fun-fact" />;
  }

  return children;
}
