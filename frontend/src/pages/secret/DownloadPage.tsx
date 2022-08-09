import React, { useEffect, useState } from "react";
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
  Spinner,
  Text as ChakraText
} from "@chakra-ui/react";
import { useColorModeValue } from "@/hooks";
import { useAppDispatch, useAppSelector } from "@/store";
import { LogLevel } from "@microsoft/signalr";
import { loggerInstance } from "@/spoke/logger";
import { useNavigate } from "react-router-dom";
import { removeSessionId } from "@/store/slices/sessionSlice";
import { ADMIN_BASE_URL, MINIO_URL } from "@/constants";
import { SettingsDropdown } from "@/components/Settings";

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

async function retriggerResult(adminSessionId: string, examSessionId: string): Promise<void> {
  const response = await fetch(
    `${ADMIN_BASE_URL}/retrigger`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        sessionId: adminSessionId,
        examSessionId: examSessionId
      })
    }
  );

  if (!response.ok) {
    throw new Error(`${response.status} ${response.statusText}`);
  }
}

export default function Download() {
  const [isLoading, setLoading] = useState(false);
  const [files, setFiles] = useState<GroupedFileEntry>({});
  const { sessionId } = useAppSelector((state) => state.session);
  const navigate = useNavigate();
  const dispatch = useAppDispatch();

  const bg = useColorModeValue("gray.100", "gray.800", "gray.900");

  useEffect(() => {
    (async () => {
      setLoading(true);
      try {
        const response = await fetch(ADMIN_BASE_URL + "/files", {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({ sessionId })
        });

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
      setLoading(false);
    })();
  }, []);

  return (
    <Flex bg={bg} justifyContent="center" w="full" minH="full" py="10" px="4">
      <SettingsDropdown disableLocaleButton />
      <Box display="flex" flexDirection="column" alignItems="center">
        <Heading mb="8" size="lg" textAlign="center" fontWeight="700">
          List of Students Data
        </Heading>
        {isLoading ? (
          <Spinner />
        ) : (
          <Accordion
            allowMultiple
            width="container.md"
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
                  <Flex
                    key={index - 1}
                    justify="space-between"
                    py="4"
                    borderTopWidth={1}
                    borderTopColor={bg}
                  >
                    <span>Regenerate results (use with caution)</span>
                    <Flex>
                      <Button
                        as="button"
                        colorScheme="orange"
                        size="sm"
                        onClick={() => retriggerResult(sessionId as string, data[0].sessionId)}
                      >
                        Regenerate
                      </Button>
                    </Flex>
                  </Flex>
                  <Flex
                    key={index}
                    justify="space-between"
                    py="4"
                    borderTopWidth={1}
                    borderTopColor={bg}
                  >
                    <span>Video</span>
                    <Flex>
                      <Button
                        as="a"
                        colorScheme="blue"
                        href={`${MINIO_URL}/public/${data[0].sessionId}_video.mp4`}
                        target="_blank"
                        size="sm"
                      >
                        Download Video
                      </Button>
                    </Flex>
                  </Flex>
                  {data.map((entry, index) => {
                    const fileUrl =
                      entry.jsonFileUrl !== null
                        ? entry.jsonFileUrl
                        : entry.csvFileUrl;
                    const entryName = parseFileName(fileUrl);
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
                          {entry.jsonFileUrl !== null ? (
                            <Button
                              as="a"
                              colorScheme="blue"
                              href={MINIO_URL + entry.jsonFileUrl}
                              target="_blank"
                              size="sm"
                            >
                              Download JSON
                            </Button>
                          ) : null}
                          {entry.csvFileUrl !== null ? (
                            <Button
                              as="a"
                              colorScheme="green"
                              href={MINIO_URL + entry.csvFileUrl}
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
        )}
      </Box>
    </Flex>
  );
}
