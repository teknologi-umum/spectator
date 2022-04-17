import React, { useEffect, useMemo, useRef } from "react";
import { ReflexContainer, ReflexElement, ReflexSplitter } from "react-reflex";
import "react-reflex/styles.css";
import {
  Editor,
  TopBar,
  Question,
  ScratchPad,
  SideBar
} from "@/components/CodingTest";
import {
  keystrokeHandler,
  mouseClickHandler,
  mouseMoveHandler,
  mouseScrollHandler,
  windowResizeHandler
} from "@/events";
import { useColorModeValue } from "@/hooks";
import { Box, Flex, theme, useEventListener } from "@chakra-ui/react";
import { useAppSelector } from "@/store";
import ToastOverlay from "@/components/ToastOverlay";
import WithTour from "@/hoc/WithTour";
import { codingTestTour } from "@/tours";
import { useTour } from "@reactour/tour";
import { sessionSpoke } from "@/spoke";

function CodingTest() {
  const { isCollapsed } = useAppSelector((state) => state.sideBar);
  const { currentQuestionNumber } = useAppSelector((state) => state.editor);
  const { tourCompleted } = useAppSelector((state) => state.session);
  const { accessToken } = useAppSelector((state) => state.session);
  const isTokenEmpty = useMemo(() => {
    return accessToken === null || accessToken === undefined;
  }, [accessToken]);

  const gray = useColorModeValue("gray.100", "gray.800", "gray.900");
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const fgDarker = useColorModeValue("gray.700", "gray.300", "gray.400");
  const videoStream = useRef<MediaStream | null>();
  const mediaRecorder = useRef<MediaRecorder | null>();

  const { setIsOpen } = useTour();

  useEventListener(
    "mousedown",
    mouseClickHandler(currentQuestionNumber, accessToken)
  );
  useEventListener(
    "mousemove",
    mouseMoveHandler(currentQuestionNumber, accessToken)
  );
  useEventListener(
    "keydown",
    keystrokeHandler(currentQuestionNumber, accessToken)
  );
  useEventListener(
    "wheel",
    mouseScrollHandler(currentQuestionNumber, accessToken)
  );
  useEventListener(
    "resize",
    windowResizeHandler(currentQuestionNumber, accessToken)
  );

  // disable right click
  // useEventListener("contextmenu", (e) => e.preventDefault());

  useEffect(() => {
    document.title = "Coding Test | Spectator";
    if (tourCompleted.codingTest) return;
    setIsOpen(true);
  }, []);

  useEffect(() => {
    if (!isTokenEmpty) {
      sessionSpoke
        .resumeExam({
          // this is actually safe since accessToken is never going to be null/undefined
          // thanks to `!isTokenEmpty` check above.
          // it's just that tsserver can't pick it up, so yeah, `as string` it is.
          accessToken: accessToken as string
        })
        .catch((err) => {
          console.error(`Unable to resume the exam session. ${err}`);
        });
    }
  }, []);

  useEffect(() => {
    if (accessToken === null) {
      return;
    }

    // We want to generate the filename here because we need it as a marker
    // for the start of the recording.
    // `MediaRecorder` splice the video stream into chunks, only the first chunk is a valid video
    // so we need to mark this first chunk.
    // A user could theoretically refresh the page and the recording would get restarted.
    // If we don't have this marker, we wouldn't know which one is the first chunk for each recording session.
    const filename = `${Date.now()}.webm`;

    (async () => {
      // TODO(elianiva): move this section into its own video stream client class
      const stream = await navigator.mediaDevices.getUserMedia({
        audio: false,
        video: {
          width: 640,
          height: 320
        }
      });
      videoStream.current = stream;
      mediaRecorder.current = new MediaRecorder(stream, {
        mimeType: "video/webm;codecs=vp9",
        videoBitsPerSecond: 200_000 // 0.2Mbits / sec
      });
      mediaRecorder.current.start(1000); // send blob every second
      mediaRecorder.current.onstart = () => {
        console.log("Start recording...");
      };
      mediaRecorder.current.ondataavailable = async (e: BlobEvent) => {
        const formData = new FormData();
        formData.append("accessToken", accessToken);
        formData.append("file", e.data, filename);
        await fetch(import.meta.env.VITE_VIDEO_STREAM_URL, {
          method: "POST",
          headers: {
            "Content-Length": e.data.size.toString()
          },
          body: formData
        });
      };
    })();

    () => {
      if (mediaRecorder.current) {
        mediaRecorder.current.stop();
        mediaRecorder.current.onstop = () => {
          console.log("Stop recording...");
        };
        videoStream.current = null;
        mediaRecorder.current = null;
      }
    };
  }, []);

  return (
    <>
      <ToastOverlay />
      <Flex w="full" h="full">
        <SideBar bg={bg} fg={fg} />
        <Box
          bg={gray}
          gap="3"
          p="3"
          w={`calc(100% - ${isCollapsed ? "65px" : "200px"})`}
          transition="width 300ms ease"
        >
          <TopBar bg={bg} fg={fgDarker} />
          <Box h="calc(100% - 3.5rem)">
            <ReflexContainer orientation="vertical">
              <ReflexElement minSize={400} style={{ overflow: "hidden" }}>
                <Question bg={bg} fg={fg} fgDarker={fgDarker} />
              </ReflexElement>

              <ReflexSplitter
                style={{
                  backgroundColor: "transparent",
                  width: theme.space[3],
                  border: "none"
                }}
              />

              <ReflexElement minSize={400} style={{ overflow: "hidden" }}>
                <ReflexContainer orientation="horizontal">
                  <ReflexElement minSize={200} style={{ overflow: "hidden" }}>
                    <Editor bg={bg} />
                  </ReflexElement>

                  <ReflexSplitter
                    style={{
                      backgroundColor: "transparent",
                      height: theme.space[3],
                      border: "none"
                    }}
                  />

                  <ReflexElement minSize={200} style={{ overflow: "hidden" }}>
                    <ScratchPad bg={bg} />
                  </ReflexElement>
                </ReflexContainer>
              </ReflexElement>
            </ReflexContainer>
          </Box>
        </Box>
      </Flex>
    </>
  );
}

export default WithTour(CodingTest, codingTestTour);
