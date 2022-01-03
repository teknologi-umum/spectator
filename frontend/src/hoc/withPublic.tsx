import React from "react";
import { Navigate } from "react-router-dom";
import { useAppSelector } from "@/store";
import type { InitialState as JwtState } from "@/store/slices/jwtSlice/types";
import type { ComponentType } from "react";

const withPublic = <P extends Record<string, unknown>>(
  WrappedComponent: ComponentType<P>
) => {
  // Naming the wrapped component is important because
  // react will complaint if you only return anonymous function
  const ComponentWithPublic = (props: P) => {
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

    return <WrappedComponent {...props} />;
  };

  return ComponentWithPublic;
};

export default withPublic;
