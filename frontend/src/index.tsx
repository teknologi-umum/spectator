import React, { Suspense } from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";
import {
  Box,
  ChakraProvider,
  ColorModeScript,
  Heading
} from "@chakra-ui/react";
import theme from "@/styles/themes";
import { store, persistor } from "@/store";
import App from "@/App";
import "@fontsource/mulish/400.css";
import "@fontsource/mulish/600.css";
import "@fontsource/mulish/700.css";
import "@fontsource/mulish/800.css";
import "@/index.css";
import "./i18n";

ReactDOM.render(
  <React.StrictMode>
    <Suspense fallback={<div>Loading....</div>}>
      <ColorModeScript initialColorMode={theme.config.initialColorMode} />
      <ChakraProvider theme={theme}>
        <Provider store={store}>
          {navigator.cookieEnabled ? (
            <PersistGate loading={null} persistor={persistor}>
              <App />
            </PersistGate>
          ) : (
            <Box>
              <Heading textAlign="center" mt="4">
                This app needs cookie access to work properly.
              </Heading>
            </Box>
          )}
        </Provider>
      </ChakraProvider>
    </Suspense>
  </React.StrictMode>,
  document.getElementById("root")
);
