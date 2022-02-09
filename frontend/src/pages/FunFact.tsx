import React from "react";
import Layout from "@/components/Layout";
import { Text, Box, Grid, Heading, Flex, Tooltip } from "@chakra-ui/react";
import { useEffect } from "react";
import { useColorModeValue } from "@/hooks";
import {
  BackspaceIcon,
  CheckmarkIcon,
  ClickIcon,
  ClockIcon,
  CodeIcon,
  CrossIcon,
  KeyboardIcon,
  QuestionIcon,
  QuestionOutlineIcon,
  RetryIcon,
  ScrollIcon,
  SpeedIcon,
  StopwatchIcon
} from "@/icons";

const FAKE_FACTS = [
  {
    label: "Session Duration",
    value: "02:34:12",
    icon: <ClockIcon width="2rem" height="2rem" />,
    color: "purple",
    description:
      "Session Duration is how long you have been conducting the entire test."
  },
  {
    label: "Coding Test Duration",
    value: "01:28:52",
    icon: <StopwatchIcon width="2rem" height="2rem" />,
    color: "pink",
    description:
      "Coding Test Duration is how long you need to finish the coding test."
  },
  {
    label: "Words Per Minute",
    value: "120",
    icon: <SpeedIcon width="2rem" height="2rem" />,
    color: "blue",
    description:
      "Words Per Minute is how many words you can type in a minute. An average person usually has around 40-60 words per minute rate."
  },
  {
    label: "Deletion Rate",
    value: "1130",
    icon: <BackspaceIcon width="2rem" height="2rem" />,
    color: "red",
    description:
      "Deletion Rate is how many times you press the backspace and delete key for the entire session."
  },
  {
    label: "Unrelated Keys",
    value: "4380",
    icon: <KeyboardIcon width="2rem" height="2rem" />,
    color: "orange",
    description:
      "Unrelated Keys is how many times you press a key that is not related to the test such as typing outside of the code editor and pressing random F-keys."
  },
  {
    label: "Mouse Clicks",
    value: "40",
    icon: <ClickIcon width="2rem" height="2rem" />,
    color: "cyan",
    description:
      "Mouse Clicks is how many times you click the mouse during the entire session."
  },
  {
    label: "Mouse Scrolls",
    value: "23",
    icon: <ScrollIcon width="2rem" height="2rem" />,
    color: "purple",
    description:
      "Mouse Scrolls is how many times you scroll the mouse during the entire session."
  },
  {
    label: "Favourite Language",
    value: "Javascript",
    icon: <CodeIcon width="2rem" height="2rem" />,
    color: "blue",
    description:
      "Favourite Language is the programming language you used the most during the coding test."
  },
  {
    label: "Correct Answers",
    value: "3",
    icon: <CheckmarkIcon width="2rem" height="2rem" />,
    color: "green",
    description:
      "Correct Answers is how many times you answer the coding test correctly."
  },
  {
    label: "Wrong Answers",
    value: "1",
    icon: <CrossIcon width="2rem" height="2rem" />,
    color: "red",
    description:
      "Wrong Answers is how many times you answer the coding test incorrectly."
  },
  {
    label: "Unanswered Questions",
    value: "1",
    icon: <QuestionIcon width="2rem" height="2rem" />,
    color: "yellow",
    description:
      "Words Per Minute is how many words you can type in a minute. An average person usually has around 40-60 words per minute rate."
  },
  {
    label: "Submission Attempts",
    value: "13",
    icon: <RetryIcon width="2rem" height="2rem" />,
    color: "orange",
    description:
      "Submission Attemps is the total number of times you have attempted to submit the test whether it be correct answer or not."
  }
];

export default function FunFact() {
  const gray = useColorModeValue("gray.500", "gray.800", "gray.900");
  const fgDarker = useColorModeValue("gray.700", "gray.300", "gray.400");

  useEffect(() => {
    document.title = "Fun Fact | Spectator";
  }, []);

  return (
    <Layout>
      <Box mb="10">
        <Heading fontSize="5xl" textAlign="center" fontWeight="800" mb="2">
          FUN FACT
        </Heading>
        <Text textAlign="center" fontSize="2xl" color={fgDarker}>
          These are some fun facts about you based on our observation during the
          test.
        </Text>
      </Box>
      <Grid
        w="full"
        align="center"
        gap="4"
        templateColumns="repeat(auto-fit, minmax(18rem, 1fr))"
        maxW="container.xl"
        mx="auto"
      >
        {FAKE_FACTS.map(({ label, value, description, icon, color }, idx) => (
          <Flex
            bg="white"
            key={idx}
            rounded="md"
            shadow="sm"
            gap="2"
            py="12"
            px="3"
            direction="column"
            justify="center"
            align="center"
            position="relative"
          >
            <Box position="absolute" left="4" top="4" color="gray.400">
              <Tooltip
                label={description}
                placement="auto-start"
                padding="3"
                fontSize="md"
              >
                <span>
                  <QuestionOutlineIcon width="1.5rem" height="1.5rem" />
                </span>
              </Tooltip>
            </Box>
            <Box
              p="4"
              rounded="full"
              bg={`${color}.100`}
              color={`${color}.600`}
            >
              {icon}
            </Box>
            <Text fontSize="5xl" fontWeight="800">
              {value}
            </Text>
            <Text color={gray} fontSize="xl" fontWeight="600">
              {label}
            </Text>
          </Flex>
        ))}
      </Grid>
    </Layout>
  );
}
