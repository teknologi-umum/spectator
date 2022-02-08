import React from "react";
import { MenuItem } from "@chakra-ui/react";
import { MenuDropdown } from "@/components/CodingTest";
import { THEMES } from "@/models/Theme";
import { useColorMode } from "@/hooks";

interface LocaleButtonProps {
  bg: string;
  fg: string;
}

export default function ThemeButton({ bg, fg }: LocaleButtonProps) {
  const { colorMode, setColorMode } = useColorMode();

  return (
    <MenuDropdown
      dropdownWidth="10rem"
      bg={bg}
      fg={fg}
      title={colorMode}
    >
      {THEMES.map((theme, idx) => (
        <MenuItem textTransform="capitalize" key={idx} onClick={() => setColorMode(theme)}>
          <span>{theme}</span>
        </MenuItem>
      ))}
    </MenuDropdown>
  );
}
