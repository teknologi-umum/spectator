import { Navigate } from "react-router-dom";
import { useAppSelector } from "@/store";

export default function PrivateRoute({ children }: { children: JSX.Element }) {
  const { jwt, jwtPayload, hasFinished } = useAppSelector((state) => state.jwt);
  const hasExpired = jwtPayload ? jwtPayload.exp - Date.now() < 0 : true;

  if (!jwt || hasExpired) {
    return <Navigate to="/" />;
  }

  if (jwt && hasFinished) {
    return <Navigate to="/fun-fact" />;
  }

  return children;
}
