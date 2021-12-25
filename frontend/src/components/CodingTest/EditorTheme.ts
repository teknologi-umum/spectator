import { EditorView } from "@codemirror/view";
import { tags, HighlightStyle } from "@codemirror/highlight";
import { theme } from "@chakra-ui/react";

export function getHighlightTheme(mode: "dark" | "light") {
  const isDarkMode = mode === "dark";

  const c = {
    blue: isDarkMode ? theme.colors.blue[400] : theme.colors.blue[700],
    lightBlue: isDarkMode ? theme.colors.blue[200] : theme.colors.blue[500],
    red: isDarkMode ? theme.colors.red[300] : theme.colors.red[500],
    gray: isDarkMode ? theme.colors.gray[400] : theme.colors.gray[600],
    orange: isDarkMode ? theme.colors.orange[300] : theme.colors.orange[500],
    normal: isDarkMode ? theme.colors.gray[200] : theme.colors.gray[700]
  };

  return HighlightStyle.define([
    { tag: tags.keyword, color: c.red },
    { tag: tags.number, color: c.blue },
    { tag: tags.bool, color: c.blue },
    { tag: tags.null, color: c.blue },
    { tag: tags.comment, color: c.gray },
    { tag: tags.function, color: theme.colors.pink[600] },
    { tag: tags.string, color: c.lightBlue },
    { tag: tags.propertyName, color: c.blue },
    { tag: tags.punctuation, color: c.gray },
    { tag: tags.paren, color: c.gray },
    { tag: tags.brace, color: c.gray },
    { tag: tags.bracket, color: c.gray },
    { tag: tags.operator, color: c.red },
    { tag: tags.typeName, color: c.orange },
    { tag: tags.constant, color: c.blue },
    { tag: tags.name, color: c.normal }
  ]).extension;
}

export function getEditorTheme(mode: "dark" | "light") {
  const isDarkMode = mode === "dark";

  const c = {
    bg: isDarkMode ? theme.colors.gray[700] : theme.colors.white,
    bgDarker: isDarkMode ? theme.colors.blue[800] : theme.colors.blue[50],
    caret: isDarkMode ? theme.colors.gray[200] : theme.colors.gray[700]
  };

  return EditorView.theme(
    {
      ".cm-editor": {
        backgroundColor: c.bg
      },
      "&.cm-editor.cm-focused": {
        outline: "none"
      },
      ".cm-content": {
        lineHeight: "1.625em",
        verticalAlign: "center",
        backgroundColor: c.bg
      },
      ".cm-cursor": {
        borderLeft: `1.25px solid ${c.caret}`
      },
      ".cm-gutter": {
        backgroundColor: c.bg
      },
      ".cm-gutterElement.cm-activeLineGutter": {
        backgroundColor: c.bgDarker
      },
      ".cm-line.cm-activeLine": {
        backgroundColor: c.bgDarker
      }
    },
    { dark: isDarkMode }
  );
}
