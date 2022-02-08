import React, { useEffect, useState } from "react";
import { Button, Flex, Select, Text } from "@chakra-ui/react";
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
import { mutate } from "@/utils/fakeSubmissionCallback";
import { useTranslation } from "react-i18next";
import { jwtDecode } from "@/utils/jwtDecode";
import { ClockIcon } from "@/icons";

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
