import { Navigate } from "react-router-dom";
import { useAppSelector } from "@/store";
import type { InitialState as JwtState } from "@/store/slices/jwtSlice/types";
import type { ComponentType } from "react";

const withFinal = <P extends Record<string, unknown>>(
  WrappedComponent: ComponentType<P>
) => {
  // Naming the wrapped component is important because
  // react will complaint if you only return anonymous function
  const ComponentWithFinal = (props: P) => {
    const { jwt, hasFinished } = useAppSelector<JwtState>((state) => state.jwt);

    if (jwt === "") {
      return <Navigate to="/" />;
    }

    if (jwt !== "" && !hasFinished) {
      return <Navigate to="/coding-test" />;
    }

    return <WrappedComponent {...props} />;
  };

  return ComponentWithFinal;
};

export default withFinal;
