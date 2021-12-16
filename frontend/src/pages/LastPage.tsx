import { Box, Container, Heading } from "@chakra-ui/react";

export default function LastPage() {
  return (
    <Container>
      <Box bg="white" shadow="md" p="6" rounded="md">
        <Heading fontSize="4xl" textAlign="center" fontWeight="700">
          Thank You!
        </Heading>
      </Box>
    </Container>
  );
}

