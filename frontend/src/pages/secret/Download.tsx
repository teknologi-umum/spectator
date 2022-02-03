import Layout from "@/components/Layout";
import ThemeButton from "@/components/ThemeButton";
import { Box, Button, Table, TableCaption, Tbody, Td, Th, Thead, Tr } from "@chakra-ui/react";

export default function Download() {
  return (
    <Layout>
      <ThemeButton position="fixed" />
      <Box
      height="100vh"
      display="flex"
      flexDirection="column"
      justifyContent="center"
      alignItems="center"
    >
      <Box
        width="50%"
        bg="white"
      >
        <Table>
          <Thead>
            <Tr>
              <Th>Student Name</Th>
              <Th>Action</Th>
            </Tr>
          </Thead>
          <Tbody>
            <Tr>
              <Td>Steve</Td>
              <Td>
                <Button bg="blue.400" color="white" >Download</Button>
              </Td>
            </Tr>
            <Tr>
              <Td>Alex</Td>
              <Td>
                <Button bg="blue.400" color="white" >Download</Button>
              </Td>
            </Tr>
            <Tr>
              <Td>Jedidta</Td>
              <Td>
                <Button bg="blue.400" color="white" >Download</Button>
              </Td>
            </Tr>
          </Tbody>
        </Table>
      </Box>
    </Box>
    </Layout>
  )
}