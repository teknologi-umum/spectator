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
  setLockedToCurrentQuestion,
  setSnapshot,
  setSolution
} from "@/store/slices/editorSlice";
import type { EditorSnapshot } from "@/models/EditorSnapshot";
import { LANGUAGES, Language } from "@/models/Language";
import { useAppDispatch, useAppSelector } from "@/store";
import { useColorModeValue } from "@/hooks";
import { ClockIcon } from "@/icons";
import { CodingResultToast } from "@/components/Toast";
import { MenuDropdown, ThemeButton, LocaleButton } from "@/components/TopBar";
import { sessionSpoke } from "@/spoke";
import { Solution } from "@/models/Solution";
import { loggerInstance } from "@/spoke/logger";
import { LogLevel } from "@microsoft/signalr";
import { setExamResult } from "@/store/slices/examResultSlice";
import { setQuestionTabIndex } from "@/store/slices/codingTestSlice";
import { SubmissionResult } from "@/models/SubmissionResult";
import { FinishedModal } from "@/components/CodingTest";

function toReadableTime(ms: number): string {
  const seconds = ms / 1000;
  const s = Math.floor(seconds % 60);
  const m = Math.floor((seconds / 60) % 60);
  const h = Math.floor((seconds / (60 * 60)) % 24);
  return [h, m.toString().padStart(2, "0"), s.toString().padStart(2, "0")].join(
    ":"
  );
}

interface TopBarProps {
  bg: string;
  fg: string;
  forfeitExam: () => void;
}

export default function TopBar({ bg, fg, forfeitExam }: TopBarProps) {
  const navigate = useNavigate();
  const toast = useToast();
  const dispatch = useAppDispatch();
  const { t } = useTranslation("translation", {
    keyPrefix: "translations.ui"
  });
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

  const [isModalOpen, setModalOpen] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const [testing, setTesting] = useState(false);

  useEffect(() => {
    const timer = setInterval(() => {
      setTime((prev: number) => prev - 1000);
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  useEffect(() => {
    const endTimer = setTimeout(async () => {
      if (accessToken === null) return;

      setModalOpen(true);
      const result = await sessionSpoke.passDeadline({ accessToken });
      dispatch(setExamResult(result));
      navigate("/sam-test");
    }, time);

    return () => clearTimeout(endTimer);
  });

  const currentSnapshot: EditorSnapshot | undefined =
    snapshotByQuestionNumber[currentQuestionNumber];
  const isSubmitted =
    currentSnapshot !== undefined ? currentSnapshot.submissionSubmitted : false;

  useEffect(() => {
    if (accessToken === null) return;
    (async () => {
      const allSnapshots = Object.values(snapshotByQuestionNumber);
      if (allSnapshots.length < 6) {
        // if they haven't tried all of the questions
        // just don't bother checking if they have been accepted or not
        return;
      }

      const isSessionFinished = allSnapshots.reduce((prev, curr) => {
        // only count correct answers
        return prev && curr.submissionSubmitted && curr.submissionAccepted;
      }, true);

      if (isSessionFinished) {
        setModalOpen(true);
        // automatically end the exam when all of their submissions have been accepted
        // and this is the last submission that was accepted
        try {
          const result = await sessionSpoke.endExam({ accessToken });
          dispatch(setExamResult(result));
          navigate("/sam-test");
        } catch (err) {
          if (err instanceof Error) {
            loggerInstance.log(LogLevel.Error, err.message);
          }
        }
      }
    })();
  }, [snapshotByQuestionNumber]);

  async function submitSolution(submissionType: "submit" | "test") {
    if (
      currentQuestionNumber === null ||
      accessToken === null ||
      currentSnapshot === undefined
    ) {
      return;
    }

    const solution = new Solution(
      currentLanguage,
      currentSnapshot.solutionByLanguage[currentLanguage]
    );

    try {
      const submissionData = {
        accessToken,
        language: solution.language,
        directives: solution.getDirective(),
        solution: solution.content,
        scratchPad: currentSnapshot.scratchPad,
        questionNumber: currentQuestionNumber
      };

      let submissionResult: SubmissionResult;
      dispatch(setLockedToCurrentQuestion(true));
      if (submissionType === "submit") {
        setSubmitting(true);
        submissionResult = await sessionSpoke.submitSolution(submissionData);
      } else if (submissionType === "test") {
        setTesting(true);
        submissionResult = await sessionSpoke.testSolution(submissionData);
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
          // mark as refactored only if it has been submitted before
          submissionRefactored: currentSnapshot.submissionSubmitted,
          submissionSubmitted: currentSnapshot.submissionSubmitted
            ? true // don't change the value if it's already set to true
            : submissionType === "submit",
          testResults: submissionResult.testResults,
          samTestResult: currentSnapshot.samTestResult
        })
      );

      dispatch(setLockedToCurrentQuestion(false));
      setSubmitting(false);
      setTesting(false);

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

      if (submissionType === "submit" && currentQuestionNumber !== 0) {
        navigate("/sam-test");
      }
    } catch (err) {
      if (err instanceof Error) {
        loggerInstance.log(LogLevel.Error, err.message);
      }
    }
  }

  function resetBoilerplate() {
    // we set it to empty because the editor will fill it with the boilerplate when it's empty
    dispatch(setSolution(""));
  }

  return (
    <>
      <FinishedModal isOpen={isModalOpen} />
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
            colorScheme="teal"
            opacity="75%"
            _hover={{
              opacity: "100%"
            }}
            h="full"
            onClick={resetBoilerplate}
            data-tour="topbar-step-6"
          >
            {t("reset")}
          </Button>
          <Button
            px="4"
            colorScheme="red"
            opacity="75%"
            _hover={{
              opacity: "100%"
            }}
            h="full"
            onClick={forfeitExam}
            data-tour="topbar-step-6"
          >
            {t("surrender")}
          </Button>
          <Button
            px="4"
            colorScheme="blue"
            variant="outline"
            h="full"
            isLoading={testing}
            onClick={() => submitSolution("test")}
            data-tour="topbar-step-7"
          >
            Test
          </Button>
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
        </Flex>
      </Flex>
    </>
  );
}
