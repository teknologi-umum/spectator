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

const FAKE_DATA = [
  {
    title: "words_per_minute",
    color: "blue",
    icon: <SpeedIcon width="48px" height="48px" />,
    value: 90
  },
  {
    title: "deletion_rate",
    color: "red",
    icon: <BackspaceIcon width="48px" height="48px" />,
    value: "4%"
  },
  {
    title: "submission_attempts",
    color: "green",
    icon: <RetryIcon width="48px" height="48px" />,
    value: 69
  }
];

export default function FunFact() {
  const [isLoading, setLoading] = useState(true);
  const { t } = useTranslation();
  const gray = useColorModeValue("gray.500", "gray.800", "gray.900");
  const fgDarker = useColorModeValue("gray.700", "gray.300", "gray.400");

  useEffect(() => {
    const id = setTimeout(() => {
      setLoading(false);
    }, 2500);
    return () => clearTimeout(id);
  }, []);

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
            {FAKE_DATA.map(({ title, color, icon, value }, idx) => (
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
                      {t(`translation.translations.funfact.${title}.title`)}
                    </Heading>
                    <Tooltip
                      label={t(
                        `translation.translations.funfact.${title}.description`
                      )}
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
