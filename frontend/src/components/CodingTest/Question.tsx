import React, { FC, useMemo } from "react";
import { useTranslation } from "react-i18next";
import {
  Box,
  Flex,
  Heading,
  ListItem,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs,
  Text,
  UnorderedList
} from "@chakra-ui/react";
import ReactMarkdown from "react-markdown";
import { useAppDispatch, useAppSelector } from "@/store";
import { useColorModeValue } from "@/hooks";
import Result from "@/components/CodingTest/Result";
import { setQuestionTabIndex } from "@/store/slices/codingTestSlice";

interface QuestionProps {
  bg: string;
  fg: string;
  fgDarker: string;
}

function buildParagraph(color: string) {
  const MDParagraph: FC = ({ children }) => (
    <Text fontSize="16" lineHeight="6" color={color} py="2">
      {children}
    </Text>
  );
  return MDParagraph;
}

function buildUnorderedList() {
  const MDUnorderedList: FC = ({ children }) => (
    <UnorderedList>{children}</UnorderedList>
  );
  return MDUnorderedList;
}

function buildListItem() {
  const MDListItem: FC = ({ children }) => <ListItem>{children}</ListItem>;
  return MDListItem;
}

function buildPre(codeBg: string, fg: string) {
  const Pre: FC = ({ children }) => (
    <Box
      as="pre"
      bg={codeBg}
      color={fg}
      p="3"
      rounded="md"
      mb="4"
      overflowX="auto"
      fontSize="14"
    >
      {children}
    </Box>
  );
  return Pre;
}

export default function Question({ bg, fg, fgDarker }: QuestionProps) {
  const dispatch = useAppDispatch();
  const { t } = useTranslation();
  const codeBg = useColorModeValue("gray.200", "gray.800", "gray.900");
  const borderBg = useColorModeValue("gray.300", "gray.500", "gray.600");

  const { currentQuestionNumber, snapshotByQuestionNumber } = useAppSelector(
    (state) => state.editor
  );
  const { questionTabIndex } = useAppSelector((state) => state.codingTest);

  const currentSnapshot = useMemo(
    () => snapshotByQuestionNumber[currentQuestionNumber],
    [snapshotByQuestionNumber, currentQuestionNumber]
  );

  const isResultTabDisabled = useMemo(
    () =>
      currentSnapshot?.testResults === null ||
      currentSnapshot?.testResults?.length > 0,
    [currentSnapshot?.testResults]
  );

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
          <Tab
            color={fgDarker}
            isDisabled={isResultTabDisabled}
            data-tour="question-step-2"
          >
            {t("translation.translations.ui.your_result")}
          </Tab>
        </TabList>

        <TabPanels h="full">
          <TabPanel p="2" h="full">
            <Box p="4" overflowY="auto" flex="1" h="full">
              <Heading size="lg" color={fg}>
                {t(`question.questions.${currentQuestionNumber - 1}.title`)}
              </Heading>
              <ReactMarkdown
                components={{
                  p: buildParagraph(fgDarker),
                  ul: buildUnorderedList(),
                  li: buildListItem(),
                  pre: buildPre(codeBg, fg)
                }}
              >
                {t(`question.questions.${currentQuestionNumber - 1}.question`)}
              </ReactMarkdown>
            </Box>
          </TabPanel>
          <TabPanel p="2" h="full">
            <Result fg={fg} fgDarker={fgDarker} />
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Flex>
  );
}
