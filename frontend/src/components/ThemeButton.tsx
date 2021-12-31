import React from "react";
import { MoonIcon, SunIcon } from "@chakra-ui/icons";
import { Button, useColorMode, useColorModeValue } from "@chakra-ui/react";

interface ThemeButtonProps {
  position: "fixed" | "relative";
}
export default function ThemeButton({ position }: ThemeButtonProps) {
  const { colorMode, toggleColorMode } = useColorMode();
  const bg = useColorModeValue("white", "gray.700");
  const fg = useColorModeValue("gray.800", "gray.100");

  return (
    <Button
      position={position}
      left={position === "fixed" ? 4 : "initial"}
      top={position === "fixed" ? 4 : "initial"}
      onClick={toggleColorMode}
      bg={bg}
      fg={fg}
    >
      {colorMode === "light" ? <MoonIcon /> : <SunIcon />}
    </Button>
  );
}
