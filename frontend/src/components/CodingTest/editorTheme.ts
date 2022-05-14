import { EditorView } from "@codemirror/view";
import { tags } from "@lezer/highlight";
import { HighlightStyle, syntaxHighlighting } from "@codemirror/language";
import theme from "@/styles/themes";
import { Extension } from "@codemirror/state";

export function getHighlightTheme(
  mode: "light" | "dimmed" | "dark"
): Extension {
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

  const style = HighlightStyle.define([
    { tag: [tags.keyword, tags.operator], color: c[mode].red },
    {
      tag: [
        tags.number,
        tags.bool,
        tags.null,
        tags.propertyName,
        tags.constant(tags.variableName)
      ],
      color: c[mode].blue
    },
    {
      tag: [
        tags.comment,
        tags.punctuation,
        tags.paren,
        tags.brace,
        tags.bracket
      ],
      color: c[mode].gray
    },
    { tag: tags.function(tags.variableName), color: theme.colors.pink[600] },
    { tag: tags.string, color: c[mode].lightBlue },
    { tag: tags.typeName, color: c[mode].orange },
    { tag: tags.name, color: c[mode].normal }
  ]);

  return syntaxHighlighting(style);
}

interface EditorThemeOption {
  mode: "light" | "dimmed" | "dark";
  fontSize: number;
}

export function getEditorTheme({
  mode,
  fontSize
}: EditorThemeOption): Extension {
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
      bgDarker: theme.colors.gray[600],
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
    },
    { dark: mode === "dark" }
  );
}
