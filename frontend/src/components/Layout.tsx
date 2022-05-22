import React, { useEffect, useMemo } from "react";
import { useColorModeValue } from "@/hooks";
import { Box } from "@chakra-ui/react";
import type { FC } from "react";
import ToastOverlay from "@/components/Toast/ToastOverlay";
import { eventSpoke, sessionSpoke } from "@/spoke";
import { useAppDispatch, useAppSelector } from "@/store";
import { Locale as DtoLocale } from "@/stub/enums";
import { setAccessToken } from "@/store/slices/sessionSlice";
import { HubConnectionState, LogLevel } from "@microsoft/signalr";
import { loggerInstance } from "@/spoke/logger";

interface LayoutProps {
  display?: "flex" | "block";
}
const Layout: FC<LayoutProps> = ({ display = "block", children }) => {
  const bg = useColorModeValue("gray.100", "gray.800", "gray.900");

  const dispatch = useAppDispatch();
  const { locale } = useAppSelector((state) => state.locale);
  const { accessToken } = useAppSelector((state) => state.session);
  const { connectionState } = useAppSelector((state) => state.signalR);

  const isTokenEmpty = useMemo(
    () => accessToken === null || accessToken === undefined,
    [accessToken]
  );
  const isHubDisconnected = useMemo(
    () => connectionState === HubConnectionState.Disconnected,
    [connectionState]
  );

  useEffect(() => {
    let dtoLocale: DtoLocale;
    switch (locale) {
    case "EN":
      dtoLocale = DtoLocale.EN;
      break;
    case "ID":
      dtoLocale = DtoLocale.ID;
      break;
    default:
      console.error(`Unknown locale: ${locale}`);
      return;
    }

    if (isTokenEmpty) {
      sessionSpoke
        .startSession({ locale: dtoLocale })
        .then((sessionReply) => {
          sessionSpoke.setAccessToken(sessionReply.accessToken);
          eventSpoke.setAccessToken(sessionReply.accessToken);
          dispatch(setAccessToken(sessionReply.accessToken));
        })
        .catch((err) => {
          if (err instanceof Error) {
            loggerInstance.log(
              LogLevel.Error,
              `Unable to start session: ${err.message}`
            );
          }
        });
    } else if (!isTokenEmpty && isHubDisconnected) {
      sessionSpoke.resumeSession().catch((err) => {
        if (err instanceof Error) {
          loggerInstance.log(
            LogLevel.Error,
            `Unable to resume session: ${err.message}`
          );
        }
      });
    }
  }, []);

  return (
    <>
      <ToastOverlay />
      <Box
        bg={bg}
        alignItems="center"
        justifyContent="center"
        w="full"
        minH="full"
        py="10"
        px="4"
        display={display}
      >
        {children}
      </Box>
    </>
  );
};

export default Layout;
