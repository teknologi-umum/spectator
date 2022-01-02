import { getEditorTheme, getHighlightTheme } from "@/components/CodingTest";
import { useAppSelector } from "@/store";
import { useEffect, useState } from "react";
import type { InitialState as EditorState } from "@/store/slices/editorSlice/types";
// import type { InitialState as ThemeState } from "@/store/slices/appSlice/types";
import { useColorMode } from "@/hooks/";

export function useCodemirrorTheme() {
  const { fontSize } = useAppSelector<EditorState>((state) => state.editor);
  // const { currentTheme } = useAppSelector<ThemeState>((state) => state.app);
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
