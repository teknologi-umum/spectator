import React, { useMemo } from "react";
import { useColorModeValue } from "@/hooks";
import { useAppSelector } from "@/store";
import {
  Accordion,
  AccordionButton,
  AccordionIcon,
  AccordionItem,
  AccordionPanel,
  Badge,
  Box,
  Code,
  Flex,
  Text,
  Tooltip
} from "@chakra-ui/react";
import {
  CheckmarkIcon,
  CrossIcon,
  WarningIcon,
  QuestionOutlineIcon
} from "@/icons";
import { ResultCase } from "@/models/TestResult";
import { EditorSnapshot } from "@/models/EditorSnapshot";
import { useTranslation } from "react-i18next";

interface OutputBoxProps {
  expected: string;
  actual: string;
  argument: string;
}
function OutputBox({ expected, actual, argument }: OutputBoxProps) {
  return (
    <Box>
      <Text>
        Caller: <Code>{argument}</Code>
      </Text>
      <Text>
        Expected: <Code>{expected}</Code>
      </Text>
      <Text>
        Actual: <Code>{actual}</Code>
      </Text>
    </Box>
  );
}

interface ResultProps {
  fg: string;
  fgDarker: string;
}

export default function Result({ fg, fgDarker }: ResultProps) {
  const { t } = useTranslation("translation", {
    keyPrefix: "translations"
  });

  const resultBg = useColorModeValue("gray.50", "gray.600", "gray.900");
  const green = useColorModeValue("green.500", "green.400", "green.300");
  const orange = useColorModeValue("orange.500", "orange.400", "orange.300");
  const yellow = useColorModeValue("yellow.500", "yellow.400", "yellow.300");
  const red = useColorModeValue("red.500", "red.400", "red.300");

  const { snapshotByQuestionNumber, currentQuestionNumber } = useAppSelector(
    (state) => state.editor
  );
  const currentSnapshot: EditorSnapshot | undefined =
    snapshotByQuestionNumber[currentQuestionNumber];
  // normalise the output to either an array of result or null
  // instead of dealing with undefined/null/empty array
  const testResults = useMemo(() => {
    if (currentSnapshot === undefined) return null;
    if (currentSnapshot.testResults === null) return null;
    if (currentSnapshot.testResults.length < 1) return null;
    return currentSnapshot.testResults;
  }, [currentSnapshot]);

  const badgeColours = {
    [ResultCase.Passing]: "green",
    [ResultCase.Failing]: "red",
    [ResultCase.RuntimeError]: "orange",
    [ResultCase.CompileError]: "yellow",
    [ResultCase.InvalidInput]: "orange"
  };

  const resultIcons = {
    [ResultCase.Passing]: (
      <Box color={green}>
        <CheckmarkIcon />
      </Box>
    ),
    [ResultCase.Failing]: (
      <Box color={red}>
        <CrossIcon />
      </Box>
    ),
    [ResultCase.RuntimeError]: (
      <Box color={orange}>
        <WarningIcon />
      </Box>
    ),
    [ResultCase.CompileError]: (
      <Box color={yellow}>
        <WarningIcon />
      </Box>
    ),
    [ResultCase.InvalidInput]: (
      <Box color={orange}>
        <CrossIcon />
      </Box>
    )
  };

  const humanisedResultCase = {
    [ResultCase.Passing]: "Passing",
    [ResultCase.Failing]: "Failing",
    [ResultCase.RuntimeError]: "Runtime Error",
    [ResultCase.CompileError]: "Compile Error",
    [ResultCase.InvalidInput]: "Invalid Input"
  };

  return (
    <Box overflowY="auto" p="4" h="full">
      <Accordion allowToggle allowMultiple>
        {testResults === null ? (
          <Text>No tests have been run.</Text>
        ) : (
          testResults.map((testResult, index) => {
            const itemBadgeColour = badgeColours[testResult.resultCase];
            const itemResultIcon = resultIcons[testResult.resultCase];
            const status = humanisedResultCase[testResult.resultCase];

            return (
              <AccordionItem
                key={index}
                border="none"
                background={resultBg}
                mb="3"
                rounded="sm"
                _expanded={{ borderRadius: "sm" }}
              >
                <AccordionButton
                  color={fg}
                  _hover={{ borderRadius: "sm" }}
                  rounded="sm"
                >
                  <Flex gap="2" align="center" flex="1" textAlign="left">
                    {itemResultIcon}
                    <Text fontWeight="bold" color={fgDarker}>
                      Test Result #{index + 1}
                    </Text>
                    <Badge colorScheme={itemBadgeColour}>{status}</Badge>
                    <Tooltip
                      hasArrow
                      label={t(
                        `error_type.${status
                          .toLowerCase()
                          .split(" ")
                          .join("_")}`
                      )}
                    >
                      <span>
                        <QuestionOutlineIcon />
                      </span>
                    </Tooltip>
                  </Flex>
                  <AccordionIcon />
                </AccordionButton>
                <AccordionPanel pb={4} color={fgDarker}>
                  {testResult.resultCase === ResultCase.Passing && (
                    <>
                      <Text>Passed!</Text>
                      <OutputBox
                        expected={testResult.passingTest.expectedStdout}
                        actual={testResult.passingTest.actualStdout}
                        argument={testResult.passingTest.argumentsStdout}
                      />
                    </>
                  )}

                  {testResult.resultCase === ResultCase.RuntimeError && (
                    <Text>{testResult.runtimeError.stderr}</Text>
                  )}

                  {testResult.resultCase === ResultCase.CompileError && (
                    <Text>{testResult.compileError.stderr}</Text>
                  )}

                  {testResult.resultCase === ResultCase.InvalidInput && (
                    <Text>{t("error_type.invalid_input")}</Text>
                  )}

                  {testResult.resultCase === ResultCase.Failing && (
                    <OutputBox
                      expected={testResult.failingTest.expectedStdout}
                      actual={testResult.failingTest.actualStdout}
                      argument={testResult.failingTest.argumentsStdout}
                    />
                  )}
                </AccordionPanel>
              </AccordionItem>
            );
          })
        )}
      </Accordion>
    </Box>
  );
}
