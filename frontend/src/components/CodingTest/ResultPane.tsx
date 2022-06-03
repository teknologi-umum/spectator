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
import { EditorSnapshot } from "@/models/EditorSnapshot";

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
  const resultBg = useColorModeValue("gray.50", "gray.600", "gray.900");
  const green = useColorModeValue("green.500", "green.400", "green.300");
  const orange = useColorModeValue("orange.500", "orange.400", "orange.300");
  const yellow = useColorModeValue("yellow.500", "yellow.400", "yellow.300");
  const red = useColorModeValue("red.500", "red.400", "red.300");

  const { snapshotByQuestionNumber, currentQuestionNumber } = useAppSelector(
    (state) => state.editor
  );
  const currentSnapshot = useMemo<EditorSnapshot | undefined>(
    () => snapshotByQuestionNumber[currentQuestionNumber],
    [currentQuestionNumber]
  );
  // normalise the output to either an array of result or null
  // instead of dealing with undefined/null/empty array
  const testResults = useMemo(() => {
    if (currentSnapshot === undefined) return null;
    if (currentSnapshot.testResults === null) return null;
    if (currentSnapshot.testResults.length < 1) return null;
    return currentSnapshot.testResults;
  }, [currentSnapshot]);

  const badgeColours = {
    Passing: "green",
    Failing: "red",
    RuntimeError: "orange",
    CompileError: "yellow"
  };

  const resultIcons = {
    Passing: (
      <Box color={green}>
        <CheckmarkIcon />
      </Box>
    ),
    Failing: (
      <Box color={red}>
        <CrossIcon />
      </Box>
    ),
    RuntimeError: (
      <Box color={orange}>
        <WarningIcon />
      </Box>
    ),
    CompileError: (
      <Box color={yellow}>
        <WarningIcon />
      </Box>
    )
  };

  const humanisedResultCase = {
    Passing: "Passing",
    Failing: "Failing",
    RuntimeError: "Runtime Error",
    CompileError: "Compile Error"
  };

  return (
    <Box overflowY="auto" p="4" h="full">
      <Accordion allowToggle allowMultiple>
        {testResults === null ? (
          <Text>No tests have been run.</Text>
        ) : (
          testResults.map((testResult, index) => {
            const itemBadgeColour = badgeColours[testResult.status];
            const itemResultIcon = resultIcons[testResult.status];
            const status = humanisedResultCase[testResult.status];

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
                  </Flex>
                  <AccordionIcon />
                </AccordionButton>
                <AccordionPanel pb={4} color={fgDarker}>
                  {testResult.status === "Passing" && <Text>Passed!</Text>}

                  {testResult.status === "RuntimeError" && (
                    <Text>{testResult.result.stderr}</Text>
                  )}

                  {testResult.status === "CompileError" && (
                    <Text>{testResult.result.stderr}</Text>
                  )}

                  {testResult.status === "Failing" && (
                    <OutputBox
                      expected={testResult.result.expectedStdout}
                      actual={testResult.result.actualStdout}
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
