import React from "react";
import { Flex } from "@chakra-ui/react";
import { useTranslation } from "react-i18next";
import { useColorModeValue } from "@/hooks";
import { ThemeButton, LocaleButton } from "../TopBar";

export function SettingsDropdown() {
  const { t } = useTranslation("translation", {
    keyPrefix: "translations"
  });
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");

  return (
    <Flex
      gap={2}
      position="fixed"
      left={4}
      top={4}
      data-tour="step-1"
      zIndex={10}
    >
      <ThemeButton bg={bg} fg={fg} title={t("ui.theme")} />
      <LocaleButton bg={bg} fg={fg} />
    </Flex>
  );
}
