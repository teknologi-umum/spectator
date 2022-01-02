import React, { useEffect, useState } from "react";
import type { UIEventHandler } from "react";
import {
  Heading,
  Tabs,
  TabList,
  TabPanels,
  Tab,
  TabPanel,
  Box
} from "@chakra-ui/react";
import CodeMirror from "@uiw/react-codemirror";
import { lineNumbers } from "@codemirror/gutter";
import { useCodemirrorTheme, useDebounce } from "@/hooks";
import { useAppSelector, useAppDispatch } from "@/store";
import { setScratchPad } from "@/store/slices/editorSlice";

interface ScratchPadProps {
  bg: string;
  onScroll: UIEventHandler<HTMLDivElement>;
}
export default function ScratchPad({ bg, onScroll }: ScratchPadProps) {
  const dispatch = useAppDispatch();
  const [theme, highlightTheme] = useCodemirrorTheme();
  const { currentQuestion } = useAppSelector((state) => state.question);
  const { scratchPads } = useAppSelector((state) => state.editor);

  const [value, setValue] = useState("");
  const debouncedValue = useDebounce(value, 500);

  useEffect(() => {
    const currentScratchPad = scratchPads.find(
      (scratchPad) => scratchPad.questionNo === currentQuestion
    );

    if (currentScratchPad !== undefined) {
      setValue(currentScratchPad.value);
    } else {
      dispatch(
        setScratchPad({
          questionNo: currentQuestion,
          value: ""
        })
      );
    }
  }, []);

  useEffect(() => {
    dispatch(
      setScratchPad({
        questionNo: currentQuestion,
        value: debouncedValue
      })
    );
  }, [debouncedValue]);

  function handleChange(value: string) {
    setValue(value);
  }

  return (
    <Box bg={bg} rounded="md" shadow="md" flex="1" h="full">
      <Tabs isLazy h="full">
        <TabList>
          <Tab>Scratch Pad</Tab>
          <Tab>Output</Tab>
        </TabList>

        <TabPanels h="full">
          <TabPanel p="2" h="full" tabIndex={-1}>
            <CodeMirror
              value={value}
              height="8rem"
              extensions={[highlightTheme, lineNumbers()]}
              theme={theme}
              style={{ height: "calc(100% - 2.75rem)" }}
              onScroll={onScroll}
              onChange={handleChange}
            />
          </TabPanel>
          <TabPanel p="2" tabIndex={-1}>
            <Heading>Sandbox</Heading>
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Box>
  );
}
