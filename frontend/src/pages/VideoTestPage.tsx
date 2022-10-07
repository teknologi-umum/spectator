import React, { useEffect, useMemo, useRef, useState } from "react";
import {
  Box,
  Button,
  HStack,
  MenuItemOption,
  MenuOptionGroup,
  Stack,
  type ToastId,
  useToast,
  Flex,
  Text
} from "@chakra-ui/react";
import { useColorModeValue, useVideoSources } from "@/hooks";
import Layout from "@/components/Layout";
import { SettingsDropdown } from "@/components/Settings";
import { MenuDropdown } from "@/components/TopBar";
import { getUserMedia } from "@/utils/getUserMedia";
import { useNavigate } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "@/store";
import {
  allowVideoPermission,
  setVideoDeviceId
} from "@/store/slices/sessionSlice";
import { ToastBase } from "@/components/Toast";
import { CrossIcon } from "@/icons";
import { useTranslation } from "react-i18next";

export default function VideoTestPage() {
  // hooks
  const toast = useToast();
  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const { hasPermission, deviceId } = useAppSelector((state) => state.session);
  const { t } = useTranslation("translation", {
    keyPrefix: "translations.video_test"
  });

  // styles
  const videoBackground = useColorModeValue("gray.400", "gray.700", "gray.800");
  const toastBg = useColorModeValue("white", "gray.600", "gray.700");
  const red = useColorModeValue("red.500", "red.400", "red.300");

  // local states
  const videoElement = useRef<HTMLVideoElement | null>(null);
  const [isAllowed, setAllowed] = useState(hasPermission);
  const videoSources = useVideoSources({ isAllowed });
  const [videoStream, setVideoStream] = useState<MediaStream | null>(null);
  // if it's allowed previously, set the initial video stream
  useEffect(() => {
    // get initial videostream when it's already allowed
    getUserMedia(deviceId).then((stream) => setVideoStream(stream));
  }, [isAllowed]);

  const activeSourceName = useMemo(() => {
    if (videoStream === null || videoStream === undefined) return "Unknown";
    const sourceName = videoStream.getTracks()?.[0].label;
    return sourceName;
  }, [videoStream]);

  function showAlert() {
    const id = toast({
      position: "top-right",
      render: () => (
        <ToastBase
          bg={toastBg}
          fg={red}
          borderColor={red}
          onClick={() => toast.close(id as ToastId)}
        >
          <Flex align="center" gap="3" color={red}>
            <CrossIcon width="1.25rem" height="1.25rem" /> {t("notification")}
          </Flex>
        </ToastBase>
      )
    });
  }

  async function acquirePermission() {
    // this is using the old way of checking permission since firefox doesn't support permissions API for camera
    try {
      const stream = await getUserMedia();
      setVideoStream(stream);
      // eslint-disable-next-line no-console
      console.debug(
        "Acquired video stream, please open the arrow on the right.",
        stream
      );
      setAllowed(true);
    } catch (err: unknown) {
      setAllowed(false);
      showAlert();
    }
  }

  async function changeVideoSource(deviceId: string) {
    if (deviceId === "unknown") {
      dispatch(setVideoDeviceId(null));
      setVideoStream(null);
    }

    const newStream = await getUserMedia(deviceId);
    setVideoStream(newStream);
    dispatch(setVideoDeviceId(deviceId));
    // eslint-disable-next-line no-console
    console.debug(
      "This message means that video device ID and video stream has been successfully set."
    );
  }

  function startCodingTest() {
    dispatch(allowVideoPermission());
    // eslint-disable-next-line no-console
    console.debug(
      "This message means that we have successfully set the allow video permission flag."
    );
    navigate("/coding-test");
  }

  useEffect(() => {
    // only try starting the preview when the user has given us the permission
    if (videoElement.current === null || !isAllowed) return;
    try {
      videoElement.current.srcObject = videoStream;
    } catch (err) {
      if (
        err instanceof DOMException &&
        err.message.toLowerCase() === "permission denied"
      ) {
        showAlert();
      }
      // eslint-disable-next-line no-console
      console.error(err);
    }
  }, [isAllowed, videoStream]);

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
                content: `"${t("no_video_placeholder")}"`,
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
          <Stack justify="center" align="center">
            <HStack spacing={4}>
              <MenuDropdown
                dropdownWidth="10rem"
                title={activeSourceName ?? t("dropdown_title")}
              >
                <MenuOptionGroup
                  type="radio"
                  onChange={(deviceId: string | string[]) => {
                    // no need to handle `string[]` since deviceId will never be an array of string
                    changeVideoSource(deviceId as string);
                  }}
                  value={deviceId ?? "unknown"}
                >
                  <MenuItemOption textTransform="capitalize" value="unknown">
                    <span>Unknown</span>
                  </MenuItemOption>
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
                {t("start_coding_test")}
              </Button>
            </HStack>
            <Text>{t("instruction")}</Text>
          </Stack>
        ) : (
          <Button colorScheme="blue" onClick={acquirePermission}>
            {t("allow_permission")}
          </Button>
        )}
      </Stack>
    </Layout>
  );
}
