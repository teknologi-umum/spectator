import { Navigate } from "react-router-dom";
import { useAppSelector } from "@/store";

export default function FinalRoute({ children }: { children: JSX.Element }) {
  const { jwt, hasFinished } = useAppSelector((state) => state.jwt);

  if (!jwt) {
    return <Navigate to="/" />;
  }

  console.log(hasFinished);

  if (jwt && !hasFinished) {
    return <Navigate to="/coding-test" />;
  }

  return children;
}
