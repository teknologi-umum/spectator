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
import { useEffect, useState } from "react";
import { ReflexContainer, ReflexElement, ReflexSplitter } from "react-reflex";
import "react-reflex/styles.css";

// TODO: ini soal ambil dari json atau sejenisnya, jangan langsung tulis disini
export default function CodingTest() {
  const [time, setTime] = useState(90 * 60); // 90 minutes

  useEffect(() => {
    const timer = setInterval(() => setTime((prev) => prev - 1), 1000);
    return () => clearInterval(timer);
  });

  const connection = useSignalR("fake_hub_url");
  useEventListener("click", mouseClickHandler(connection));
  useEventListener("mousemove", mouseMoveHandler(connection));
  useEventListener("keydown", keystrokeHandler(connection));
  useEventListener("contextmenu", (e) => e.preventDefault());

  const gray = useColorModeValue("gray.100", "gray.800");
  const bg = useColorModeValue("white", "gray.700");
  const fg = useColorModeValue("gray.800", "gray.100");
  const codeBg = useColorModeValue("gray.200", "gray.800");

  return (
    <Box w="full" h="full" bg={gray} gap="3" p="3">
      <Menu bg={bg} fg={fg} time={time} />
      <Box h="calc(100% - 3.5rem)">
        <ReflexContainer orientation="vertical">
          <ReflexElement minSize={400} style={{ overflow: "hidden" }}>
            <Question bg={bg} fg={fg} time={time} codeBg={codeBg} />
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
