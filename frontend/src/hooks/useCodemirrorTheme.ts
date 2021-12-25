import { getEditorTheme, getHighlightTheme } from "@/components/CodingTest";
import { useColorMode } from "@chakra-ui/react";
import { useEffect, useState } from "react";

export function useCodemirrorTheme() {
  const { colorMode } = useColorMode();
  const [theme, setTheme] = useState(getEditorTheme(colorMode));
  const [highlightTheme, setHighlightTheme] = useState(
    getHighlightTheme(colorMode)
  );

  useEffect(() => {
    setTheme(getEditorTheme(colorMode));
  }, [colorMode]);

  useEffect(() => {
    setHighlightTheme(getHighlightTheme(colorMode));
  }, [colorMode]);

  return [theme, highlightTheme];
}
