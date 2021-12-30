import { Tabs, TabList, TabPanels, Tab, TabPanel, Box } from "@chakra-ui/react";
import CodeMirror, { keymap } from "@uiw/react-codemirror";
import { defaultKeymap } from "@codemirror/commands";
import { lineNumbers } from "@codemirror/gutter";
import { javascript } from "@codemirror/lang-javascript";
import { php } from "@codemirror/lang-php";
import { java } from "@codemirror/lang-java";
import { cpp } from "@codemirror/lang-cpp";
import { python } from "@codemirror/lang-python";
import { useCodemirrorTheme } from "@/hooks";
import { questions } from "@/data/questions.json";
import { useAppSelector, useAppDispatch } from "@/store";
import { setSolution } from "@/store/slices/editorSlice";
import type { InitialState as EditorState } from "@/store/slices/editorSlice/types";
import type { InitialState as QuestionState } from "@/store/slices/questionSlice/types";
import { UIEventHandler, useEffect, useState, useMemo } from "react";
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
  onScroll: UIEventHandler<HTMLDivElement>;
}

export default function Editor({ bg, onScroll }: EditorProps) {
  const dispatch = useAppDispatch();
  const [theme, highlightTheme] = useCodemirrorTheme();
  const { currentQuestion } = useAppSelector<QuestionState>(
    (state) => state.question
  );
  const { currentLanguage } = useAppSelector<EditorState>(
    (state) => state.editor
  );
  // memoized the question
  const stringifiedQuestion = useMemo(() => {
    return JSON.stringify(
      questions[currentQuestion].templates[currentLanguage]
    );
  }, [currentQuestion, currentLanguage]);

  const [innerCode, setInnerCode] = useState(stringifiedQuestion);
  const debouncedInnerCode = useDebounce(innerCode, 1000);

  useEffect(() => {
    // need further discussion about scratchPad payload since
    // editor and scratchpad are two separate components
    dispatch(
      setSolution({
        questionNo: currentQuestion,
        language: currentLanguage,
        code: JSON.stringify(debouncedInnerCode),
        scratchPad: ""
      })
    );
  }, [debouncedInnerCode]);

  function handleChange(value: string) {
    setInnerCode(JSON.stringify(value));
  }

  return (
    <Box bg={bg} rounded="md" shadow="md" flex="1" h="full">
      <Tabs h="full">
        <TabList>
          <Tab>Your Solution</Tab>
        </TabList>
        <TabPanels h="full">
          <TabPanel p="2" h="full" position="relative" tabIndex={-1}>
            <CodeMirror
              value={innerCode ? JSON.parse(innerCode) : ""}
              extensions={[
                highlightTheme,
                lineNumbers(),
                LANGUAGES[currentLanguage],
                keymap.of([
                  ...defaultKeymap,
                  {
                    key: "Ctrl-c",
                    run: () => {
                      /* noop */
                      return true;
                    },
                    preventDefault: true
                  },
                  {
                    key: "Ctrl-v",
                    run: () => {
                      /* noop */
                      return true;
                    },
                    preventDefault: true
                  }
                ])
              ]}
              theme={theme}
              style={{ height: "calc(100% - 2.75rem)" }}
              onScroll={onScroll}
              onChange={handleChange}
            />
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Box>
  );
}
