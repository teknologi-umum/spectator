import { useColorMode } from "@/hooks";
import { Box, Select} from "@chakra-ui/react";
import { useColorModeValue } from "@/hooks/";
import { FormEvent } from "react";
import { Theme } from "@/store/slices/appSlice/types";

interface ThemeButtonProps {
  position: "fixed" | "relative";
}

const THEME = ["light", "dimmed", "dark"];

export default function ThemeButton({ position }: ThemeButtonProps) {
  const { setColorMode } = useColorMode();
  
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");

  return (
    <Box 
      display="inline-block"
      position={position}
      left={position === "fixed" ? 4 : "initial"}
      top={position === "fixed" ? 4 : "initial"}
      >
      
      <Select
        onChange={(e:FormEvent<HTMLSelectElement>) => setColorMode(e.currentTarget.value as Theme)}
        bg={bg}
        textTransform="capitalize"
        w="8rem"
        border="none"
        color={fg}
      >
        {
          THEME.map((val, idx) => (
            <option 
              key={idx} 
              value={val}
            >
              {val}
            </option>
          ))
        }
      </Select>
    </Box>
  );
}
