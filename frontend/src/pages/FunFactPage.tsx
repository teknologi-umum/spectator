import React, { useState } from "react";
import { Text, Box, Heading, Flex, Spinner, Tooltip } from "@chakra-ui/react";
import { useEffect } from "react";
import { useColorModeValue } from "@/hooks";
import {
  BackspaceIcon,
  RetryIcon,
  SpeedIcon,
  QuestionOutlineIcon
} from "@/icons";
import { useTranslation } from "react-i18next";
import { useAppSelector } from "@/store";
import { ExamResult_FunFact } from "@/stub/session";
import { removeAccessToken } from "@/store/slices/sessionSlice";

interface FunFactData {
  name: string;
  color: "blue" | "red" | "green";
  icon: JSX.Element;
  value: string;
}

function mapFunFactToList(funfact: ExamResult_FunFact | undefined) {
  if (funfact === undefined) return [];

  const result = Object.entries(funfact).reduce((prev, [key, value]) => {
    switch (key) {
    case "wordsPerMinute":
      return prev.concat({
        name: "words_per_minute",
        color: "blue",
        icon: <SpeedIcon width="48px" height="48px" />,
        value: value
      });
    case "deletionRate":
      return prev.concat({
        name: "deletion_rate",
        color: "red",
        icon: <BackspaceIcon width="48px" height="48px" />,
        value: `${value}%`
      });
    case "submissionAttempts":
      return prev.concat({
        name: "submission_attempts",
        color: "green",
        icon: <RetryIcon width="48px" height="48px" />,
        value: value
      });
    default:
      return prev;
    }
  }, [] as FunFactData[]);

  return result;
}

export default function FunFact() {
  const { examResult } = useAppSelector((state) => state.examResult);
  const { hasPermission } = useAppSelector((state) => state.session);
  const { t } = useTranslation("translation", {
    keyPrefix: "translations.funfact"
  });

  // styles
  const gray = useColorModeValue("gray.500", "gray.200", "gray.100");
  const fgDarker = useColorModeValue("gray.700", "gray.300", "gray.400");

  // local states
  const [funfactData, setFunfactData] = useState<FunFactData[]>([]);
  const [isLoading, setLoading] = useState(true);

  useEffect(() => {
    const id = setTimeout(() => {
      if (examResult === null) throw new Error("examResult is null");
      setFunfactData(mapFunFactToList(examResult.funFact));
      setLoading(false);
      removeAccessToken();
    }, 2500);
    return () => clearTimeout(id);
  }, [hasPermission]);

  useEffect(() => {
    document.title = "Fun Fact | Spectator";
  }, []);

  return (
    <Flex maxW="container.lg" mx="auto" h="full">
      {isLoading ? (
        <Flex
          direction="column"
          align="center"
          justify="center"
          gap="10"
          w="full"
        >
          <Spinner thickness="8px" size="xl" speed="1s" color={gray} />
          <Text fontSize="3xl" color={fgDarker}>
            Calculating your personal result...
          </Text>
        </Flex>
      ) : (
        <Flex gap="8rem" w="full" align="center" justify="center">
          <Heading
            fontSize="5xl"
            textAlign="center"
            fontWeight="800"
            mb="2"
            flex="1"
          >
            Yay, you did it!
          </Heading>
          <Flex gap="4rem" direction="column" flex="1">
            {funfactData.map(({ name: title, color, icon, value }, idx) => (
              <Flex gap="6" align="center" key={idx}>
                <Box
                  p="2"
                  bg={`${color}.200`}
                  color={`${color}.600`}
                  rounded="full"
                >
                  {icon}
                </Box>
                <Box>
                  <Flex align="center" gap="2">
                    <Heading fontSize="2xl" color={gray}>
                      {t(`${title}.title`)}
                    </Heading>
                    <Tooltip
                      label={t(`${title}.description`)}
                      placement="auto-start"
                      padding="3"
                      fontSize="md"
                    >
                      <Text as="span" color={gray}>
                        <QuestionOutlineIcon />
                      </Text>
                    </Tooltip>
                  </Flex>
                  <Text fontSize="4xl" fontWeight="800">
                    {value}
                  </Text>
                </Box>
              </Flex>
            ))}
          </Flex>
        </Flex>
      )}
    </Flex>
  );
}
