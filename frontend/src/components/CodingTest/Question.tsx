import React, { useMemo } from "react";
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
import { UIEventHandler } from "react";
import { useTranslation } from "react-i18next";
import Result from "./Result";

interface QuestionProps {
  bg: string;
  fg: string;
  fgDarker: string;
  onScroll: UIEventHandler<HTMLDivElement>;
}

export default function Question({
  bg,
  fg,
  fgDarker,
  onScroll
}: QuestionProps) {
  const codeBg = useColorModeValue("gray.200", "gray.800", "gray.900");
  const borderBg = useColorModeValue("gray.300", "gray.400", "gray.400");
  const { currentQuestionNumber, snapshotByQuestionNumber } = useAppSelector(
    (state) => state.editor
  );
  const currentSnapshot = useMemo(
    () => snapshotByQuestionNumber[currentQuestionNumber],
    [currentQuestionNumber]
  );

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
      <Tabs h="calc(100% - 2.75rem)" isLazy>
        <TabList borderColor={borderBg}>
          <Tab color={fgDarker}>{t("translation.translations.ui.prompt")}</Tab>
          <Tab
            color={fgDarker}
            isDisabled={currentSnapshot?.testResults !== null}
          >
            {t("translation.translations.ui.your_result")}
          </Tab>
        </TabList>

        <TabPanels h="full">
          <TabPanel p="2" h="full">
            <Box p="4" overflowY="auto" flex="1" h="full" onScroll={onScroll}>
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
