import { Box, useColorMode, useColorModeValue } from "@chakra-ui/react";
import type { FC } from "react";

// eslint-disable-next-line react/prop-types
const Layout: FC = ({ children }) => {
  const bg = useColorModeValue('white', 'gray.800')

  return (
    <Box bg={bg} alignItems="center" w="full" minH="full" py="10" px="4">
      {children}
    </Box>
  );
};

export default Layout;
