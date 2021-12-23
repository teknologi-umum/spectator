import { MoonIcon, SunIcon } from "@chakra-ui/icons";
import { Button, useColorMode } from "@chakra-ui/react";

export default function ThemeButton() {
  const { colorMode, toggleColorMode } = useColorMode();
  return (
    <Button position="fixed" left="4" top="4" onClick={toggleColorMode}>
      {colorMode === "light" ? <MoonIcon /> : <SunIcon />}
    </Button>
  );
}
