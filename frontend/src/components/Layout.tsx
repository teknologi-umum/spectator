import React from "react";
import { useColorModeValue } from "@/hooks";
import { Box } from "@chakra-ui/react";
import type { FC } from "react";
import ConnectionToastOverlay from "./ToastOverlay";

const Layout: FC = ({ children }) => {
  const bg = useColorModeValue("gray.100", "gray.800", "gray.900");

  return (
    <>
      <ConnectionToastOverlay />
      <Box bg={bg} alignItems="center" w="full" minH="full" py="10" px="4">
        {children}
      </Box>
    </>
  );
};

export default Layout;
