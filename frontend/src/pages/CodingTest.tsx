import {
  Grid,
  useEventListener,
  useColorModeValue
} from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { useSignalR } from "@/hooks";
import { mouseClickHandler, mouseMoveHandler } from "@/events";
import { Question, Editor, Menu, Scratchpad } from "@/components/CodingTest";

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

  const gray = useColorModeValue("gray.100", "gray.800");
  const border = useColorModeValue("gray.200", "gray.600");
  const bg = useColorModeValue("white", "gray.700");
  const fg = useColorModeValue("gray.800", "gray.100");
  const codeBg = useColorModeValue("gray.200", "gray.700");

  return (
    <Grid
      w="full"
      h="full"
      gridTemplateColumns="1fr 1fr"
      gridTemplateRows="2.5rem 1fr 1fr"
      bg={gray}
      gap="3"
      p="3"
    >
      <Menu bg={bg} fg={fg} time={time} />
      <Question bg={bg} border={border} fg={fg} time={time} codeBg={codeBg} />
      <Editor bg={bg} />
      <Scratchpad bg={bg} />
    </Grid>
  );
}
