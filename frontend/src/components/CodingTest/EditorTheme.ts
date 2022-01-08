import { EditorView } from "@codemirror/view";
import { tags, HighlightStyle } from "@codemirror/highlight";
import theme from "@/styles/themes";

export function getHighlightTheme(mode: "light" | "dimmed" | "dark") {
  const c = {
    light: {
      blue: theme.colors.blue[700],
      lightBlue: theme.colors.blue[500],
      red: theme.colors.red[500],
      gray: theme.colors.gray[600],
      orange: theme.colors.orange[500],
      normal: theme.colors.gray[700]
    },
    dimmed: {
      blue: theme.colors.blue[400],
      lightBlue: theme.colors.blue[200],
      red: theme.colors.red[300],
      gray: theme.colors.gray[400],
      orange: theme.colors.orange[300],
      normal: theme.colors.gray[200]
    },
    dark: {
      blue: theme.colors.blue[400],
      lightBlue: theme.colors.blue[200],
      red: theme.colors.red[300],
      gray: theme.colors.gray[400],
      orange: theme.colors.orange[300],
      normal: theme.colors.gray[200]
    }
  };

  return HighlightStyle.define([
    { tag: tags.keyword, color: c[mode].red },
    { tag: tags.number, color: c[mode].blue },
    { tag: tags.bool, color: c[mode].blue },
    { tag: tags.null, color: c[mode].blue },
    { tag: tags.comment, color: c[mode].gray },
    { tag: tags.function, color: theme.colors.pink[600] },
    { tag: tags.string, color: c[mode].lightBlue },
    { tag: tags.propertyName, color: c[mode].blue },
    { tag: tags.punctuation, color: c[mode].gray },
    { tag: tags.paren, color: c[mode].gray },
    { tag: tags.brace, color: c[mode].gray },
    { tag: tags.bracket, color: c[mode].gray },
    { tag: tags.operator, color: c[mode].red },
    { tag: tags.typeName, color: c[mode].orange },
    { tag: tags.constant, color: c[mode].blue },
    { tag: tags.name, color: c[mode].normal }
  ]).extension;
}

export function getEditorTheme({
  mode,
  fontSize
}: {
  mode: "light" | "dimmed" | "dark";
  fontSize: number;
}) {
  const c = {
    light: {
      gray: theme.colors.gray[600],
      bg: theme.colors.white,
      bgDarker: theme.colors.blue[50],
      caret: theme.colors.gray[700]
    },
    dimmed: {
      gray: theme.colors.gray[400],
      bg: theme.colors.gray[700],
      bgDarker: theme.colors.gray[700],
      caret: theme.colors.gray[200]
    },
    dark: {
      gray: theme.colors.gray[400],
      bg: theme.colors.gray[800],
      bgDarker: theme.colors.gray[700],
      caret: theme.colors.gray[200]
    }
  };

  return EditorView.theme(
    {
      "&.cm-editor": {
        backgroundColor: c[mode].bg,
        height: "100%"
      },
      "&.cm-editor.cm-focused": {
        outline: "none"
      },
      ".cm-content": {
        lineHeight: "1.625em",
        verticalAlign: "center",
        fontSize: fontSize + "px",
        height: "100%"
      },
      "&.cm-focused .cm-selectionBackground": {
        background: c[mode].bgDarker
      },
      ".cm-selectionBackground": {
        background: c[mode].bgDarker
      },
      ".cm-cursor": {
        borderLeft: `1.25px solid ${c[mode].caret}`,
        backgroundColor: c[mode].gray
      },
      ".cm-gutters": {
        border: "none",
        backgroundColor: c[mode].bg,
        color: c[mode].gray,
        fontSize: fontSize + "px"
      },
      ".cm-gutterElement.cm-activeLineGutter": {
        backgroundColor: c[mode].bgDarker
      },
      ".cm-line.cm-activeLine": {
        backgroundColor: c[mode].bgDarker
      }
    }
  );
}
