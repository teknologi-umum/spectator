import {
  Box,
  Button,
  Flex,
  Grid,
  GridItem,
  Heading,
  Text,
  useEventListener,
  useColorModeValue
} from "@chakra-ui/react";
import { mouseClickHandler, mouseMoveHandler } from "@/events";
import { useSignalR } from "@/hooks";

// TODO: ini soal ambil dari json atau sejenisnya, jangan langsung tulis disini
export default function CodingTest() {
  const connection = useSignalR("fake_hub_url");
  useEventListener("click", mouseClickHandler(connection));
  useEventListener("mousemove", mouseMoveHandler(connection));

  const gray = useColorModeValue("gray.100", "gray.700");
  const lightGray = useColorModeValue("gray.50", "gray.800");
  const border = useColorModeValue("gray.200", "gray.600");
  const bg = useColorModeValue("white", "gray.700");
  const fg = useColorModeValue("gray.800", "gray.100");

  return (
    <Grid
      w="full"
      h="full"
      gridTemplateColumns="1fr 1fr"
      gridTemplateRows="1fr 1fr"
      bg={gray}
      gap="2px"
    >
      <GridItem
        colStart={1}
        colEnd={2}
        rowStart={1}
        rowEnd={3}
        bg={lightGray}
        position="relative"
        h="full"
        resize="horizontal"
        display="flex"
        flexDir="column"
      >
        <Box
          w="full"
          borderBottom="2px"
          borderBottomColor={border}
          bg={bg}
          maxH="full"
        >
          <Flex alignItems="center">
            <Box borderRight="2px" borderRightColor={border} p="4">
              <Text fontWeight="bold" fontSize="lg" color={fg}>
                00:00:00
              </Text>
            </Box>
            <Box p="4">
              <Text fontWeight="bold" fontSize="lg" color={fg}>
                Twinkle Twinkle Little Star
              </Text>
            </Box>
          </Flex>
        </Box>
        <Box p="6" maxH="calc(100% - 8rem)" overflowY="auto" flex="1">
          <Text fontSize="16" lineHeight="6">
            Questions:
          </Text>
          <Text fontSize="16" lineHeight="6">
            Print twinkle twinkle little star lyrics by using 2 variables
          </Text>

          <Text fontSize="16" lineHeight="6" mt="4">
            String a : Twinkle twinkle little star \nHow I wonder what you are{" "}
            <br /> String b : Up above the world so high \nLike a diamond in the
            sky
          </Text>

          <Text fontSize="16" lineHeight="6" mt="4">
            Examples:
          </Text>

          <Box as="pre" bg="gray.700" color={fg} p="4" mt="4">
            {`1 variable x: Never gonna give you up
2 variable y: Never gonna run around and desert you
3 variable z: Never gonna let you down
4 call x
5 call z
6 call y
`}
          </Box>

          <Text fontSize="16" lineHeight="6" mt="4">
            Output:
          </Text>

          <Box as="pre" bg="gray.700" color={fg} p="4" mt="4">
            {`Never gonna give you up
Never gonna let you down
Never gonna run around and desert you
`}
          </Box>
        </Box>
        <Flex
          w="full"
          bottom="0"
          justifyContent="end"
          alignItems="center"
          gap="4"
          p="4"
          borderTop="2px"
          borderTopColor={border}
          bg={bg}
        >
          <Button px="6" colorScheme="red">
            Surrender
          </Button>
          <Button px="6" colorScheme="gray">
            Prev
          </Button>
          <Button px="6" colorScheme="blue">
            Next
          </Button>
        </Flex>
      </GridItem>
      <GridItem colStart={2} colEnd={3} rowStart={1} rowEnd={2} bg={bg}>
        <Heading>Codemirror</Heading>
      </GridItem>
      <GridItem colStart={2} colEnd={3} rowStart={2} rowEnd={3} bg={bg}>
        <Heading>Visible Test Cases</Heading>
      </GridItem>
    </Grid>
  );
}
