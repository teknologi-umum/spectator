import { getEditorTheme, getHighlightTheme } from "@/components/CodingTest";
import { useAppSelector } from "@/store";
import { useColorMode } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import type { InitialState as EditorState } from "@/store/slices/editorSlice/types";

export function useCodemirrorTheme() {
  const { fontSize } = useAppSelector<EditorState>((state) => state.editor);
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
