import { getEditorTheme, getHighlightTheme } from "@/components/CodingTest";
import { useAppSelector } from "@/store";
import { useEffect, useState } from "react";
import type { InitialState as EditorState } from "@/store/slices/editorSlice/types";
import type { InitialState as ThemeState } from "@/store/slices/appSlice/types";
import { useColorMode } from ".";

export function useCodemirrorTheme() {
  const { fontSize } = useAppSelector<EditorState>((state) => state.editor);
  const { currentTheme } = useAppSelector<ThemeState>((state) => state.app);
  const { colorMode } = useColorMode();
  const [theme, setTheme] = useState(
    getEditorTheme({ mode: currentTheme, fontSize })
  );

  const [highlightTheme, setHighlightTheme] = useState(
    getHighlightTheme(currentTheme)
  );

  useEffect(() => {
    setTheme(getEditorTheme({ mode: currentTheme, fontSize }));
  }, [currentTheme, fontSize]);

  useEffect(() => {
    setHighlightTheme(getHighlightTheme(currentTheme));
  }, [currentTheme, fontSize]);

  return [theme, highlightTheme];
}
