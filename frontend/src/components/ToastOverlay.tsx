import React, { useEffect, useMemo, useState } from "react";
import { Box, Flex, ToastId, useToast } from "@chakra-ui/react";
import { useAppSelector } from "@/store";
import { HubConnectionState } from "@microsoft/signalr";
import { CheckmarkIcon, CrossIcon, InfoIcon } from "@/icons";
import { useColorModeValue } from "@/hooks";

export default function ToastOverlay() {
  const toast = useToast();
  const { connectionState: state } = useAppSelector((state) => state.signalR);
  const [borderColor, setBorderColor] = useState("");
  const toastBg = useColorModeValue("white", "gray.600", "gray.700");
  const toastFg = useColorModeValue("gray.700", "gray.600", "gray.700");
  const green = useColorModeValue("green.500", "green.400", "green.300");
  const blue = useColorModeValue("blue.500", "blue.400", "blue.300");
  const red = useColorModeValue("red.500", "red.400", "red.300");

  const toastContent = useMemo(() => {
    switch (state) {
    case HubConnectionState.Connected:
      setBorderColor(green);
      return (
        <Flex align="center" gap="2" color={green}>
          <CheckmarkIcon width="1.25rem" height="1.25rem" /> Session {state}!
        </Flex>
      );
    case HubConnectionState.Disconnected:
      setBorderColor(red);
      return (
        <Flex align="center" gap="2" color={red}>
          <CrossIcon width="1.25rem" height="1.25rem" /> Session {state}!
        </Flex>
      );
    default:
      setBorderColor(blue);
      return (
        <Flex align="center" gap="2" color={blue}>
          <InfoIcon width="1.25rem" height="1.25rem" /> Session {state}!
        </Flex>
      );
    }
  }, [state]);

  useEffect(() => {
    const id = toast({
      position: "top-right",
      render: () => (
        <Box
          bg={toastBg}
          color={toastFg}
          borderLeft="4px"
          borderColor={borderColor}
          p={4}
          borderRadius="md"
          fontSize="md"
          fontWeight="bold"
          textAlign="left"
          shadow="sm"
          onClick={() => toast.close(id as ToastId)}
          cursor="pointer"
        >
          {toastContent}
        </Box>
      )
    });
  }, [state]);

  return <></>;
}
