import React, { useEffect, useState } from "react";
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
import type { Language } from "@/models/Language";
import { LANGUAGES } from "@/models/Language";
import { useAppDispatch, useAppSelector } from "@/store";
import { useTranslation } from "react-i18next";
import { useColorModeValue } from "@/hooks";
import { ClockIcon } from "@/icons";
import FeedbackToast from "@/components/FeedbackToast";
import {
  MenuDropdown,
  ThemeButton,
  LocaleButton
} from "@/components/CodingTest";
import { sessionSpoke } from "@/spoke";
import { useNavigate } from "react-router-dom";

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
  const toast = useToast();
  const toastBg = useColorModeValue("white", "gray.600", "gray.700");
  const toastFg = useColorModeValue("gray.700", "gray.600", "gray.700");
  const green = useColorModeValue("green.500", "green.400", "green.300");
  const red = useColorModeValue("red.500", "red.400", "red.300");
  const dispatch = useAppDispatch();
  const {
    currentQuestionNumber,
    fontSize,
    currentLanguage,
    snapshotByQuestionNumber
  } = useAppSelector((state) => state.editor);
  const { accessToken } = useAppSelector((state) => state.session);
  const navigate = useNavigate();

  const { t } = useTranslation();

  const { deadlineUtc } = useAppSelector((state) => state.editor);
  const [time, setTime] = useState(deadlineUtc ? deadlineUtc - Date.now() : 0);

  useEffect(() => {
    const timer = setInterval(() => {
      setTime((prev: number) => prev - 1000);
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  useEffect(() => {
    const endTimer = setInterval(async () => {
      if (accessToken === null) return;

      await sessionSpoke.endExam({ accessToken });
      navigate("/sam-test");
    }, time);

    return () => clearInterval(endTimer);
  });

  const currentSnapshot = snapshotByQuestionNumber[currentQuestionNumber!];
  const isSubmitted =
    currentSnapshot !== undefined ? currentSnapshot.submissionAccepted : false;
  const isRefactored =
    currentSnapshot !== undefined
      ? currentSnapshot.submissionRefactored
      : false;

  async function handleSubmit() {
    if (currentQuestionNumber === null || accessToken === null) return;

    try {
      const submissionResult = await sessionSpoke.submitSolution({
        // FIXME(elianiva): fix this dumb thing
        accessToken,
        language: 1,
        solution: "",
        scratchPad: "",
        questionNumber: currentQuestionNumber
      });

      dispatch(
        setSnapshot({
          language: currentLanguage,
          questionNumber: currentQuestionNumber,
          scratchPad: currentSnapshot?.scratchPad || "",
          solutionByLanguage: {
            ...currentSnapshot?.solutionByLanguage,
            [currentLanguage]: ""
          },
          submissionAccepted: submissionResult.accepted,
          submissionRefactored: currentSnapshot?.submissionSubmitted || false,
          submissionSubmitted: true,
          testResults: submissionResult.testResults
        })
      );

      const id = toast({
        position: "top-right",
        render: () => (
          <FeedbackToast
            bg={toastBg}
            fg={toastFg}
            green={green}
            red={red}
            isCorrect={submissionResult.accepted}
            onClick={() => toast.close(id!)}
          />
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
          onClick={() => {
            // TODO(elianiva): do we need this?
          }}
          data-tour="topbar-step-7"
        >
          Test
        </Button>
        {!isRefactored && (
          <Button
            px="4"
            colorScheme="blue"
            h="full"
            onClick={() => handleSubmit()}
            data-tour="topbar-step-8"
          >
            {isSubmitted ? "Refactor" : "Submit"}
          </Button>
        )}
      </Flex>
    </Flex>
  );
}
