import React, { useEffect, useMemo, useRef, useState } from "react";
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
import { useAppSelector } from "@/store";
import { useColorModeValue } from "@/hooks";
import { useTranslation } from "react-i18next";
import Result from "./Result";

interface QuestionProps {
  bg: string;
  fg: string;
  fgDarker: string;
}

export default function Question({ bg, fg, fgDarker }: QuestionProps) {
  const isMounted = useRef(false);
  const [tabIndex, setTabIndex] = useState(0);

  const codeBg = useColorModeValue("gray.200", "gray.800", "gray.900");
  const borderBg = useColorModeValue("gray.300", "gray.500", "gray.600");
  const { currentQuestionNumber, snapshotByQuestionNumber } = useAppSelector(
    (state) => state.editor
  );
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

  useEffect(() => {
    if (isMounted.current) {
      // move to the result tab whenever the current submission changes
      setTabIndex(1);
    } else {
      isMounted.current = true;
    }
  }, [currentSnapshot?.testResults]);

  const { t } = useTranslation();

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
      {/* TODO(elianiva): should automatically switch to 'your result' after pressing submit */}
      <Tabs
        index={tabIndex}
        onChange={(idx) => setTabIndex(idx)}
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
                  p: ({ children }) => (
                    <Text fontSize="16" lineHeight="6" color={fgDarker} py="2">
                      {children}
                    </Text>
                  ),
                  ul: ({ children }) => (
                    <UnorderedList>{children}</UnorderedList>
                  ),
                  li: ({ children }) => <ListItem>{children}</ListItem>,
                  pre: ({ children }) => (
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
                  )
                }}
              >
                {t(`question.questions.${currentQuestionNumber - 1}.question`)}
              </ReactMarkdown>
            </Box>
            1
          </TabPanel>
          <TabPanel p="2" h="full">
            <Result fg={fg} fgDarker={fgDarker} />
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Flex>
  );
}
