import React, { FormEvent } from "react";
import { Box, Select } from "@chakra-ui/react";
import { useColorMode } from "@/hooks";
import { useColorModeValue } from "@/hooks/";
import { Theme } from "@/models/Theme";

interface ThemeButtonProps {
  position: "fixed" | "relative";
}

const THEME = ["light", "dimmed", "dark"];

export default function ThemeButton({ position }: ThemeButtonProps) {
  const { colorMode, setColorMode } = useColorMode();

  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fgDarker = useColorModeValue("gray.700", "gray.400", "gray.400");

  return (
    <Box
      display="inline-block"
      position={position}
      left={position === "fixed" ? 4 : "initial"}
      top={position === "fixed" ? 4 : "initial"}
    >
      <Select
        onChange={(e: FormEvent<HTMLSelectElement>) =>
          setColorMode(e.currentTarget.value as Theme)
        }
        bg={bg}
        textTransform="capitalize"
        w="8rem"
        border="none"
        color={fgDarker}
        value={colorMode}
      >
        {THEME.map((val, idx) => (
          <option key={idx} value={val}>
            {val}
          </option>
        ))}
      </Select>
    </Box>
  );
}
