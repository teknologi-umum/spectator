import React from "react";
import { useTranslation } from "react-i18next";
import { IndonesiaFlagIcon, UnitedKingdomFlagIcon } from "@/icons";
import { Flex, MenuItem } from "@chakra-ui/react";
import { MenuDropdown } from "@/components/CodingTest";

interface LocaleButtonProps {
  bg: string;
  fg: string;
}

export default function LocaleButton({ bg, fg }: LocaleButtonProps) {
  const { t, i18n } = useTranslation();

  return (
    <MenuDropdown
      dropdownWidth="10rem"
      bg={bg}
      fg={fg}
      title={t("translation.translations.ui.language")}
    >
      <MenuItem onClick={() => i18n.changeLanguage("en-US")}>
        <Flex gap={3} align="center">
          <UnitedKingdomFlagIcon width="1.25rem" height="1.25rem" />
          <span>English</span>
        </Flex>
      </MenuItem>
      <MenuItem onClick={() => i18n.changeLanguage("id-ID")}>
        <Flex gap={3} align="center">
          <IndonesiaFlagIcon width="1.25rem" height="1.25rem" />
          <span>Indonesia</span>
        </Flex>
      </MenuItem>
    </MenuDropdown>
  );
}
