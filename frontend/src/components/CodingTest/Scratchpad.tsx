import React, { useEffect, useState } from "react";
import { Tabs, TabList, TabPanels, Tab, TabPanel, Box } from "@chakra-ui/react";
import CodeMirror from "@uiw/react-codemirror";
import { lineNumbers } from "@codemirror/view";
import { useCodemirrorTheme, useColorModeValue, useDebounce } from "@/hooks";
import { useAppSelector, useAppDispatch } from "@/store";
import { setScratchPad } from "@/store/slices/editorSlice";
import { useTranslation } from "react-i18next";

interface ScratchPadProps {
  bg: string;
}
export default function ScratchPad({ bg }: ScratchPadProps) {
  const { t } = useTranslation("translation", {
    keyPrefix: "translations.ui"
  });
  const dispatch = useAppDispatch();
  const [theme, highlightTheme] = useCodemirrorTheme();
  const borderBg = useColorModeValue("gray.300", "gray.500", "gray.600");
  const fgDarker = useColorModeValue("gray.700", "gray.400", "gray.400");
  const { currentQuestionNumber, snapshotByQuestionNumber } = useAppSelector(
    (state) => state.editor
  );

  const [value, setValue] = useState("");
  const debouncedValue = useDebounce(value, 500);


  useEffect(() => {
    const currentScratchPad =
      snapshotByQuestionNumber[currentQuestionNumber!]?.scratchPad;

    if (currentScratchPad !== undefined) {
      setValue(currentScratchPad);
    } else {
      dispatch(setScratchPad(""));
    }
  }, []);

  useEffect(() => {
    dispatch(setScratchPad(debouncedValue));
  }, [debouncedValue]);

  return (
    <Box
      bg={bg}
      rounded="md"
      shadow="md"
      flex="1"
      h="full"
      data-tour="scratchpad-step-1"
    >
      <Tabs isLazy h="full">
        <TabList borderColor={borderBg} color={fgDarker}>
          <Tab>{t("scratchpad")}</Tab>
        </TabList>
        <TabPanels h="full">
          <TabPanel p="2" h="full" tabIndex={-1}>
            <CodeMirror
              value={value}
              height="8rem"
              extensions={[highlightTheme, lineNumbers()]}
              theme={theme}
              style={{ height: "calc(100% - 2.75rem)" }}
              onChange={(v) => setValue(v)}
            />
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Box>
  );
}
