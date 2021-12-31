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
import { questions } from "@/data/questions.json";
import { useAppSelector } from "@/store";
import type { InitialState as QuestionState } from "@/store/slices/questionSlice/types";
import { useColorModeValue } from "@/hooks/";

interface QuestionProps {
  bg: string;
  fg: string;
  fgDarker: string;
}

export default function Question({ bg, fg, fgDarker }: QuestionProps) {
  const codeBg = useColorModeValue("gray.200", "gray.500", "gray.700");
  const { currentQuestion } = useAppSelector<QuestionState>(
    (state) => state.question
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
      {/* TODO(elianiva): should automatically switch to 'your result' after pressing submit */}
      <Tabs h="calc(100% - 2.75rem)" isLazy>
        <TabList>
          <Tab color={fgDarker}>Prompt</Tab>
          <Tab color={fgDarker}>Your Result</Tab>
        </TabList>

        <TabPanels h="full">
          <TabPanel p="2" h="full">
            <Box p="4" overflowY="auto" flex="1" h="full">
              <Heading size="lg" color={fg}>
                {questions[currentQuestion].title}
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
