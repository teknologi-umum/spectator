import { useColorModeValue } from "@/hooks";
import { EditorState } from "@/models/EditorState";
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
import React from "react";

// TODO(elianiva): replace this with a data coming from redux
const FAKE_RESULTS: EditorState = {
  currentQuestionNumber: 1,
  fontSize: 14,
  questions: [],
  currentLanguage: "javascript",
  deadlineUtc: 0,
  snapshotByQuestionNumber: {
    1: {
      questionNumber: 1,
      language: "javascript",
      scratchPad: "",
      submissionAccepted: false,
      submissionRefactored: false,
      submissionSubmitted: true,
      solutionByLanguage: {
        javascript: "function foo() { return 1; }",
        java: "",
        c: "",
        cpp: "",
        python: "",
        php: ""
      },
      testResults: [
        { testNumber: 1, status: "Passing" },
        {
          testNumber: 2,
          status: "RuntimeError",
          stderr: "Trying to access a non-existent variable"
        },
        {
          testNumber: 3,
          status: "CompileError",
          stderr:
            "Failed to compile: error: 'yeet' was not declared in this scope"
        },
        {
          testNumber: 4,
          status: "Failing",
          expectedStdout: "2",
          actualStdout: "1"
        },
        {
          testNumber: 5,
          status: "Failing",
          expectedStdout: "{ \"foo\": \"bar\" }",
          actualStdout: "1"
        },
        { testNumber: 6, status: "Passing" }
      ]
    }
  }
};

interface ResultProps {
  fg: string;
  fgDarker: string;
}

export default function Result({ fg, fgDarker }: ResultProps) {
  const resultBg = useColorModeValue("gray.50", "gray.600", "gray.900");

  return (
    <Box>
      <Accordion allowToggle allowMultiple>
        {FAKE_RESULTS.snapshotByQuestionNumber[1].testResults!.map(
          (testResult, index) => {
            // replace `PascalCase` with `Sentence Case`
            const status = testResult.status
              .replace(/([A-Z]+)/g, " $1")
              .replace(/([A-Z][a-z])/g, " $1");

            const badgeColour =
              testResult.status === "Passing"
                ? "green"
                : testResult.status === "CompileError"
                  ? "orange"
                  : "red";

            return (
              <AccordionItem
                key={index}
                border="none"
                background={resultBg}
                mb="3"
                rounded="sm"
              >
                <AccordionButton color={fg}>
                  <Flex gap="2" align="center" flex="1" textAlign="left">
                    <Text fontWeight="bold">
                      Test Result #{testResult.testNumber}{" "}
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
          }
        )}
      </Accordion>
    </Box>
  );
}
