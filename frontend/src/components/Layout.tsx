import { Box } from "@chakra-ui/react";
import { ReactChild } from "react";

export default function Layout({ children }: { children: ReactChild }) {
  return (
    <Box bg="gray.100" alignItems="center" w="full" minH="full" py="10" px="4">
      {children}
    </Box>
  );
}
