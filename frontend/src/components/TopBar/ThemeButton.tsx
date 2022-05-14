import React from "react";
import { MenuItemOption, MenuOptionGroup } from "@chakra-ui/react";
import { MenuDropdown } from "@/components/TopBar";
import { THEMES } from "@/models/Theme";
import type { Theme } from "@/models/Theme";
import { useColorMode } from "@/hooks";

interface LocaleButtonProps {
  title?: string;
  bg: string;
  fg: string;
}

export default function ThemeButton({ bg, fg, title, ...rest }: LocaleButtonProps) {
  const { colorMode, setColorMode } = useColorMode();

  return (
    <MenuDropdown
      dropdownWidth="10rem"
      bg={bg}
      fg={fg}
      title={title || colorMode}
      {...rest}
    >
      <MenuOptionGroup
        type="radio"
        value={colorMode}
        onChange={(value) => setColorMode(value as Theme)}
      >
        {THEMES.map((theme, idx) => (
          <MenuItemOption textTransform="capitalize" key={idx} value={theme}>
            <span>{theme}</span>
          </MenuItemOption>
        ))}
      </MenuOptionGroup>
    </MenuDropdown>
  );
}
