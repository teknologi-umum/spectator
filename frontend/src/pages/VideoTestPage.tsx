import React, { useEffect, useMemo, useRef, useState } from "react";
import {
  Box,
  Button,
  HStack,
  MenuItemOption,
  MenuOptionGroup,
  Stack
} from "@chakra-ui/react";
import { useColorModeValue, useVideoSources } from "@/hooks";
import Layout from "@/components/Layout";
import { SettingsDropdown } from "@/components/Settings";
import { MenuDropdown } from "@/components/TopBar";
import { getUserMedia } from "@/utils/getUserMedia";
import { useNavigate } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "@/store";
import { allowVideoPermission } from "@/store/slices/sessionSlice";

export default function VideoTestPage() {
  // hooks
  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const { hasPermission } = useAppSelector((state) => state.session);

  // styles
  const videoBackground = useColorModeValue("gray.400", "gray.700", "gray.800");

  // local states
  const videoElement = useRef<HTMLVideoElement | null>(null);
  const [videoStream, setVideoStream] = useState<MediaStream | null>(null);
  const [isAllowed, setAllowed] = useState(hasPermission);
  const videoSources = useVideoSources({ isAllowed });
  const activeSourceName = useMemo(() => {
    if (videoStream === null) return "Unknown";
    const sourceName = videoStream.getTracks()?.[0].label ?? "Unknown";
    return sourceName;
  }, [videoStream]);

  async function acquirePermission() {
    // this is using the old way of checking permission since firefox doesn't support permissions API for camera
    try {
      const stream = await getUserMedia();
      setVideoStream(stream);
      setAllowed(true);
    } catch (err: unknown) {
      setAllowed(false);
      alert("Please allow camera permission and refresh the page.");
    }
  }

  async function changeVideoSource(deviceId: string) {
    const newStream = await getUserMedia(deviceId);
    setVideoStream(newStream);
  }

  function startCodingTest() {
    dispatch(allowVideoPermission());
    navigate("/coding-test");
  }

  useEffect(() => {
    // only try starting the preview when the user has given us the permission
    if (videoElement.current === null || !isAllowed) return;
    try {
      videoElement.current.srcObject = videoStream;
    } catch (err) {
      if (err instanceof DOMException) {
        if (err.message === "Permission denied") {
          alert("Please allow camera access");
        }
      }
      // eslint-disable-next-line no-console
      console.error(err);
    }
  }, [isAllowed]);

  return (
    <Layout display="flex">
      <SettingsDropdown />
      <Stack spacing={4} align="center">
        <Box
          position="relative"
          w={640}
          h={360}
          rounded="md"
          overflow="hidden"
          bg={videoBackground}
          sx={{
            "&::before": !isAllowed
              ? {
                content: "\"Your video will appear here\"",
                position: "absolute",
                left: "50%",
                top: "50%",
                transform: "translateX(-50%) translateY(-50%)",
                fontSize: "1.5rem"
              }
              : {}
          }}
        >
          <video ref={videoElement} autoPlay width={640} height={360} />
        </Box>
        {isAllowed ? (
          <HStack spacing={4}>
            <MenuDropdown
              dropdownWidth="10rem"
              title={activeSourceName ?? "Select video source"}
            >
              <MenuOptionGroup
                type="radio"
                onChange={(deviceId) => changeVideoSource(deviceId as string)}
              >
                {videoSources.map((source) => (
                  <MenuItemOption
                    key={source.deviceId}
                    textTransform="capitalize"
                    value={source.deviceId ?? ""}
                  >
                    <span>{source.label}</span>
                  </MenuItemOption>
                ))}
              </MenuOptionGroup>
            </MenuDropdown>
            <Button colorScheme="blue" onClick={startCodingTest}>
              Start the Coding Test
            </Button>
          </HStack>
        ) : (
          <Button colorScheme="blue" onClick={acquirePermission}>
            Allow Permission
          </Button>
        )}
      </Stack>
    </Layout>
  );
}
