import {
  Heading,
  GridItem,
  Tabs,
  TabList,
  TabPanels,
  Tab,
  TabPanel
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
    <GridItem
      colStart={2}
      colEnd={3}
      rowStart={3}
      rowEnd={4}
      bg={bg}
      rounded="md"
      shadow="md"
    >
      <Tabs>
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
    </GridItem>
  );
}
