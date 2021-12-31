import React from "react";
import { Box, useColorModeValue } from "@chakra-ui/react";
import type { FC } from "react";

const Layout: FC = ({ children }) => {
  const bg = useColorModeValue("gray.100", "gray.800");

  return (
    <Box bg={bg} alignItems="center" w="full" minH="full" py="10" px="4">
      {children}
    </Box>
  );
};

export default Layout;
