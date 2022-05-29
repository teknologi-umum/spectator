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
  Text
} from "@chakra-ui/react";
import { CheckmarkIcon, CrossIcon, WarningIcon } from "@/icons";
import { ResultCase } from "@/models/TestResult";

interface OutputBoxProps {
  expected: string;
  actual: string;
}
function OutputBox({ expected, actual }: OutputBoxProps) {
  return (
    <Box>
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
  const { snapshotByQuestionNumber, currentQuestionNumber } = useAppSelector(
    (state) => state.editor
  );
  const currentSnapshot = useMemo(
    () => snapshotByQuestionNumber[currentQuestionNumber],
    [currentQuestionNumber]
  );
  const resultBg = useColorModeValue("gray.50", "gray.600", "gray.900");
  const green = useColorModeValue("green.500", "green.400", "green.300");
  const orange = useColorModeValue("orange.500", "orange.400", "orange.300");
  const yellow = useColorModeValue("yellow.500", "yellow.400", "yellow.300");
  const red = useColorModeValue("red.500", "red.400", "red.300");

  const badgeColours = {
    [ResultCase.Passing]: "green",
    [ResultCase.Failing]: "red",
    [ResultCase.RuntimeError]: "orange",
    [ResultCase.CompileError]: "yellow"
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
    )
  };

  const humanisedResultCase = {
    [ResultCase.Passing]: "Passing",
    [ResultCase.Failing]: "Failing",
    [ResultCase.RuntimeError]: "Runtime Error",
    [ResultCase.CompileError]: "Compile Error"
  };

  return (
    <Box overflowY="auto" p="4" h="full">
      <Accordion allowToggle allowMultiple>
        {currentSnapshot?.testResults &&
          currentSnapshot.testResults.map((testResult, index) => {
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
                      Test Result #{testResult.testNumber}
                    </Text>
                    <Badge colorScheme={itemBadgeColour}>{status}</Badge>
                  </Flex>
                  <AccordionIcon />
                </AccordionButton>
                <AccordionPanel pb={4} color={fgDarker}>
                  {testResult.resultCase === ResultCase.Passing && (
                    <Text>Passed!</Text>
                  )}

                  {testResult.resultCase === ResultCase.RuntimeError && (
                    <Text>{testResult.runtimeError.stderr}</Text>
                  )}

                  {testResult.resultCase === ResultCase.CompileError && (
                    <Text>{testResult.compileError.stderr}</Text>
                  )}

                  {testResult.resultCase === ResultCase.Failing && (
                    <OutputBox
                      expected={testResult.failingTest.expectedStdout}
                      actual={testResult.failingTest.actualStdout}
                    />
                  )}
                </AccordionPanel>
              </AccordionItem>
            );
          })}
      </Accordion>
    </Box>
  );
}
