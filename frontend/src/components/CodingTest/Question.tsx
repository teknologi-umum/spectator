import React from "react";
import type { UIEventHandler } from "react";
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
  UnorderedList,
  useColorModeValue
} from "@chakra-ui/react";
import ReactMarkdown from "react-markdown";
// TODO: this should be automatically inferred (en/id) when we have proper i18n
import { questions } from "@/data/en/questions.json";
import { useAppSelector } from "@/store";
import type { UIEventHandler } from "react";

interface QuestionProps {
  bg: string;
  fg: string;
  onScroll: UIEventHandler<HTMLDivElement>;
}

export default function Question({ bg, fg, onScroll }: QuestionProps) {
  const codeBg = useColorModeValue("gray.200", "gray.800");
  const { currentQuestion } = useAppSelector((state) => state.question);

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
        <TabList>
          <Tab>Prompt</Tab>
          <Tab>Your Result</Tab>
        </TabList>

        <TabPanels h="full">
          <TabPanel p="2" h="full">
            <Box p="4" overflowY="auto" flex="1" h="full" onScroll={onScroll}>
              <Heading size="lg" color={fg}>
                {questions[currentQuestion].title}
              </Heading>
              <ReactMarkdown
                components={{
                  p: ({ children }) => (
                    <Text fontSize="16" lineHeight="6" color={fg} py="2">
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
                {questions[currentQuestion].question}
              </ReactMarkdown>
            </Box>
          </TabPanel>
          <TabPanel p="2">
            <Heading>Result</Heading>
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Flex>
  );
}
