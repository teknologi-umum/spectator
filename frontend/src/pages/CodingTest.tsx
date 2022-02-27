import React, { useEffect } from "react";
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
  scrollHandler
} from "@/events";
import { useColorModeValue } from "@/hooks";
import { Box, Flex, theme, useEventListener } from "@chakra-ui/react";
import { useAppSelector } from "@/store";
import ToastOverlay from "@/components/ToastOverlay";
import WithTour from "@/hoc/WithTour";
import { codingTestTour } from "@/tours";
import { useTour } from "@reactour/tour";

function CodingTest() {
  const { isCollapsed } = useAppSelector((state) => state.sideBar);
  const { currentQuestionNumber } = useAppSelector(
    (state) => state.editor
  );
  const { tourCompleted } = useAppSelector((state) => state.session);

  useEventListener("mousedown", mouseClickHandler(currentQuestionNumber));
  useEventListener("mousemove", mouseMoveHandler(currentQuestionNumber));
  useEventListener("keydown", keystrokeHandler( currentQuestionNumber));

  useEventListener("scroll", scrollHandler(currentQuestionNumber));

  // disable right click
  // useEventListener("contextmenu", (e) => e.preventDefault());

  const gray = useColorModeValue("gray.100", "gray.800", "gray.900");
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const fgDarker = useColorModeValue("gray.700", "gray.300", "gray.400");

  const { setIsOpen} = useTour();

  useEffect(() => {
    document.title = "Coding Test | Spectator";
    console.log(tourCompleted);
    if (tourCompleted.codingTest) return;
    setIsOpen(true);
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
                <Question
                  bg={bg}
                  fg={fg}
                  fgDarker={fgDarker}
                  onScroll={scrollHandler(connection, currentQuestionNumber)}
                />
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
                    <Editor
                      bg={bg}
                      onScroll={scrollHandler(
                        connection,
                        currentQuestionNumber
                      )}
                    />
                  </ReflexElement>

                  <ReflexSplitter
                    style={{
                      backgroundColor: "transparent",
                      height: theme.space[3],
                      border: "none"
                    }}
                  />

                  <ReflexElement minSize={200} style={{ overflow: "hidden" }}>
                    <ScratchPad
                      bg={bg}
                      onScroll={scrollHandler(
                        connection,
                        currentQuestionNumber
                      )}
                    />
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
