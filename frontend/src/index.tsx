import React, { Suspense } from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";
import {
  Flex,
  ChakraProvider,
  ColorModeScript,
  Heading,
  Spinner
} from "@chakra-ui/react";
import theme from "@/styles/themes";
import { store, persistor } from "@/store";
import App from "@/App";
import "@fontsource/mulish/400.css";
import "@fontsource/mulish/600.css";
import "@fontsource/mulish/700.css";
import "@fontsource/mulish/800.css";
import "@/i18n";

ReactDOM.render(
  <React.StrictMode>
    <ChakraProvider theme={theme}>
      <Suspense
        fallback={
          <Flex align="center" justify="center" w="full" h="full">
            <Spinner size="lg" color="blue.500" />
          </Flex>
        }
      >
        <ColorModeScript initialColorMode={theme.config.initialColorMode} />
        <Provider store={store}>
          {navigator.cookieEnabled ? (
            <PersistGate loading={null} persistor={persistor}>
              <App />
            </PersistGate>
          ) : (
            <Flex align="center" justify="center" w="full" h="full">
              <Heading textAlign="center">
                This app needs cookie access to work properly.
              </Heading>
            </Flex>
          )}
        </Provider>
      </Suspense>
    </ChakraProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
