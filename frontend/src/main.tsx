import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { ChakraProvider, ColorModeScript } from "@chakra-ui/react";
import theme from "@/styles/themes";
import store from "@/store";
import App from "@/App";
import "@fontsource/mulish/400.css";
import "@fontsource/mulish/600.css";
import "@fontsource/mulish/700.css";
import "@/index.css";

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
