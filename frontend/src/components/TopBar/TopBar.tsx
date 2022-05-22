import React, { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import {
  Button,
  Flex,
  MenuItemOption,
  MenuOptionGroup,
  Text,
  useToast
} from "@chakra-ui/react";
import {
  setFontSize,
  setLanguage,
  setSnapshot
} from "@/store/slices/editorSlice";
import type { EditorSnapshot } from "@/models/EditorSnapshot";
import { LANGUAGES, Language } from "@/models/Language";
import { useAppDispatch, useAppSelector } from "@/store";
import { useColorModeValue } from "@/hooks";
import { ClockIcon } from "@/icons";
import CodingResultToast from "@/components/Toast/CodingResultToast";
import { MenuDropdown, ThemeButton, LocaleButton } from "@/components/TopBar";
import { sessionSpoke } from "@/spoke";
import { Solution } from "@/models/Solution";
import { loggerInstance } from "@/spoke/logger";
import { LogLevel } from "@microsoft/signalr";
import { setExamResult } from "@/store/slices/examResultSlice";
import { setQuestionTabIndex } from "@/store/slices/codingTestSlice";
import { SubmissionResult } from "@/stub/session";

function toReadableTime(ms: number): string {
  const seconds = ms / 1000;
  const s = Math.floor(seconds % 60);
  const m = Math.floor((seconds / 60) % 60);
  const h = Math.floor((seconds / (60 * 60)) % 24);
  return [h, m.toString().padStart(2, "0"), s.toString().padStart(2, "0")].join(
    ":"
  );
}

interface MenuProps {
  bg: string;
  fg: string;
}

export default function TopBar({ bg, fg }: MenuProps) {
  const navigate = useNavigate();
  const toast = useToast();
  const dispatch = useAppDispatch();
  const { t } = useTranslation();
  const {
    currentQuestionNumber,
    fontSize,
    currentLanguage,
    snapshotByQuestionNumber
  } = useAppSelector((state) => state.editor);
  const { accessToken } = useAppSelector((state) => state.session);
  const { deadlineUtc } = useAppSelector((state) => state.editor);
  const [time, setTime] = useState(deadlineUtc ? deadlineUtc - Date.now() : 0);

  const toastBg = useColorModeValue("white", "gray.600", "gray.700");
  const toastFg = useColorModeValue("gray.700", "gray.600", "gray.700");
  const green = useColorModeValue("green.500", "green.400", "green.300");
  const red = useColorModeValue("red.500", "red.400", "red.300");

  useEffect(() => {
    const timer = setInterval(() => {
      setTime((prev: number) => prev - 1000);
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  useEffect(() => {
    const endTimer = setInterval(async () => {
      if (accessToken === null) return;

      const result = await sessionSpoke.passDeadline({ accessToken });
      dispatch(setExamResult(result));
      navigate("/sam-test");
    }, time);

    return () => clearInterval(endTimer);
  });

  const currentSnapshot: EditorSnapshot | undefined =
    snapshotByQuestionNumber[currentQuestionNumber];
  const isSubmitted =
    currentSnapshot !== undefined ? currentSnapshot.submissionAccepted : false;
  const isRefactored =
    currentSnapshot !== undefined
      ? currentSnapshot.submissionRefactored
      : false;

  const [submitting, setSubmitting] = useState(false);

  async function submitSolution(submissionType: "submit" | "test") {
    if (
      currentQuestionNumber === null ||
      accessToken === null ||
      currentSnapshot === undefined
    ) {
      return;
    }

    setSubmitting(true);

    const solution = new Solution(
      currentLanguage,
      currentSnapshot.solutionByLanguage[currentLanguage]
    );

    try {
      let submissionResult: SubmissionResult;
      if (submissionType === "submit") {
        submissionResult = await sessionSpoke.submitSolution({
          accessToken,
          language: solution.language,
          directives: solution.getDirective(),
          solution: solution.content,
          scratchPad: currentSnapshot.scratchPad,
          questionNumber: currentQuestionNumber
        });
      } else if (submissionType === "test") {
        submissionResult = await sessionSpoke.testSolution({
          accessToken,
          language: solution.language,
          directives: solution.getDirective(),
          solution: solution.content,
          scratchPad: currentSnapshot.scratchPad,
          questionNumber: currentQuestionNumber
        });
      } else {
        throw new Error("Invalid type");
      }

      dispatch(
        setSnapshot({
          language: currentLanguage,
          questionNumber: currentQuestionNumber,
          scratchPad: currentSnapshot.scratchPad,
          solutionByLanguage: currentSnapshot.solutionByLanguage,
          submissionAccepted: submissionResult.accepted,
          submissionRefactored: currentSnapshot.submissionSubmitted,
          submissionSubmitted: submissionType === "submit",
          testResults: submissionResult.testResults
        })
      );

      setSubmitting(false);

      // move to the result tab
      dispatch(setQuestionTabIndex("result"));

      const id = toast({
        position: "top-right",
        render: () => (
          <CodingResultToast
            bg={toastBg}
            fg={toastFg}
            green={green}
            red={red}
            isCorrect={submissionResult.accepted}
            onClick={() => toast.close(id!)}
          />
        )
      });

      const allSnapshots = Object.values(snapshotByQuestionNumber);
      if (allSnapshots.length < 6) {
        // if they haven't submitted all of their submissions
        // just don't bother checking if they have been accepted or not
        return;
      }

      // this will only be true when every submissions have been accepted
      const isLastSolution = allSnapshots.reduce((acc, curr) => {
        return curr.submissionAccepted && acc;
      }, false);

      if (isLastSolution) {
        // automatically end the exam when all of their submissions have been accepted
        // and this is the last submission that was accepted
        try {
          const result = await sessionSpoke.endExam({ accessToken });
          dispatch(setExamResult(result));
          navigate("/fun-fact");
        } catch (err) {
          if (err instanceof Error) {
            loggerInstance.log(LogLevel.Error, err.message);
          }
        }
      }
    } catch (err) {
      if (err instanceof Error) {
        loggerInstance.log(LogLevel.Error, err.message);
      }
    }
  }

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
        data-tour="topbar-step-1"
      >
        <ClockIcon />
        <Text fontWeight="medium" fontSize="lg">
          {toReadableTime(time)}
        </Text>
      </Flex>
      <Flex alignItems="center" gap="3" w="32rem">
        <ThemeButton bg={bg} fg={fg} data-tour="topbar-step-2" />
        <MenuDropdown
          dropdownWidth="10rem"
          bg={bg}
          fg={fg}
          title={currentLanguage}
          data-tour="topbar-step-3"
          data-testid="editor-language-select"
        >
          <MenuOptionGroup
            type="radio"
            value={isSubmitted ? currentSnapshot.language : currentLanguage}
            onChange={(value) => dispatch(setLanguage(value as Language))}
          >
            {LANGUAGES.map((lang, idx) => (
              <MenuItemOption
                textTransform="capitalize"
                key={idx}
                value={lang}
                isDisabled={
                  isSubmitted ? lang !== currentSnapshot?.language : false
                }
              >
                <span>{lang === "cpp" ? "C++" : lang}</span>
              </MenuItemOption>
            ))}
          </MenuOptionGroup>
        </MenuDropdown>
        <MenuDropdown
          dropdownWidth="10rem"
          bg={bg}
          fg={fg}
          title={fontSize + "px"}
          data-tour="topbar-step-4"
          data-testid="editor-fontsize-select"
        >
          <MenuOptionGroup
            type="radio"
            value={String(fontSize)}
            onChange={(value) => {
              dispatch(setFontSize(parseInt(value as string)));
            }}
          >
            {Array(9)
              .fill(0)
              .map((_, idx: number) => {
                const fontSize = idx + 12;
                return (
                  <MenuItemOption key={idx} value={String(fontSize)}>
                    <span>{fontSize}px</span>
                  </MenuItemOption>
                );
              })}
          </MenuOptionGroup>
        </MenuDropdown>
        <LocaleButton bg={bg} fg={fg} data-tour="topbar-step-5" />
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
          onClick={async () => {
            if (accessToken === null) return;
            await sessionSpoke.forfeitExam({ accessToken });
          }}
          data-tour="topbar-step-6"
        >
          {t("translation.translations.ui.surrender")}
        </Button>
        <Button
          px="4"
          colorScheme="blue"
          variant="outline"
          h="full"
          isLoading={submitting}
          onClick={() => submitSolution("test")}
          data-tour="topbar-step-7"
        >
          Test
        </Button>
        {!isRefactored && (
          <Button
            px="4"
            colorScheme="blue"
            h="full"
            isLoading={submitting}
            onClick={() => submitSolution("submit")}
            data-tour="topbar-step-8"
          >
            {isSubmitted ? "Refactor" : "Submit"}
          </Button>
        )}
      </Flex>
    </Flex>
  );
}
