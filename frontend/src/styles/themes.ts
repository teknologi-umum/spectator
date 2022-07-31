import { extendTheme } from "@chakra-ui/react";

const theme = extendTheme({
  initialColorMode: "system",
  useSystemColorMode: true,
  fonts: {
    heading: "Mulish",
    body: "Mulish"
  },
  styles: {
    global: (props: { colorMode: "light" | "dark" }) => ({
      "html, body, #root": {
        width: "full",
        height: "full"
      },
      "*::-webkit-scrollbar": {
        width: "1rem"
      },
      "*::-webkit-scrollbar-corner": {
        backgroundColor: "rgba(0, 0, 0, 0)"
      },
      "*::-webkit-scrollbar-thumb": {
        border: "0.3rem solid rgba(0, 0, 0, 0)",
        backgroundClip: "padding-box",
        borderRadius: "9999px",
        backgroundColor: props.colorMode === "light" ? "gray.300" : "gray.600"
      }
    })
  },
  colors: {
    white: "#FFFFFF",
    gray: {
      900: "#11151c",
      800: "#1d2430",
      700: "#2D3748",
      600: "#4A5568",
      500: "#5b6a82",
      400: "#A0AEC0",
      300: "#CBD5E0",
      200: "#E2E8F0",
      100: "#EDF2F7"
    },
    blue: {
      700: "#2C5282",
      500: "#3182CE",
      400: "#4299E1",
      200: "#90CDF4",
      50: "#EBF8FF"
    },
    pink: {
      600: "#B83280"
    },
    red: {
      500: "#E53E3E",
      300: "#FC8181"
    },
    orange: {
      500: "#DD6B20",
      300: "#F6AD55"
    }
  }
});

export default theme;
