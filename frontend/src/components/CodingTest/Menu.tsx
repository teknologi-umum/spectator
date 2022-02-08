// TODO(elianiva): this file is kinda big, should probably refactor it later

import React, { useEffect, useState } from "react";
import {
  Box,
  Button,
  Flex,
  Select,
  Text,
  ToastId,
  useToast
} from "@chakra-ui/react";
import ThemeButton from "../ThemeButton";
import {
  setFontSize,
  setLanguage,
  setSnapshot
} from "@/store/slices/editorSlice";
import type { Language } from "@/models/Language";
import { useAppDispatch, useAppSelector } from "@/store";
import { useColorModeValue } from "@/hooks";
import theme from "@/styles/themes";
import { useTranslation } from "react-i18next";
import { jwtDecode } from "@/utils/jwtDecode";
import { CheckmarkIcon, ClockIcon, CrossIcon } from "@/icons";

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
  const toast = useToast();
  const toastBg = useColorModeValue("white", "gray.600", "gray.700");
  const toastFg = useColorModeValue("gray.700", "gray.600", "gray.700");
  const green = useColorModeValue("green.500", "green.400", "green.300");
  const red = useColorModeValue("red.500", "red.400", "red.300");
  const dispatch = useAppDispatch();
  const optionBg = useColorModeValue(
    theme.colors.white,
    theme.colors.gray[700],
    theme.colors.gray[800]
  );
  const {
    currentQuestionNumber,
    fontSize,
    currentLanguage,
    snapshotByQuestionNumber
  } = useAppSelector((state) => state.editor);

  const { t } = useTranslation();

  const { accessToken } = useAppSelector((state) => state.session);
  const decoded = accessToken ? jwtDecode(accessToken) : null;
  const [time, setTime] = useState(
    decoded ? decoded.iat + decoded.exp - Date.now() : 0
  );

  useEffect(() => {
    const timer = setInterval(() => {
      setTime((prev: number) => prev - 1000);
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  const currentSnapshot = snapshotByQuestionNumber[currentQuestionNumber!];
  const isSubmitted =
    currentSnapshot !== undefined ? currentSnapshot.submissionAccepted : false;
  const isRefactored =
    currentSnapshot !== undefined
      ? currentSnapshot.submissionRefactored
      : false;

  function handleSubmit() {
    if (currentQuestionNumber === null) return;

    // TODO(elianiva): submit the actual thing to backend using SignalR
    const isCorrect = Math.random() < 0.5;
    try {
      // TODO(elianiva): submit the actual submission
      // const submissionResult = await sessionSpoke.submitSolution({
      //   // FIXME(elianiva): fix this dumb thing
      //   language: LanguageEnum[currentLanguage.toUpperCase()],
      //   solution: "",
      //   scratchPad: "",
      //   questionNumber: currentQuestionNumber
      // });

      dispatch(
        setSnapshot({
          language: currentLanguage,
          questionNumber: currentQuestionNumber,
          scratchPad: currentSnapshot?.scratchPad || "",
          solutionByLanguage: {
            ...currentSnapshot?.solutionByLanguage,
            [currentLanguage]: ""
          },
          submissionAccepted: isCorrect,
          submissionRefactored: currentSnapshot?.submissionSubmitted || false,
          submissionSubmitted: true,
          // TODO(elianiva): replace this with the actual submission result
          testResults: [
            { testNumber: 1, status: "Passing" },
            {
              testNumber: 2,
              status: "Failing",
              expectedStdout: "1",
              actualStdout: "2"
            },
            { testNumber: 3, status: "CompileError", stderr: "Unexpected '('" },
            {
              testNumber: 4,
              status: "RuntimeError",
              stderr: "Couldn't found 'foo' in current scope"
            },
            {
              testNumber: 5,
              status: "CompileError",
              stderr: "Invalid data type for 'foo'"
            }
          ]
        })
      );

      const id = toast({
        position: "top-right",
        render: () => (
          <Box
            bg={toastBg}
            color={toastFg}
            borderLeft="4px"
            borderColor={isCorrect ? green : red}
            p={4}
            borderRadius="md"
            fontSize="md"
            fontWeight="bold"
            textAlign="left"
            shadow="sm"
            onClick={() => toast.close(id as ToastId)}
            cursor="pointer"
          >
            <Flex align="center" gap="2" color={isCorrect ? green : red}>
              {isCorrect ? (
                <>
                  <CheckmarkIcon width="1.25rem" height="1.25rem" /> Correct answer!
                </>
              ) : (
                <>
                  <CrossIcon width="1.25rem" height="1.25rem" /> Wrong answer!
                </>
              )}
            </Flex>
          </Box>
        )
      });
    } catch (e) {
      // eslint-disable-next-line no-console
      console.error(e);
    }
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
        <ClockIcon />
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
            dispatch(setLanguage(language as Language));
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
              value={currentSnapshot?.language ?? ""}
            >
              {currentSnapshot?.language === "cpp"
                ? "C++"
                : currentSnapshot?.language}
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
            dispatch(setFontSize(fontSize));
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
                  style={{
                    textTransform: "capitalize",
                    backgroundColor: optionBg
                  }}
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
          opacity="75%"
          _hover={{
            opacity: "100%"
          }}
          h="full"
          onClick={() => {
            // TODO(elianiva): implement proper surrender logic properly
            //                 it's now temporarily used for previous question
            //                 to make testing easier
            // dispatch(prevQuestion());
          }}
        >
          {t("translation.translations.ui.surrender")}
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
        {!isRefactored && (
          <Button
            px="4"
            colorScheme="blue"
            h="full"
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
