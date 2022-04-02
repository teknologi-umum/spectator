import React from "react";
import ThemeButton from "@/components/CodingTest/TopBar/ThemeButton";
import {
  Box,
  Button,
  Flex,
  Heading,
  Table,
  Tbody,
  Td,
  Th,
  Thead,
  Tr
} from "@chakra-ui/react";
import { LocaleButton } from "@/components/CodingTest";
import { useColorModeValue } from "@/hooks";
import { useTranslation } from "react-i18next";

interface FakeData {
  id: number;
  name: string;
}

const FAKE_DATA: FakeData[] = [
  { id: 1, name: "John Doe" },
  { id: 2, name: "Jane Doe" },
  { id: 3, name: "Jack Smith" },
  { id: 4, name: "James Bond" },
  { id: 5, name: "Adrian Smith" },
  { id: 6, name: "Jill Rock" },
  { id: 7, name: "John Doe" },
  { id: 8, name: "Jane Doe" },
  { id: 9, name: "Jack Smith" },
  { id: 10, name: "James Bond" },
  { id: 11, name: "Adrian Smith" },
  { id: 12, name: "Jill Rock" }
];

export default function Download() {
  const { t } = useTranslation();

  const boxBg = useColorModeValue("white", "gray.700", "gray.800");
  const bg = useColorModeValue("gray.100", "gray.800", "gray.900");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const fgDarker = useColorModeValue("gray.700", "gray.400", "gray.400");

  return (
    <Flex
      bg={bg}
      justifyContent="center"
      w="full"
      minH="full"
      py="10"
      px="4"
    >
      <Flex gap={2} position="fixed" left={4} top={4} data-tour="step-1">
        <ThemeButton
          bg={boxBg}
          fg={fg}
          title={t("translation.translations.ui.theme")}
        />
        <LocaleButton bg={boxBg} fg={fg} />
      </Flex>
      <Box display="flex" flexDirection="column" alignItems="center">
        <Heading mb="8" size="lg" textAlign="center" fontWeight="700">
          List of Students Data
        </Heading>
        <Box bg={boxBg} rounded="md" maxH="32rem" overflowY="auto">
          <Table variant="simple" size="md" w="32rem">
            <Thead>
              <Tr>
                <Th color={fgDarker}>Student Name</Th>
                <Th color={fgDarker} isNumeric>Action</Th>
              </Tr>
            </Thead>
            <Tbody>
              {FAKE_DATA.map((data, index) => (
                <Tr key={index}>
                  <Td>{data.name}</Td>
                  <Td isNumeric>
                    <Button colorScheme="blue" size="sm">
                      Download
                    </Button>
                  </Td>
                </Tr>
              ))}
            </Tbody>
          </Table>
        </Box>
      </Box>
    </Flex>
  );
}
