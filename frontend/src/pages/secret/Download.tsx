import React, { useEffect } from "react";
import { LocaleButton, ThemeButton } from "@/components/TopBar";
import {
  Box,
  Button,
  Flex,
  Heading,
  HStack,
  Table,
  Tbody,
  Td,
  Th,
  Thead,
  Tr
} from "@chakra-ui/react";
import { useColorModeValue } from "@/hooks";
import { useTranslation } from "react-i18next";
import { useAppSelector } from "@/store";
import { LogLevel } from "@microsoft/signalr";
import { loggerInstance } from "@/spoke/logger";

interface FakeData {
  id: number;
  studentNumber: string;
}

const FAKE_DATA: FakeData[] = [...Array(20).fill(0)].map((_, n) => ({
  id: n + 1,
  studentNumber: n.toString().padStart(8, "0"),
  fileUrlJson: "https://fake.url.dont.go.here",
  fileUrlCsv: "https://fake.url.dont.go.here"
}));

export default function Download() {
  const { sessionId } = useAppSelector((state) => state.session);
  const { t } = useTranslation();

  const boxBg = useColorModeValue("white", "gray.700", "gray.800");
  const bg = useColorModeValue("gray.100", "gray.800", "gray.900");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const fgDarker = useColorModeValue("gray.700", "gray.400", "gray.400");

  useEffect(() => {
    (async () => {
      try {
        const response = await fetch(
          import.meta.env.VITE_ADMIN_URL + "/files",
          {
            method: "POST",
            headers: {
              "Content-Type": "application/json"
            },
            body: JSON.stringify({ sessionId })
          }
        );
        const data = await response.json();
        console.log(data);
      } catch (err) {
        if (err instanceof Error) {
          loggerInstance.log(LogLevel.Error, err.message);
        }
      }
    })();
  }, []);

  return (
    <Flex bg={bg} justifyContent="center" w="full" minH="full" py="10" px="4">
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
          <Table variant="simple" size="md" w="42rem">
            <Thead>
              <Tr>
                <Th color={fgDarker} fontSize="sm">
                  Student Number
                </Th>
                <Th color={fgDarker} isNumeric fontSize="sm">
                  Action
                </Th>
              </Tr>
            </Thead>
            <Tbody>
              {FAKE_DATA.map((data, index) => (
                <Tr key={index}>
                  <Td>{data.studentNumber}</Td>
                  <Td isNumeric>
                    <HStack gap="2" justify="end">
                      <Button colorScheme="blue" size="sm">
                        Download JSON
                      </Button>
                      <Button colorScheme="green" size="sm">
                        Download CSV
                      </Button>
                    </HStack>
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
