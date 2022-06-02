import React from "react";
import { useTranslation } from "react-i18next";
import {
  Flex,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs
} from "@chakra-ui/react";
import { useAppDispatch, useAppSelector } from "@/store";
import { useColorModeValue } from "@/hooks";
import { setQuestionTabIndex } from "@/store/slices/codingTestSlice";
import ResultPane from "@/components/CodingTest/ResultPane";
import QuestionPane from "@/components/CodingTest/QuestionPane";

interface ProblemProps {
  bg: string;
  fg: string;
  fgDarker: string;
}

export default function Problem({ bg, fg, fgDarker }: ProblemProps) {
  const dispatch = useAppDispatch();
  const { t } = useTranslation();
  const borderBg = useColorModeValue("gray.300", "gray.500", "gray.600");

  const { questionTabIndex } = useAppSelector((state) => state.codingTest);

  return (
    <Flex
      direction="column"
      flex="1"
      position="relative"
      bg={bg}
      minW="80"
      h="full"
      rounded="md"
      shadow="md"
    >
      <Tabs
        index={questionTabIndex}
        onChange={(idx) => {
          dispatch(setQuestionTabIndex(idx === 1 ? "result" : "question"));
        }}
        isLazy
        h="calc(100% - 2.75rem)"
      >
        <TabList borderColor={borderBg}>
          <Tab color={fgDarker} data-tour="question-step-1">
            {t("translation.translations.ui.prompt")}
          </Tab>
          <Tab color={fgDarker} data-tour="question-step-2">
            {t("translation.translations.ui.your_result")}
          </Tab>
        </TabList>

        <TabPanels h="full">
          <TabPanel p="2" h="full">
            <QuestionPane fg={fg} fgDarker={fgDarker} />
          </TabPanel>
          <TabPanel p="2" h="full">
            <ResultPane fg={fg} fgDarker={fgDarker} />
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Flex>
  );
}
