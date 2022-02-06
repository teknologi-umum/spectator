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

  return (
    <Box overflowY="auto" p="4" h="full">
      <Accordion allowToggle allowMultiple>
        {currentSnapshot?.testResults &&
          currentSnapshot.testResults.map((testResult, index) => {
            // replace `PascalCase` with `Sentence Case`
            const status = testResult.status
              .replace(/([A-Z]+)/g, " $1")
              .replace(/([A-Z][a-z])/g, " $1");

            const badgeColour = {
              Passing: "green",
              RuntimeError: "orange",
              CompileError: "yellow",
              Failing: "red"
            }[testResult.status];

            const resultIcon = {
              Passing: <Box color={green}><CheckmarkIcon /></Box>,
              RuntimeError: <Box color={orange}><WarningIcon /></Box>,
              CompileError: <Box color={yellow}><WarningIcon /></Box>,
              Failing: <Box color={red}><CrossIcon /></Box>
            }[testResult.status];

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
                    {resultIcon}
                    <Text fontWeight="bold" color={fgDarker}>
                      Test Result #{testResult.testNumber}
                    </Text>
                    <Badge colorScheme={badgeColour}>{status}</Badge>
                  </Flex>
                  <AccordionIcon />
                </AccordionButton>
                <AccordionPanel pb={4} color={fgDarker}>
                  {testResult.status === "Passing" && <Text>Passed!</Text>}

                  {(testResult.status === "CompileError" ||
                    testResult.status === "RuntimeError") && (
                    <Text>{testResult.stderr}</Text>
                  )}

                  {testResult.status === "Failing" && (
                    <Box>
                      <Text>
                        Expected: <Code>{testResult.expectedStdout}</Code>
                      </Text>
                      <Text>
                        Actual: <Code>{testResult.actualStdout}</Code>
                      </Text>
                    </Box>
                  )}
                </AccordionPanel>
              </AccordionItem>
            );
          })}
      </Accordion>
    </Box>
  );
}
