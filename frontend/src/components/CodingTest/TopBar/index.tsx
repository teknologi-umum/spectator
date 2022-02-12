import React, { useEffect, useState } from "react";
import {
  Button,
  Flex,
  MenuItem,
  MenuItemOption,
  MenuOptionGroup,
  Text
} from "@chakra-ui/react";
import {
  setFontSize,
  setLanguage,
  setSnapshot
} from "@/store/slices/editorSlice";
import type { Language } from "@/models/Language";
import { useAppDispatch, useAppSelector } from "@/store";
import { mutate } from "@/utils/fakeSubmissionCallback";
import { useTranslation } from "react-i18next";
import { jwtDecode } from "@/utils/jwtDecode";
import { ClockIcon } from "@/icons";
import {
  MenuDropdown,
  ThemeButton,
  LocaleButton
} from "@/components/CodingTest";

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
  fg: string;
}

export default function TopBar({ bg, fg }: MenuProps) {
  const dispatch = useAppDispatch();
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

  const recordedSubmission = snapshotByQuestionNumber[currentQuestionNumber!];
  const isSubmitted =
    recordedSubmission !== undefined
      ? recordedSubmission.submissionAccepted
      : false;
  const isRefactored =
    recordedSubmission !== undefined
      ? recordedSubmission.submissionRefactored
      : false;

  function handleSubmit() {
    if (currentQuestionNumber === null) return;

    const currentSnapshot = snapshotByQuestionNumber[currentQuestionNumber];

    mutate(currentSnapshot, {
      onSuccess: (res) => {
        dispatch(
          setSnapshot({
            ...res.data,
            submissionSubmitted: true
          })
        );
      }
    });
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
        <ThemeButton bg={bg} fg={fg} data-tour="topbar-step-2"/>
        <MenuDropdown
          dropdownWidth="10rem"
          bg={bg}
          fg={fg}
          title={currentLanguage}
          data-tour="topbar-step-3"
        >
          <MenuOptionGroup
            type="radio"
            value={isSubmitted ? recordedSubmission.language : currentLanguage}
            onChange={(value) => dispatch(setLanguage(value as Language))}
          >
            {LANGUAGES.map((lang, idx) => (
              <MenuItemOption
                textTransform="capitalize"
                key={idx}
                value={lang}
                isDisabled={
                  isSubmitted ? lang !== recordedSubmission?.language : false
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
          onClick={() => {
            // TODO(elianiva): implement proper surrender logic properly
            //                 it's now temporarily used for previous question
            //                 to make testing easier
            // dispatch(prevQuestion());
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
            // TODO(elianiva): send the code to backend for execution
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
            onClick={() => {
              // TODO(elianiva): only allow to continue when they have the correct answer
              // dispatch(nextQuestion());

              handleSubmit();
            }}
            data-tour="topbar-step-8"
          >
            {isSubmitted ? "Refactor" : "Submit"}
          </Button>
        )}
      </Flex>
    </Flex>
  );
}
