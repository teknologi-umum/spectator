import { Button, Flex, GridItem, Select, Text } from "@chakra-ui/react";
import ThemeButton from "../ThemeButton";

function toReadableTime(seconds: number): string {
  const s = Math.floor(seconds % 60);
  const m = Math.floor((seconds / 60) % 60);
  const h = Math.floor((seconds / (60 * 60)) % 24);
  return [h, m.toString().padStart(2, "0"), s.toString().padStart(2, "0")].join(
    ":"
  );
}

const LANGUAGES = ["javascript", "java", "php", "python", "c++"];

interface MenuProps {
  bg: string;
  fg: string;
  time: number;
}
export default function Menu({ bg, fg, time }: MenuProps) {
  return (
    <GridItem
      colStart={1}
      colEnd={3}
      rowStart={1}
      rowEnd={2}
      display="flex"
      justifyContent="stretch"
      gap="2"
    >
      <Flex
        bg={bg}
        color={fg}
        justifyContent="center"
        alignItems="center"
        h="full"
        px="4"
        rounded="md"
      >
        <Text fontWeight="bold" fontSize="lg">
          {toReadableTime(time)}
        </Text>
      </Flex>
      <ThemeButton position="relative" />
      <Flex alignItems="center" gap="2" w="14rem">
        <Select bg={bg} textTransform="capitalize" w="8rem" border="none">
          {LANGUAGES.map((lang, idx) => (
            <option
              key={idx}
              value={lang}
              style={{ textTransform: "capitalize" }}
            >
              {lang}
            </option>
          ))}
        </Select>
        <Select bg={bg} textTransform="capitalize" w="6rem" border="none">
          {Array(6)
            .fill(0)
            .map((_, idx: number) => (
              <option
                key={idx}
                value={idx + 12}
                style={{ textTransform: "capitalize" }}
              >
                {idx + 12}px
              </option>
            ))}
        </Select>
      </Flex>
      <Flex alignItems="center" gap="2" ml="auto">
        <Button px="4" colorScheme="red" h="full">
          Surrender
        </Button>
        <Button px="4" colorScheme="blue" h="full">
          Submit
        </Button>
      </Flex>
    </GridItem>
  );
}
