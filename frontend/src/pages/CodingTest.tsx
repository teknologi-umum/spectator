import { Editor, Menu, Question, Scratchpad } from "@/components/CodingTest";
import {
  keystrokeHandler,
  mouseClickHandler,
  mouseMoveHandler,
  scrollHandler
} from "@/events";
import { withProtected } from "@/hoc";
import { useSignalR } from "@/hooks";
import {
  Box,
  theme,
  useColorModeValue,
  useEventListener
} from "@chakra-ui/react";
import { useEffect } from "react";
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
  useEventListener("contextmenu", (e) => e.preventDefault());

  const gray = useColorModeValue("gray.100", "gray.800");
  const bg = useColorModeValue("white", "gray.700");
  const fg = useColorModeValue("gray.800", "gray.100");

  useEffect(() => {
    document.title = "Coding Test | Spectator";
  }, []);

  return (
    <Box w="full" h="full" bg={gray} gap="3" p="3">
      <Menu bg={bg} fg={fg} />
      <Box h="calc(100% - 3.5rem)">
        <ReflexContainer orientation="vertical">
          <ReflexElement minSize={400} style={{ overflow: "hidden" }}>
            <Question
              bg={bg}
              fg={fg}
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
                <Scratchpad
                  bg={bg}
                  onScroll={scrollHandler(connection, currentQuestion)}
                />
              </ReflexElement>
            </ReflexContainer>
          </ReflexElement>
        </ReflexContainer>
      </Box>
    </Box>
  );
}

export default withProtected(CodingTest);
