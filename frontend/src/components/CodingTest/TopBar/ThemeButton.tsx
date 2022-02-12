import React from "react";
import { MenuItemOption, MenuOptionGroup } from "@chakra-ui/react";
import { MenuDropdown } from "@/components/CodingTest";
import { THEMES } from "@/models/Theme";
import type { Theme } from "@/models/Theme";
import { useColorMode } from "@/hooks";

interface LocaleButtonProps {
  title?: string;
  bg: string;
  fg: string;
}

export default function ThemeButton({ bg, fg, title }: LocaleButtonProps) {
  const { colorMode, setColorMode } = useColorMode();

  return (
    <MenuDropdown
      dropdownWidth="10rem"
      bg={bg}
      fg={fg}
      title={title || colorMode}
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
