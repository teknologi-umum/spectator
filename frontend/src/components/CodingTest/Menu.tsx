import React, { useEffect, useState } from "react";
import { Button, Flex, Select, Text } from "@chakra-ui/react";
import { TimeIcon } from "@chakra-ui/icons";
import ThemeButton from "../ThemeButton";
import {
  changeFontSize,
  changeCurrentLanguage
} from "@/store/slices/editorSlice";
import { useAppDispatch, useAppSelector } from "@/store";
import { useColorModeValue } from "@/hooks";
import theme from "@/styles/themes";
import {
  prevQuestion,
  nextQuestion,
  setSubmission
} from "@/store/slices/questionSlice";
import { mutate } from "@/utils/fakeSubmissionCallback";

function toReadableTime(ms: number): string {
  const seconds = ms / 1000;
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
  fgDarker: string;
}

export default function Menu({ bg, fgDarker }: MenuProps) {
  const dispatch = useAppDispatch();
  const optionBg = useColorModeValue(theme.colors.white, theme.colors.gray[700], theme.colors.gray[800]);
  const { currentQuestion, submissions } = useAppSelector(
    (state) => state.question
  );
  const { fontSize, currentLanguage, solutions } = useAppSelector(
    (state) => state.editor
  );

  const {
    jwtPayload: { exp, iat }
  } = useAppSelector((state) => state.jwt);
  const [time, setTime] = useState(iat + exp - Date.now());

  useEffect(() => {
    const timer = setInterval(() => {
      setTime((prev) => prev - 1000);
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  const recordedSubmission = submissions.find(
    (submission) => submission.questionNo === currentQuestion
  );
  const isSubmitted =
    recordedSubmission !== undefined ? recordedSubmission.isSubmitted : false;
  const isRefactored =
    recordedSubmission !== undefined ? recordedSubmission.isRefactored : false;

  console.log(isSubmitted);

  function handleSubmit() {
    const currentSolution = solutions.filter(
      (solution) =>
        solution.questionNo === currentQuestion &&
        solution.language === currentLanguage
    )[0];

    mutate(currentSolution, {
      onSuccess: (res) => {
        dispatch(
          setSubmission({
            ...res.data,
            isSubmitted: true,
            isRefactored: res.data.submissionType === "refactor"
          })
        );
      }
    });
  }

  return (
    <Flex display="flex" justifyContent="stretch" gap="3" h="2.5rem" mb="3">
      <Flex
        bg={bg}
        color={fgDarker}
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
          color={fgDarker}
          bg={bg}
          textTransform="capitalize"
          w="8rem"
          border="none"
          value={currentLanguage}
          onChange={(e) => {
            const language = e.currentTarget.value;
            dispatch(changeCurrentLanguage(language));
          }}
          data-testid="editor-language-select"
        >
          {!isSubmitted ? (
            <>
              {LANGUAGES.map((lang, idx) => (
                <option
                  key={idx}
                  value={lang}
                  style={{ textTransform: "capitalize" }}
                >
                  {lang === "cpp" ? "C++" : lang}
                </option>
              ))}
            </>
          ) : (
            <option
              style={{ textTransform: "capitalize", backgroundColor: optionBg }}
              value={recordedSubmission?.language ?? ""}
            >
              {recordedSubmission?.language === "cpp"
                ? "C++"
                : recordedSubmission?.language}
            </option>
          )}
        </Select>
        <Select
          color={fgDarker}
          bg={bg}
          textTransform="capitalize"
          w="6rem"
          border="none"
          value={fontSize}
          onChange={(e) => {
            const fontSize = parseInt(e.currentTarget.value);
            dispatch(changeFontSize(fontSize));
          }}
          data-testid="editor-fontsize-select"
        >
          {Array(9)
            .fill(0)
            .map((_, idx: number) => {
              const fontSize = idx + 12;
              return (
                <option
                  key={idx}
                  value={fontSize}
                  style={{ textTransform: "capitalize", backgroundColor: optionBg }}
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
          background="red.500"
          opacity="60%"
          _hover={{
            opacity: "100%"
          }}
          color="white"
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
          _hover={{
            bg: "blue.600",
            borderColor: "white",
            color: "white"
          }}
          onClick={() => {
            // TODO(elianiva): send the code to backend for execution
          }}
        >
          Test
        </Button>
        {!isRefactored && (
          <Button
            px="4"
            background="blue.500"
            color= "white"
            h="full"
            _hover={{
              bg: "gray.800",
              borderColor: "white",
            }}
            onClick={() => {
              // TODO(elianiva): only allow to continue when they have the correct answer
              // dispatch(nextQuestion());

              handleSubmit();
            }}
          >
            {isSubmitted ? "Refactor" : "Submit"}
          </Button>
        )}
      </Flex>
    </Flex>
  );
}
