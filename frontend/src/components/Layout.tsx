import { Box } from "@chakra-ui/react";

// eslint-disable-next-line react/prop-types
const Layout: React.FC = ({ children }) => {
  return (
    <Box bg="gray.100" alignItems="center" w="full" minH="full" py="10" px="4">
      {children}
    </Box>
  );
};

export default Layout;
