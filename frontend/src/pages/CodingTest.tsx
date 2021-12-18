import {
  Badge,
  Box,
  Button,
  Flex,
  Grid,
  GridItem,
  Heading,
  Text
} from "@chakra-ui/react";

// TODO: ini soal ambil dari json atau sejenisnya, jangan langsung tulis disini
export default function CodingTest() {
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
          p="4"
          borderBottom="2px"
          borderBottomColor="gray.200"
          bg="white"
          maxH="full"
        >
          <Heading fontSize="lg" color="gray.700" display="inline">
            Day 1: Sonar Sweep
          </Heading>
          <Badge colorScheme="green" fontWeight="bold" ml="2">
            Easy
          </Badge>
        </Box>
        <Box p="6" maxH="calc(100% - 8rem)" overflowY="auto" flex="1">
          <Text fontSize="16" lineHeight="6">
            You&apos;re minding your own business on a ship at sea when the
            overboard alarm goes off! You rush to see if you can help.
            Apparently, one of the Elves tripped and accidentally sent the
            sleigh keys flying into the ocean!
          </Text>

          <Text fontSize="16" lineHeight="6" mt="4">
            Before you know it, you&apos;re inside a submarine the Elves keep
            ready for situations like this. It&apos;s covered in Christmas
            lights (because of course it is), and it even has an experimental
            antenna that should be able to track the keys if you can boost its
            signal strength high enough; there&apos;s a little meter that
            indicates the antenna&apos;s signal strength by displaying 0-50
            stars.
          </Text>

          <Text fontSize="16" lineHeight="6" mt="4">
            Your instincts tell you that in order to save Christmas, you&apos;ll
            need to get all fifty stars by December 25th.
          </Text>

          <Text fontSize="16" lineHeight="6" mt="4">
            Collect stars by solving puzzles. Two puzzles will be made available
            on each day in the Advent calendar; the second puzzle is unlocked
            when you complete the first. Each puzzle grants one star. Good luck!
          </Text>

          <Text fontSize="16" lineHeight="6" mt="4">
            As the submarine drops below the surface of the ocean, it
            automatically performs a sonar sweep of the nearby sea floor. On a
            small screen, the sonar sweep report (your puzzle input) appears:
            each line is a measurement of the sea floor depth as the sweep looks
            further and further away from the submarine.
          </Text>

          <Text fontSize="16" lineHeight="6" mt="4">
            For example, suppose you had the following report:
          </Text>

          <Box as="pre" bg="gray.700" color="white" p="4" mt="4">{`199
200
208
210
200
207
240
269
260
263`}</Box>

          <Text fontSize="16" lineHeight="6" mt="4">
            This report indicates that, scanning outward from the submarine, the
            sonar sweep found depths of 199, 200, 208, 210, and so on.
          </Text>

          <Text fontSize="16" lineHeight="6" mt="4">
            The first order of business is to figure out how quickly the depth
            increases, just so you know what you&apos;re dealing with - you
            never know if the keys will get carried into deeper water by an
            ocean current or a fish or something.
          </Text>

          <Text fontSize="16" lineHeight="6" mt="4">
            To do this, count the number of times a depth measurement increases
            from the previous measurement. (There is no measurement before the
            first measurement.) In the example above, the changes are as
            follows:
          </Text>

          <Box
            as="pre"
            bg="gray.700"
            color="white"
            p="4"
            mt="4"
          >{`199 (N/A - no previous measurement)
200 (increased)
208 (increased)
210 (increased)
200 (decreased)
207 (increased)
240 (increased)
269 (increased)
260 (decreased)
263 (increased)`}</Box>

          <Text fontSize="16" lineHeight="6" mt="4">
            In this example, there are 7 measurements that are larger than the
            previous measurement.
          </Text>

          <Text fontSize="16" lineHeight="6" mt="4">
            How many measurements are larger than the previous measurement?
          </Text>
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
