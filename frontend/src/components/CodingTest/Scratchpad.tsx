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
      <Tabs isLazy h="full">
        <TabList>
          <Tab>Scratchpad</Tab>
          <Tab>Output</Tab>
        </TabList>

        <TabPanels h="full">
          <TabPanel p="2" h="full" tabIndex={-1}>
            <CodeMirror
              height="8rem"
              extensions={[highlightTheme, lineNumbers()]}
              theme={theme}
              style={{ height: "calc(100% - 2.75rem)" }}
            />
          </TabPanel>
          <TabPanel p="2" tabIndex={-1}>
            <Heading>Sandbox</Heading>
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Box>
  );
}
