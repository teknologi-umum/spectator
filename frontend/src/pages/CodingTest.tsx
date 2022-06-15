import React, { useEffect, useMemo } from "react";
import { ReflexContainer, ReflexElement, ReflexSplitter } from "react-reflex";
import "react-reflex/styles.css";
import {
  Box,
  Button,
  Flex,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Text,
  theme,
  ToastId,
  useDisclosure,
  useEventListener,
  useToast
} from "@chakra-ui/react";
import { useTour } from "@reactour/tour";
import WithTour from "@/hoc/WithTour";
import { codingTestTour } from "@/tours";
import { Editor, Problem, ScratchPad, SideBar } from "@/components/CodingTest";
import ToastBase from "@/components/Toast/ToastBase";
import ToastOverlay from "@/components/Toast/ToastOverlay";
import { TopBar } from "@/components/TopBar";
import {
  keystrokeHandler,
  mouseUpHandler,
  mouseDownHandler,
  mouseMoveHandler,
  mouseScrollHandler,
  windowResizeHandler
} from "@/events";
import { useColorModeValue, useVideoRecorder, RecordingStatus } from "@/hooks";
import { useAppSelector } from "@/store";
import { sessionSpoke } from "@/spoke";
import { CrossIcon, VideoIcon } from "@/icons";
import { loggerInstance } from "@/spoke/logger";
import { LogLevel } from "@microsoft/signalr";
import { useTranslation } from "react-i18next";

function CodingTest() {
  const { t } = useTranslation("translation", {
    keyPrefix: "translations"
  });
  const { isCollapsed } = useAppSelector((state) => state.codingTest);
  const { currentQuestionNumber } = useAppSelector((state) => state.editor);
  const { tourCompleted } = useAppSelector((state) => state.session);
  const { accessToken } = useAppSelector((state) => state.session);
  const isTokenEmpty = useMemo(
    () => accessToken === null || accessToken === undefined,
    [accessToken]
  );

  const gray = useColorModeValue("gray.100", "gray.800", "gray.900");
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const fgDarker = useColorModeValue("gray.700", "gray.300", "gray.400");
  const green = useColorModeValue("green.500", "green.400", "green.300");
  const red = useColorModeValue("red.500", "red.400", "red.300");
  const toastBg = useColorModeValue("white", "gray.600", "gray.700");
  const toastFg = useColorModeValue("gray.700", "gray.600", "gray.700");

  const toast = useToast();
  const { setIsOpen } = useTour();

  const { isOpen, onOpen, onClose } = useDisclosure();

  const recordingStatus = useVideoRecorder(accessToken);

  useEventListener(
    "mousedown",
    mouseDownHandler(currentQuestionNumber, accessToken)
  );
  useEventListener(
    "mouseup",
    mouseUpHandler(currentQuestionNumber, accessToken)
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

  // TODO(elianiva): uncomment this on production, disable right click
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
          loggerInstance.log(LogLevel.Error, err);

          const id = toast({
            position: "top-right",
            render: () => (
              <ToastBase
                bg={toastBg}
                fg={toastFg}
                borderColor={red}
                onClick={() => toast.close(id as ToastId)}
              >
                <Flex align="center" gap="2" color={red}>
                  <CrossIcon width="1.25rem" height="1.25rem" /> Unable to
                  resume the exam.
                </Flex>
              </ToastBase>
            )
          });
        });
    }
  }, []);

  useEffect(() => {
    if (recordingStatus === RecordingStatus.UNKNOWN) return;

    const accentColor =
      recordingStatus === RecordingStatus.STARTED ? green : red;
    const id = toast({
      position: "top-right",
      render: () => (
        <ToastBase
          bg={toastBg}
          fg={toastFg}
          borderColor={accentColor}
          onClick={() => toast.close(id as ToastId)}
        >
          <Flex align="center" gap="2" color={accentColor}>
            <VideoIcon width="1.25rem" height="1.25rem" /> Recording{" "}
            {recordingStatus}!
          </Flex>
        </ToastBase>
      )
    });
  }, [recordingStatus]);

  async function forfeitExam() {
    if (accessToken === null) return;
    await sessionSpoke.forfeitExam({ accessToken });
  }

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
          <TopBar bg={bg} fg={fgDarker} forfeitExam={onOpen} />
          <Box h="calc(100% - 3.5rem)">
            <ReflexContainer orientation="vertical">
              <ReflexElement minSize={400} style={{ overflow: "hidden" }}>
                <Problem bg={bg} fg={fg} fgDarker={fgDarker} />
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
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent bg={bg} color={fg}>
          <ModalHeader fontSize="2xl">
            {t("surrender.title")}
          </ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Text fontSize="lg" lineHeight="7">
              {t("surrender.body")}
            </Text>
          </ModalBody>

          <ModalFooter>
            <Button
              colorScheme="blue"
              variant="outline"
              mr={3}
              onClick={onClose}
            >
              {t("ui.cancel")}
            </Button>
            <Button colorScheme="red" onClick={forfeitExam} data-tour="step-2">
              {t("ui.surrender")}
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
}

export default WithTour(CodingTest, codingTestTour);
