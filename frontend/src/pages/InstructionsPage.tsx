import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Heading, Text, Container, Button, Box } from "@chakra-ui/react";
import Layout from "@/components/Layout";
import { ReactComponent as Arousal } from "@/images/arousal/arousal.svg";
import { useColorModeValue } from "@/hooks";
import { useTranslation } from "react-i18next";
import { SettingsDropdown } from "@/components/Settings";

export default function Instructions() {
  const navigate = useNavigate();
  const { t } = useTranslation("translation", {
    keyPrefix: "translations"
  });
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const fgDarker = useColorModeValue("gray.700", "gray.400", "gray.400");

  const textColor = useColorModeValue("gray.600", "gray.400", "gray.400");

  useEffect(() => {
    document.title = "Instructions | Spectator";
  }, []);

  return (
    <Layout>
      <SettingsDropdown />
      <Container maxW="container.md" bg={bg} p="6" rounded="md" shadow="md">
        <Heading size="lg" textAlign="center" mb="4" color={fg}>
          {t("instructions.title")}
        </Heading>
        <Text fontSize="18" lineHeight="8" color={textColor}>
          {t("instructions.overview")}
        </Text>

        <Heading size="md" mt="6" mb="4" color={fg}>
          1. SAM Test
        </Heading>
        <Text fontSize="18" lineHeight="8" color={textColor}>
          {t("instructions.sam_test")}
        </Text>
        <Box color={fgDarker}>
          <Arousal width="100%" height="100" viewBox="0 0 1240 140" />
        </Box>
        <Text as="label" fontSize="sm" lineHeight="8" color={textColor}>
          {t("instructions.sam_test_label")}
        </Text>

        <Heading size="md" mt="6" mb="4" color={fg}>
          2. Programming Test
        </Heading>
        <Text fontSize="18" lineHeight="8" color={textColor}>
          {t("instructions.programming_test")}
        </Text>
        <Text as="label" fontSize="sm" lineHeight="8" color={textColor}>
          {t("instructions.sam_test_label_2")}
        </Text>

        <Button
          type="submit"
          colorScheme="blue"
          onClick={() => navigate("/sam-test")}
          display="block"
          mx="auto"
          mt="6"
        >
          {t("ui.begin_test")}
        </Button>
      </Container>
    </Layout>
  );
}
