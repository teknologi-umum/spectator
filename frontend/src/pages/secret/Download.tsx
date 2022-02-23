import Layout from "@/components/Layout";
import ThemeButton from "@/components/CodingTest/TopBar/ThemeButton";
import { Box, Button, Flex, Table, TableCaption, Tbody, Td, Th, Thead, Tr } from "@chakra-ui/react";
import { LocaleButton } from "@/components/CodingTest";
import { useColorModeValue } from "@/hooks";
import { useTranslation } from "react-i18next";

export default function Download() {
  const { t } = useTranslation();
  
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const fgDarker = useColorModeValue("gray.700", "gray.400", "gray.400");
  
  return (
    <Layout>
      <Flex gap={2} position="fixed" left={4} top={4} data-tour="step-1">
        <ThemeButton
          bg={bg}
          fg={fg}
          title={t("translation.translations.ui.theme")}
        />
        <LocaleButton bg={bg} fg={fg} />
      </Flex>
      <Box
      height="100vh"
      display="flex"
      flexDirection="column"
      alignItems="center"
    >
      <Box
        width="50%"
        bg={bg}
        rounded="md"
      >
        <Table variant="simple" size="md">
          <Thead>
            <Tr>
              <Th color={fgDarker}>Student Name</Th>
              <Th color={fgDarker}>Action</Th>
            </Tr>
          </Thead>
          <Tbody>
            <Tr>
              <Td>Steve</Td>
              <Td>
                <Button bg="blue.400" color="white" size="sm" >Download</Button>
              </Td>
            </Tr>
            <Tr>
              <Td>Alex</Td>
              <Td>
                <Button bg="blue.400" color="white" size="sm" >Download</Button>
              </Td>
            </Tr>
            <Tr>
              <Td>Jedidta</Td>
              <Td>
                <Button bg="blue.400" color="white" size="sm" >Download</Button>
              </Td>
            </Tr>
          </Tbody>
        </Table>
      </Box>
    </Box>
    </Layout>
  )
}