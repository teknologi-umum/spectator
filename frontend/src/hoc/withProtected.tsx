import React, { useEffect } from "react";
import { useNavigate, Navigate } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "@/store";
import { finishSession } from "@/store/slices/jwtSlice";
import type { InitialState as JwtState } from "@/store/slices/jwtSlice/types";
import type { ComponentType } from "react";

const withProtected = <P extends Record<string, unknown>>(
  WrappedComponent: ComponentType<P>
) => {
  // Naming the wrapped component is important because
  // react will complaint if you only return anonymous function
  const ComponentWithProtected = (props: P) => {
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

    if (jwt === "" || hasExpired) {
      return <Navigate to="/" />;
    }

    if (jwt && hasFinished) {
      return <Navigate to="/fun-fact" />;
    }

    return <WrappedComponent {...props} />;
  };

  return ComponentWithProtected;
};

export default withProtected;
