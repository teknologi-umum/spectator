import { Tabs, TabList, TabPanels, Tab, TabPanel, Box } from "@chakra-ui/react";
import CodeMirror from "@uiw/react-codemirror";
import { lineNumbers } from "@codemirror/gutter";
import { javascript } from "@codemirror/lang-javascript";
import { php } from "@codemirror/lang-php";
import { java } from "@codemirror/lang-java";
import { cpp } from "@codemirror/lang-cpp";
import { python } from "@codemirror/lang-python";
import { useCodemirrorTheme } from "@/hooks";

const LANGUAGES = [
  javascript({ typescript: true }),
  php({ plain: true }),
  java(),
  cpp(),
  python()
];

const PLACEHOLDER = `import { calculateDirection } from "@/utils/getMouseDirection";
import { emit } from "@/events/emitter";

// TODO(elianiva): emit position and direction as a single event??
export function mouseMoveHandler(connection: unknown) {
  return async (e: MouseEvent) => {
    const data = {
      event: "mouse",
      value: JSON.stringify({
        x: e.pageX,
        y: e.pageY
      }),
      timestamp: Date.now()
    };

    // only emit if it's actully moving
    const direction = calculateDirection(e);
    if (direction) {
      const data = {
        event: "mouse",
        value: JSON.stringify({ direction }),
        timestamp: Date.now()
      };
    }
  };
}`;

interface EditorProps {
  bg: string;
}

export default function Editor({ bg }: EditorProps) {
  const [theme, highlightTheme] = useCodemirrorTheme();

  return (
    <Box bg={bg} rounded="md" shadow="md" flex="1" h="full">
      <Tabs h="full">
        <TabList>
          <Tab>Your Solution</Tab>
        </TabList>
        <TabPanels h="full">
          <TabPanel p="2" h="full" position="relative" tabIndex={-1}>
            <CodeMirror
              value={PLACEHOLDER}
              extensions={[highlightTheme, lineNumbers(), ...LANGUAGES]}
              theme={theme}
              style={{ height: "calc(100% - 2.75rem)" }}
            />
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Box>
  );
}
