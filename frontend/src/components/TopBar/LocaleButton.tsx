import React from "react";
import { useTranslation } from "react-i18next";
import { IndonesiaFlagIcon, UnitedKingdomFlagIcon } from "@/icons";
import { Flex, MenuItem } from "@chakra-ui/react";
import { MenuDropdown } from "@/components/TopBar";
import { sessionSpoke } from "@/spoke";
import { Locale } from "@/stub/enums";
import { useAppSelector } from "@/store";

interface LocaleButtonProps {
  bg: string;
  fg: string;
}

export default function LocaleButton({ bg, fg, ...rest }: LocaleButtonProps) {
  const { t, i18n } = useTranslation("translation", {
    keyPrefix: "translations.ui"
  });
  const { accessToken } = useAppSelector((state) => state.session);

  return (
    <MenuDropdown
      dropdownWidth="10rem"
      bg={bg}
      fg={fg}
      title={t("language")}
      {...rest}
    >
      <MenuItem
        onClick={async () => {
          if (accessToken === null) return;
          await sessionSpoke.setLocale({ accessToken, locale: Locale.EN });
          i18n.changeLanguage("en-US");
        }}
      >
        <Flex gap={3} align="center">
          <UnitedKingdomFlagIcon width="1.25rem" height="1.25rem" />
          <span>English</span>
        </Flex>
      </MenuItem>
      <MenuItem
        onClick={async () => {
          if (accessToken === null) return;
          await sessionSpoke.setLocale({ accessToken, locale: Locale.ID });
          i18n.changeLanguage("id-ID");
        }}
      >
        <Flex gap={3} align="center">
          <IndonesiaFlagIcon width="1.25rem" height="1.25rem" />
          <span>Indonesia</span>
        </Flex>
      </MenuItem>
    </MenuDropdown>
  );
}
