import React, { useMemo } from "react";
import { useAppSelector } from "@/store";
import { Navigate, Outlet, useLocation } from "react-router-dom";

export const SecretRoute = () => {
  const { sessionId } = useAppSelector((state) => state.session);
  const location = useLocation();

  const validPath = useMemo(() => {
    return sessionId === null ? "/secret/login" : "/secret/download";
  }, [sessionId]);

  if (location.pathname !== validPath) {
    return <Navigate to={validPath} />;
  }

  return <Outlet />;
};
