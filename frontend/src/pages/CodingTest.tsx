import { Editor, Menu, Question, Scratchpad } from "@/components/CodingTest";
import {
  keystrokeHandler,
  mouseClickHandler,
  mouseMoveHandler
} from "@/events";
import { useSignalR } from "@/hooks";
import {
  Box,
  theme,
  useColorModeValue,
  useEventListener
} from "@chakra-ui/react";
import { ReflexContainer, ReflexElement, ReflexSplitter } from "react-reflex";
import "react-reflex/styles.css";

// TODO: ini soal ambil dari json atau sejenisnya, jangan langsung tulis disini
export default function CodingTest() {
  const connection = useSignalR("fake_hub_url");
  useEventListener("click", mouseClickHandler(connection));
  useEventListener("mousemove", mouseMoveHandler(connection));
  useEventListener("keydown", keystrokeHandler(connection));
  useEventListener("contextmenu", (e) => e.preventDefault());

  const gray = useColorModeValue("gray.100", "gray.900");
  const bg = useColorModeValue("white", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100");

  return (
    <Box w="full" h="full" bg={gray} gap="3" p="3">
      <Menu bg={bg} fg={fg} />
      <Box h="calc(100% - 3.5rem)">
        <ReflexContainer orientation="vertical">
          <ReflexElement minSize={400} style={{ overflow: "hidden" }}>
            <Question bg={bg} fg={fg} />
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
                <Scratchpad bg={bg} />
              </ReflexElement>
            </ReflexContainer>
          </ReflexElement>
        </ReflexContainer>
      </Box>
    </Box>
  );
}
