import React, { useMemo } from "react";
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
import { useTranslation } from "react-i18next";

const FAKE_FACTS = [
  {
    label: "session_duration",
    value: "02:34:12",
    icon: <ClockIcon width="2rem" height="2rem" />,
    color: "purple"
  },
  {
    label: "coding_test_duration",
    value: "01:28:52",
    icon: <StopwatchIcon width="2rem" height="2rem" />,
    color: "pink"
  },
  {
    label: "words_per_minute",
    value: "120",
    icon: <SpeedIcon width="2rem" height="2rem" />,
    color: "blue"
  },
  {
    label: "deletion_rate",
    value: "1130",
    icon: <BackspaceIcon width="2rem" height="2rem" />,
    color: "red"
  },
  {
    label: "unrelated_keys",
    value: "4380",
    icon: <KeyboardIcon width="2rem" height="2rem" />,
    color: "orange"
  },
  {
    label: "mouse_clicks",
    value: "40",
    icon: <ClickIcon width="2rem" height="2rem" />,
    color: "cyan"
  },
  {
    label: "mouse_scrolls",
    value: "23",
    icon: <ScrollIcon width="2rem" height="2rem" />,
    color: "purple"
  },
  {
    label: "favourite_language",
    value: "Javascript",
    icon: <CodeIcon width="2rem" height="2rem" />,
    color: "blue"
  },
  {
    label: "correct_answers",
    value: "3",
    icon: <CheckmarkIcon width="2rem" height="2rem" />,
    color: "green"
  },
  {
    label: "wrong_answers",
    value: "1",
    icon: <CrossIcon width="2rem" height="2rem" />,
    color: "red"
  },
  {
    label: "unanswered_questions",
    value: "1",
    icon: <QuestionIcon width="2rem" height="2rem" />,
    color: "yellow"
  },
  {
    label: "submission_attempts",
    value: "13",
    icon: <RetryIcon width="2rem" height="2rem" />,
    color: "orange"
  }
];

export default function FunFact() {
  const { t } = useTranslation();
  const gray = useColorModeValue("gray.500", "gray.800", "gray.900");
  const fgDarker = useColorModeValue("gray.700", "gray.300", "gray.400");
  const fakeFacts = useMemo(() => {
    return FAKE_FACTS.map((fact) => {
      return {
        ...fact,
        label: t("translation.translations.funfact." + fact.label + ".title"),
        description: t(
          "translation.translations.funfact." + fact.label + ".description"
        )
      };
    });
  }, [t]);

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
          {t("translation.translations.funfact.description")}
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
        {fakeFacts.map(({ label, value, description, icon, color }, idx) => (
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
