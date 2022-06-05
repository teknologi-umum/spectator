import React, { useEffect, useState, useMemo } from "react";
import { useTranslation } from "react-i18next";
import { Tabs, TabList, TabPanels, Tab, TabPanel, Box } from "@chakra-ui/react";
import CodeMirror, { keymap } from "@uiw/react-codemirror";
import { defaultKeymap } from "@codemirror/commands";
import { indentUnit } from "@codemirror/language";
import { lineNumbers } from "@codemirror/view";
import { javascript } from "@codemirror/lang-javascript";
import { php } from "@codemirror/lang-php";
import { java } from "@codemirror/lang-java";
import { cpp } from "@codemirror/lang-cpp";
import { python } from "@codemirror/lang-python";
import { useCodemirrorTheme, useColorModeValue } from "@/hooks";
import { useAppSelector, useAppDispatch } from "@/store";
import { setSolution } from "@/store/slices/editorSlice";
import { useDebounce } from "@/hooks";

const cLike = cpp();
const LANGUAGES = {
  java: java(),
  javascript: javascript({ typescript: true }),
  php: php({ plain: true }),
  cpp: cLike,
  c: cLike,
  python: python()
};

interface EditorProps {
  bg: string;
}

export default function Editor({ bg }: EditorProps) {
  const dispatch = useAppDispatch();
  const { t } = useTranslation("question", {
    keyPrefix: "questions"
  });
  const [theme, highlightTheme] = useCodemirrorTheme();
  const borderBg = useColorModeValue("gray.300", "gray.500", "gray.600");

  const [code, setCode] = useState("");
  const debouncedCode = useDebounce(code, 500);

  const { currentQuestionNumber, currentLanguage, snapshotByQuestionNumber } =
    useAppSelector((state) => state.editor);

  const boilerplate = useMemo(
    () => t(`${currentQuestionNumber - 1}.templates.${currentLanguage}`),
    [currentQuestionNumber, currentLanguage]
  );

  const currentSolution = useMemo(() => {
    if (currentQuestionNumber === null) {
      return null;
    }

    const currentSnapshot = snapshotByQuestionNumber[currentQuestionNumber];
    return currentSnapshot?.solutionByLanguage[currentLanguage];
  }, [currentQuestionNumber, currentLanguage]);

  // at first render, we have to check if the data of current solution
  // already persisted. If so, we assign it with setCode.
  // else, we assign it with boilerplate and dispatch to persist store at the same time
  useEffect(() => {
    if (
      currentSolution !== null &&
      currentSolution !== undefined &&
      currentSolution !== ""
    ) {
      setCode(currentSolution);
      return;
    }

    setCode(boilerplate);
    dispatch(setSolution(boilerplate));
  }, [currentQuestionNumber, currentLanguage]);

  useEffect(() => {
    dispatch(setSolution(debouncedCode));
  }, [debouncedCode]);

  function handleChange(value: string) {
    setCode(value);
  }

  function noop() {
    return true;
  }

  return (
    <Box
      bg={bg}
      rounded="md"
      shadow="md"
      flex="1"
      h="full"
      data-tour="editor-step-1"
    >
      <Tabs h="full">
        <TabList borderColor={borderBg}>
          <Tab>Your Solution</Tab>
        </TabList>
        <TabPanels h="full">
          <TabPanel p="2" h="full" position="relative" tabIndex={-1}>
            <CodeMirror
              value={code}
              extensions={[
                highlightTheme,
                lineNumbers(),
                indentUnit.of(" ".repeat(4)), // 4 spaces indentation
                LANGUAGES[currentLanguage],
                keymap.of([
                  ...defaultKeymap,
                  {
                    key: "Ctrl-c",
                    run: noop,
                    preventDefault: true
                  },
                  {
                    key: "Ctrl-v",
                    run: noop,
                    preventDefault: true
                  }
                ])
              ]}
              theme={theme}
              style={{ height: "calc(100% - 2.75rem)" }}
              onChange={handleChange}
            />
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Box>
  );
}
