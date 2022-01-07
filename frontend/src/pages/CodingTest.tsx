import {
  Editor,
  Menu,
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
import { useColorModeValue, useSignalR } from "@/hooks";
import { withProtected } from "@/hoc";
import {
  Box,
  Flex,
  theme,
  useEventListener
} from "@chakra-ui/react";
import React, { useEffect } from "react";
import { ReflexContainer, ReflexElement, ReflexSplitter } from "react-reflex";
import "react-reflex/styles.css";
import { useAppSelector } from "@/store";
import type { InitialState as QuestionState } from "@/store/slices/questionSlice/types";

function CodingTest() {
  const { currentQuestion } = useAppSelector<QuestionState>(
    (state) => state.question
  );

  const connection = useSignalR("fake_hub_url");

  useEventListener("mousedown", mouseClickHandler(connection, currentQuestion));
  useEventListener("mousemove", mouseMoveHandler(connection, currentQuestion));
  useEventListener("keydown", keystrokeHandler(connection, currentQuestion));

  useEventListener("scroll", scrollHandler(connection, currentQuestion));

  // disable right click
  // useEventListener("contextmenu", (e) => e.preventDefault());

  const gray = useColorModeValue("gray.100", "gray.800", "gray.900");
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const fgDarker = useColorModeValue("gray.700", "gray.400", "gray.400");

  useEffect(() => {
    document.title = "Coding Test | Spectator";
  }, []);

  return (
    <Flex w="full" h="full">
      <SideBar bg={bg} fg={fg} />
      <Box bg={gray} gap="3" p="3">
        <Menu bg={bg} fgDarker={fgDarker} />
        <Box h="calc(100% - 3.5rem)">
          <ReflexContainer orientation="vertical">
            <ReflexElement minSize={400} style={{ overflow: "hidden" }}>
              <Question
                bg={bg}
                fg={fg}
                fgDarker={fgDarker}
                onScroll={scrollHandler(connection, currentQuestion)}
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
                    onScroll={scrollHandler(connection, currentQuestion)}
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
                    onScroll={scrollHandler(connection, currentQuestion)}
                  />
                </ReflexElement>
              </ReflexContainer>
            </ReflexElement>
          </ReflexContainer>
        </Box>
      </Box>
    </Flex>
  );
}

export default withProtected(CodingTest);
