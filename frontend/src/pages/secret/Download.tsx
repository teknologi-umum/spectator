import React, { useEffect, useState } from "react";
import { LocaleButton, ThemeButton } from "@/components/TopBar";
import {
  Accordion,
  AccordionButton,
  AccordionIcon,
  AccordionItem,
  AccordionPanel,
  Box,
  Button,
  Flex,
  Heading,
  Text as ChakraText
} from "@chakra-ui/react";
import { useColorModeValue } from "@/hooks";
import { useTranslation } from "react-i18next";
import { useAppDispatch, useAppSelector } from "@/store";
import { LogLevel } from "@microsoft/signalr";
import { loggerInstance } from "@/spoke/logger";
import { useNavigate } from "react-router-dom";
import { removeSessionId } from "@/store/slices/sessionSlice";

interface FakeData {
  id: number;
  studentNumber: string;
}

const FAKE_DATA: FakeData[] = [...Array(20).fill(0)].map((_, n) => ({
  id: n + 1,
  studentNumber: n.toString().padStart(8, "0"),
  fileUrlJson: "https://fake.url.dont.go.here",
  fileUrlCsv: "https://fake.url.dont.go.here"
}));

interface FileEntry {
  csvFileUrl: string;
  jsonFileUrl: string;
  sessionId: string;
  studentNumber: string;
}

interface GroupedFileEntry {
  [key: string]: FileEntry[];
}

// the response from backend is not very humanised, we'll transform this
// so it's easier to read by humans
const HUMANISED_FILENAME: Record<string, string> = {
  solutionaccepted: "Solution Accepted",
  solutionrejected: "Solution Rejected",
  personalinfo: "Personal Info",
  sambefore: "Before Exam SAM Test",
  samafter: "After Exam SAM Test",
  examstarted: "Exam Started",
  examended: "Exam Ended",
  examforfeited: "Exam Forfeited",
  examidereloaded: "IDE Reloaded",
  deadlinepassed: "Deadline Passed",
  mousescrolled: "Mouse Scrolled"
};

const toCapitalised = (str: string) =>
  str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();

function parseFileName(url: string) {
  // example url: /public/51946418_solution_events_solutionaccepted.csv
  return url
    .replace(/\/public\//, "") // remove `/public/` prefix
    .split(".")[0] // remove the file extension
    .split("_") // remove `{sessionId}_` prefix
    .slice(1)
    .map((word) => HUMANISED_FILENAME[word]?.split(" ") || word) // humanise the ending, it's a bit hard for human to read
    .map((word) => {
      if (typeof word === "string") return toCapitalised(word);
      return word.map(toCapitalised).join(" ");
    }) // capitalise
    .join(" ");
}

export default function Download() {
  const [files, setFiles] = useState<GroupedFileEntry>({});
  const { sessionId } = useAppSelector((state) => state.session);
  const { t } = useTranslation();
  const navigate = useNavigate();
  const dispatch = useAppDispatch();

  const boxBg = useColorModeValue("white", "gray.700", "gray.800");
  const bg = useColorModeValue("gray.100", "gray.800", "gray.900");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");

  useEffect(() => {
    (async () => {
      try {
        const response = await fetch(
          import.meta.env.VITE_ADMIN_URL + "/files",
          {
            method: "POST",
            headers: {
              "Content-Type": "application/json"
            },
            body: JSON.stringify({ sessionId })
          }
        );

        // token has expired
        if (response.status === 401) {
          dispatch(removeSessionId());
          navigate("/secret/login");
        }

        const data: FileEntry[] = await response.json();
        const dataGroupedByStudentNumber = data.reduce((acc, curr) => {
          if (acc[curr.studentNumber] === undefined) {
            acc[curr.studentNumber] = [];
          }
          acc[curr.studentNumber].push(curr);
          return acc;
        }, {} as GroupedFileEntry);
        setFiles(dataGroupedByStudentNumber);
      } catch (err) {
        if (err instanceof Error) {
          loggerInstance.log(LogLevel.Error, err.message);
        }
      }
    })();
  }, []);

  return (
    <Flex bg={bg} justifyContent="center" w="full" minH="full" py="10" px="4">
      <Flex gap={2} position="fixed" left={4} top={4} data-tour="step-1">
        <ThemeButton
          bg={boxBg}
          fg={fg}
          title={t("translation.translations.ui.theme")}
        />
        <LocaleButton bg={boxBg} fg={fg} />
      </Flex>
      <Box display="flex" flexDirection="column" alignItems="center">
        <Heading mb="8" size="lg" textAlign="center" fontWeight="700">
          List of Students Data
        </Heading>
        <Accordion
          allowMultiple
          width={"container.md"}
          backgroundColor="white"
          rounded="lg"
        >
          {Object.entries(files).map(([studentNumber, data], index) => (
            <AccordionItem key={index}>
              <AccordionButton
                _hover={{ backgroundColor: "white" }}
                rounded="lg"
              >
                <Box flex="1" textAlign="left" my="2">
                  <ChakraText fontWeight="bold" fontSize="lg">
                    {studentNumber}
                  </ChakraText>
                </Box>
                <AccordionIcon />
              </AccordionButton>
              <AccordionPanel pb={4}>
                {data.map((entry, index) => {
                  const url =
                    entry.jsonFileUrl !== ""
                      ? entry.jsonFileUrl
                      : entry.csvFileUrl;
                  const entryName = parseFileName(url);
                  return (
                    <Flex
                      key={index}
                      justify="space-between"
                      py="4"
                      borderTopWidth={1}
                      borderTopColor={bg}
                    >
                      <span>{entryName}</span>
                      <Flex>
                        {entry.jsonFileUrl !== "" ? (
                          <Button
                            as="a"
                            colorScheme="blue"
                            href={entry.jsonFileUrl}
                            target="_blank"
                            size="sm"
                          >
                            Download JSON
                          </Button>
                        ) : null}
                        {entry.csvFileUrl !== "" ? (
                          <Button
                            as="a"
                            colorScheme="green"
                            href={entry.csvFileUrl}
                            target="_blank"
                            size="sm"
                          >
                            Download CSV
                          </Button>
                        ) : null}
                      </Flex>
                    </Flex>
                  );
                })}
              </AccordionPanel>
            </AccordionItem>
          ))}
        </Accordion>
      </Box>
    </Flex>
  );
}
