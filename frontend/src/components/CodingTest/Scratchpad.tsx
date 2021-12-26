import {
  Heading,
  Tabs,
  TabList,
  TabPanels,
  Tab,
  TabPanel,
  Box
} from "@chakra-ui/react";
import CodeMirror from "@uiw/react-codemirror";
import { lineNumbers } from "@codemirror/gutter";
import { useCodemirrorTheme } from "@/hooks";

interface ScratchpadProps {
  bg: string;
}
export default function Scratchpad({ bg }: ScratchpadProps) {
  const [theme, highlightTheme] = useCodemirrorTheme();

  return (
    <Box bg={bg} rounded="md" shadow="md" flex="1" h="full">
      <Tabs isLazy>
        <TabList>
          <Tab>Scratchpad</Tab>
          <Tab>Output</Tab>
        </TabList>

        <TabPanels>
          <TabPanel p="2">
            <CodeMirror
              value=""
              height="8rem"
              extensions={[highlightTheme, lineNumbers()]}
              theme={theme}
            />
          </TabPanel>
          <TabPanel p="2">
            <Heading>Sandbox</Heading>
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Box>
  );
}
