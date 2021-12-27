import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { ChakraProvider, ColorModeScript } from "@chakra-ui/react";
// import { extendTheme, ThemeConfig } from "@chakra-ui/react";
import theme from "@/styles/themes"
import store from "@/store";
import App from "@/App";
import "@fontsource/mulish/400.css";
import "@fontsource/mulish/600.css";
import "@fontsource/mulish/700.css";
import "@/index.css";

// const theme = extendTheme({
//   initialColorMode: "light",
//   useSystemColorMode: true,
//   fonts: {
//     heading: "Mulish",
//     body: "Mulish"
//   },
//   colors: {
//     gray: {
//       900: "#010409",
//       800: "#0D1117"
//     }
//   }
// });

ReactDOM.render(
  <React.StrictMode>
    <ColorModeScript initialColorMode={theme.config.initialColorMode} />
    <ChakraProvider theme={theme}>
      <Provider store={store}>
        <App />
      </Provider>
    </ChakraProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
