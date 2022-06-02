import React from "react";
import type { FC } from "react";
import { useTranslation } from "react-i18next";
import {
  Box,
  Heading,
  ListItem,
  Text,
  UnorderedList
} from "@chakra-ui/react";
import ReactMarkdown from "react-markdown";
import { useAppSelector } from "@/store";
import { useColorModeValue } from "@/hooks";

interface QuestionPaneProps {
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

export default function QuestionPane({ fg, fgDarker }: QuestionPaneProps) {
  const { t } = useTranslation();
  const codeBg = useColorModeValue("gray.200", "gray.800", "gray.900");

  const { currentQuestionNumber } = useAppSelector((state) => state.editor);

  return (
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
  );
}
