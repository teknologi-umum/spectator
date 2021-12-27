import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { createStore } from "redux";
import { ChakraProvider, ColorModeScript } from "@chakra-ui/react";
import { extendTheme } from "@chakra-ui/react";
import rootReducer from "@/store/reducers/rootReducer";
import App from "@/App";
import "@fontsource/inter/600.css";
import "@fontsource/inter/700.css";
import "@fontsource/noto-sans";
import "@/index.css";

const store = createStore(rootReducer);

const theme = extendTheme({
  initialColorMode: "light",
  useSystemColorMode: true,
  fonts: {
    heading: "Inter",
    body: "Noto Sans"
  }
});

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
