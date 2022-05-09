import { getEditorTheme, getHighlightTheme } from "@/components/CodingTest/editorTheme";
import { useAppSelector } from "@/store";
import { useEffect, useState } from "react";
import { useColorMode } from "@/hooks/";

export function useCodemirrorTheme() {
  const { fontSize } = useAppSelector((state) => state.editor);
  const { colorMode } = useColorMode();
  const [theme, setTheme] = useState(
    getEditorTheme({ mode: colorMode, fontSize })
  );

  const [highlightTheme, setHighlightTheme] = useState(
    getHighlightTheme(colorMode)
  );

  useEffect(() => {
    setTheme(getEditorTheme({ mode: colorMode, fontSize }));
  }, [colorMode, fontSize]);

  useEffect(() => {
    setHighlightTheme(getHighlightTheme(colorMode));
  }, [colorMode, fontSize]);

  return [theme, highlightTheme];
}
