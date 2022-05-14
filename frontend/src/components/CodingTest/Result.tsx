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

  const badgeColours = useMemo(
    () => ({
      passingTest: "green",
      runtimeError: "orange",
      compileError: "yellow",
      failingTest: "red"
    }),
    []
  );

  const resultIcons = useMemo(
    () => ({
      passingTest: (
        <Box color={green}>
          <CheckmarkIcon />
        </Box>
      ),
      runtimeError: (
        <Box color={orange}>
          <WarningIcon />
        </Box>
      ),
      compileError: (
        <Box color={yellow}>
          <WarningIcon />
        </Box>
      ),
      failingTest: (
        <Box color={red}>
          <CrossIcon />
        </Box>
      )
    }),
    []
  );

  function humanizeResultStatus(status: string) {
    return status.replace(/([A-Z]+)/g, " $1").replace(/([A-Z][a-z])/g, " $1");
  }

  return (
    <Box overflowY="auto" p="4" h="full">
      <Accordion allowToggle allowMultiple>
        {currentSnapshot?.testResults &&
          currentSnapshot.testResults
            .filter((testResult) => testResult.result.oneofKind !== undefined)
            .map((testResult, index) => {
              // the compiler isn't smart enough to know that we already
              // filtered out `oneOfKind`s that are undefined so we still have to 
              // do a non-null assertion
              const status = testResult.result.oneofKind!;

              const humanizedStatus = humanizeResultStatus(status);
              const itemBadgeColour = badgeColours[status];
              const itemResultIcon = resultIcons[status];

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
                      <Badge colorScheme={itemBadgeColour}>
                        {humanizedStatus}
                      </Badge>
                    </Flex>
                    <AccordionIcon />
                  </AccordionButton>
                  <AccordionPanel pb={4} color={fgDarker}>
                    {testResult.result.oneofKind === "passingTest" && (
                      <Text>Passed!</Text>
                    )}

                    {testResult.result.oneofKind === "runtimeError" && (
                      <Text>{testResult.result.runtimeError.stderr}</Text>
                    )}

                    {testResult.result.oneofKind === "compileError" && (
                      <Text>{testResult.result.compileError.stderr}</Text>
                    )}

                    {testResult.result.oneofKind === "failingTest" && (
                      <OutputBox
                        expected={testResult.result.failingTest.expectedStdout}
                        actual={testResult.result.failingTest.actualStdout}
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
