import { useEffect, useState } from "react";
import { Button, Flex, Select, Text } from "@chakra-ui/react";
import { TimeIcon } from "@chakra-ui/icons";
import ThemeButton from "../ThemeButton";
import {
  changeFontSize,
  changeCurrentLanguage
} from "@/store/slices/editorSlice";
import { useAppDispatch, useAppSelector } from "@/store";
import { prevQuestion, nextQuestion } from "@/store/slices/questionSlice";

function toReadableTime(seconds: number): string {
  const s = Math.floor(seconds % 60);
  const m = Math.floor((seconds / 60) % 60);
  const h = Math.floor((seconds / (60 * 60)) % 24);
  return [h, m.toString().padStart(2, "0"), s.toString().padStart(2, "0")].join(
    ":"
  );
}

const LANGUAGES = ["javascript", "java", "php", "python", "c", "cpp"];

interface MenuProps {
  bg: string;
  fg: string;
}

export default function Menu({ bg, fg }: MenuProps) {
  const [time, setTime] = useState(90 * 60); // 90 minutes
  const dispatch = useAppDispatch();
  const { fontSize, currentLanguage } = useAppSelector((state) => state.editor);

  useEffect(() => {
    const timer = setInterval(() => setTime((prev) => prev - 1), 1000);
    return () => clearInterval(timer);
  });

  return (
    <Flex display="flex" justifyContent="stretch" gap="3" h="2.5rem" mb="3">
      <Flex
        bg={bg}
        color={fg}
        justifyContent="center"
        alignItems="center"
        h="full"
        px="4"
        gap="2"
        rounded="md"
      >
        <TimeIcon />
        <Text fontWeight="medium" fontSize="lg">
          {toReadableTime(time)}
        </Text>
      </Flex>
      <ThemeButton position="relative" />
      <Flex alignItems="center" gap="3" w="14rem">
        <Select
          bg={bg}
          textTransform="capitalize"
          w="8rem"
          border="none"
          value={currentLanguage}
          onChange={(e) => {
            const language = e.currentTarget.value;
            dispatch(changeCurrentLanguage(language));
          }}
        >
          {LANGUAGES.map((lang, idx) => (
            <option
              key={idx}
              value={lang}
              style={{ textTransform: "capitalize" }}
            >
              {lang === "cpp" ? "C++" : lang}
            </option>
          ))}
        </Select>
        <Select
          bg={bg}
          textTransform="capitalize"
          w="6rem"
          border="none"
          value={fontSize}
          onChange={(e) => {
            const fontSize = parseInt(e.currentTarget.value);
            dispatch(changeFontSize(fontSize));
          }}
        >
          {Array(9)
            .fill(0)
            .map((_, idx: number) => {
              const fontSize = idx + 12;
              return (
                <option
                  key={idx}
                  value={fontSize}
                  style={{ textTransform: "capitalize" }}
                >
                  {fontSize}px
                </option>
              );
            })}
        </Select>
      </Flex>
      <Flex alignItems="center" gap="3" ml="auto">
        <Button
          px="4"
          colorScheme="red"
          h="full"
          onClick={() => {
            // TODO(elianiva): implement proper surrender logic properly
            //                 it's now temporarily used for previous question
            //                 to make testing easier
            dispatch(prevQuestion());
          }}
        >
          Surrender
        </Button>
        <Button
          px="4"
          colorScheme="blue"
          variant="outline"
          h="full"
          onClick={() => {
            // TODO(elianiva): send the code to backend for execution
          }}
        >
          Test
        </Button>
        <Button
          px="4"
          colorScheme="blue"
          h="full"
          onClick={() => {
            // TODO(elianiva): only allow to continue when they have the correct answer
            dispatch(nextQuestion());
          }}
        >
          Submit
        </Button>
      </Flex>
    </Flex>
  );
}
