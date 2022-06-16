import React from "react";
import type { FC } from "react";
import { useTranslation } from "react-i18next";
import {
  Box,
  Heading,
  ListItem,
  Text,
  UnorderedList,
  Badge
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
  const { t } = useTranslation("question", {
    keyPrefix: "questions"
  });
  const codeBg = useColorModeValue("gray.200", "gray.800", "gray.900");

  const { currentQuestionNumber, snapshotByQuestionNumber } = useAppSelector(
    (state) => state.editor
  );

  const snapshot = snapshotByQuestionNumber[currentQuestionNumber];

  return (
    <Box p="4" overflowY="auto" flex="1" h="full">
      <Heading size="lg" color={fg}>
        {t(`${currentQuestionNumber - 1}.title`)} {"  "}
        {!snapshot?.submissionSubmitted && (
          <Badge fontSize="1rem" variant="outline">
            NO ATTEMPT
          </Badge>
        )}
        {snapshot?.submissionSubmitted && snapshot?.submissionAccepted && (
          <Badge fontSize="1rem" variant="subtle" colorScheme='green'>
            ACCEPTED
          </Badge>
        )}
        {snapshot?.submissionSubmitted && !snapshot?.submissionAccepted && (
          <Badge fontSize="1rem" variant="subtle" colorScheme='red'>
            REJECTED
          </Badge>
        )}
      </Heading>
      <ReactMarkdown
        components={{
          p: buildParagraph(fgDarker),
          ul: buildUnorderedList(),
          li: buildListItem(),
          pre: buildPre(codeBg, fg)
        }}
      >
        {t(`${currentQuestionNumber - 1}.question`)}
      </ReactMarkdown>
    </Box>
  );
}
