import { Tabs, TabList, TabPanels, Tab, TabPanel, Box } from "@chakra-ui/react";
import CodeMirror from "@uiw/react-codemirror";
import { lineNumbers } from "@codemirror/gutter";
import { javascript } from "@codemirror/lang-javascript";
import { php } from "@codemirror/lang-php";
import { java } from "@codemirror/lang-java";
import { cpp } from "@codemirror/lang-cpp";
import { python } from "@codemirror/lang-python";
import { useCodemirrorTheme } from "@/hooks";
// TODO: this should be automatically inferred (en/id) when we have proper i18n
import { questions } from "@/data/en/questions.json";
import { useAppSelector } from "@/store";
import type { InitialState as EditorState } from "@/store/slices/editorSlice/types";
import type { InitialState as QuestionState } from "@/store/slices/questionSlice/types";

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
  const [theme, highlightTheme] = useCodemirrorTheme();
  const { currentQuestion } = useAppSelector<QuestionState>(
    (state) => state.question
  );
  const { currentLanguage } = useAppSelector<EditorState>(
    (state) => state.editor
  );

  return (
    <Box bg={bg} rounded="md" shadow="md" flex="1" h="full">
      <Tabs h="full">
        <TabList>
          <Tab>Your Solution</Tab>
        </TabList>
        <TabPanels h="full">
          <TabPanel p="2" h="full" position="relative" tabIndex={-1}>
            <CodeMirror
              value={questions[currentQuestion].templates[currentLanguage]}
              extensions={[
                highlightTheme,
                lineNumbers(),
                LANGUAGES[currentLanguage]
              ]}
              theme={theme}
              style={{ height: "calc(100% - 2.75rem)" }}
            />
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Box>
  );
}
