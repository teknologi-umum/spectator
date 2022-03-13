import React, { useEffect, useMemo } from "react";
import { useColorModeValue } from "@/hooks";
import { Box } from "@chakra-ui/react";
import type { FC } from "react";
import ToastOverlay from "@/components/ToastOverlay";
import { eventSpoke, sessionSpoke } from "@/spoke";
import { useAppDispatch, useAppSelector } from "@/store";
import { Locale as DtoLocale } from "@/stub/enums";
import { setAccessToken } from "../store/slices/sessionSlice";
import { HubConnectionState } from "@microsoft/signalr";

const Layout: FC = ({ children }) => {
  const bg = useColorModeValue("gray.100", "gray.800", "gray.900");
  const dispatch = useAppDispatch();
  const { locale } = useAppSelector((state) => state.locale);
  const { accessToken } = useAppSelector((state) => state.session);
  const { connectionState } = useAppSelector((state) => state.signalR);
  const isTokenEmpty = useMemo(() => { 
    return (accessToken === null || accessToken === undefined);
  }, [accessToken]);
  const isHubDisconnected = useMemo(() => {
    return connectionState === HubConnectionState.Disconnected;
  }, [connectionState]);

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
          console.error(`Unable to start session. ${err}`);
        });
    } else if (!isTokenEmpty && isHubDisconnected) {
      // TODO(elianiva): figure out how to reconnect before the exam started
    }
  }, []);

  return (
    <>
      <ToastOverlay />
      <Box bg={bg} alignItems="center" w="full" minH="full" py="10" px="4">
        {children}
      </Box>
    </>
  );
};

export default Layout;
