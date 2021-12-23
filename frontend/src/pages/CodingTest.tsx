import {
  Badge,
  Box,
  Button,
  Flex,
  Grid,
  GridItem,
  Heading,
  Text,
  useEventListener
} from "@chakra-ui/react";
import { mouseClickHandler, mouseMoveHandler } from "@/events";
import { useSignalR } from "@/hooks";

// TODO: ini soal ambil dari json atau sejenisnya, jangan langsung tulis disini
export default function CodingTest() {
  const connection = useSignalR("fake_hub_url");

  useEventListener("click", mouseClickHandler(connection));
  useEventListener("mousemove", mouseMoveHandler(connection));

  return (
    <Grid
      w="full"
      h="full"
      gridTemplateColumns="1fr 1fr"
      gridTemplateRows="1fr 1fr"
      bg="gray.100"
      gap="2px"
    >
      <GridItem
        colStart={1}
        colEnd={2}
        rowStart={1}
        rowEnd={3}
        bg="gray.50"
        position="relative"
        h="full"
        resize="horizontal"
        display="flex"
        flexDir="column"
      >
        <Box
          w="full"
          borderBottom="2px"
          borderBottomColor="gray.200"
          bg="white"
          maxH="full"
        >
          <Flex alignItems="center">
            <Box borderRight="2px" borderRightColor="gray.200" p="4">
              <Text fontWeight="bold" fontSize="lg" color="gray.700">
                00:00:00
              </Text>
            </Box>
            <Box p="4">
              <Text fontWeight="bold" fontSize="lg" color="gray.700">
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
            String a : Twinkle twinkle little star \nHow I wonder what you are <br /> String b : Up above the world so high \nLike a diamond in the sky
          </Text>
          
          <Text fontSize="16" lineHeight="6" mt="4">
            Examples:
          </Text>

          <Box as="pre" bg="gray.700" color="white" p="4" mt="4">
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

          <Box as="pre" bg="gray.700" color="white" p="4" mt="4">{`Never gonna give you up
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
          borderTopColor="gray.200"
          bg="white"
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
      <GridItem colStart={2} colEnd={3} rowStart={1} rowEnd={2} bg="white">
        <Heading>Codemirror</Heading>
      </GridItem>
      <GridItem colStart={2} colEnd={3} rowStart={2} rowEnd={3} bg="white">
        <Heading>Visible Test Cases</Heading>
      </GridItem>
    </Grid>
  );
}
