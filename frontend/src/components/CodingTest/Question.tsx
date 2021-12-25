import {
  Box,
  GridItem,
  Heading,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs,
  Text
} from "@chakra-ui/react";

interface QuestionProps {
  bg: string;
  fg: string;
  codeBg: string;
  time: number;
}
export default function Question({ bg, fg, codeBg }: QuestionProps) {
  return (
    <GridItem
      colStart={1}
      colEnd={2}
      rowStart={2}
      rowEnd={4}
      bg={bg}
      position="relative"
      h="full"
      resize="horizontal"
      display="flex"
      flexDir="column"
      rounded="md"
      shadow="md"
    >
      <Tabs h="full">
        <TabList>
          <Tab>Prompt</Tab>
          <Tab>Your Result</Tab>
        </TabList>

        <TabPanels h="34rem">
          <TabPanel p="2" h="full">
            <Box p="4" overflowY="auto" flex="1" h="full">
              <Text fontWeight="bold" fontSize="lg" color={fg} mb="4">
                Twinkle Twinkle Little Star
              </Text>
              <Text fontSize="16" lineHeight="6">
                Questions:
              </Text>
              <Text fontSize="16" lineHeight="6">
                Print twinkle twinkle little star lyrics by using 2 variables
              </Text>

              <Text fontSize="16" lineHeight="6" mt="4">
                String a : Twinkle twinkle little star \nHow I wonder what you
                are <br /> String b : Up above the world so high \nLike a
                diamond in the sky
              </Text>

              <Text fontSize="16" lineHeight="6" mt="4">
                Examples:
              </Text>

              <Box
                as="pre"
                bg={codeBg}
                color={fg}
                p="4"
                mt="4"
                rounded="sm"
              >
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

              <Box as="pre" bg={codeBg} color={fg} p="4" mt="4">
                {`Never gonna give you up
Never gonna let you down
Never gonna run around and desert you
`}
              </Box>
            </Box>
          </TabPanel>
          <TabPanel p="2">
            <Heading>Result</Heading>
          </TabPanel>
        </TabPanels>
      </Tabs>
    </GridItem>
  );
}
