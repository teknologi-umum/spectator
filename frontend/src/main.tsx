import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import App from "./App";
import { createStore } from "redux";
import rootReducer from "./store/reducers/rootReducer";
import { Provider } from "react-redux";
import { ChakraProvider } from "@chakra-ui/react";

const store = createStore(rootReducer);

ReactDOM.render(
  <React.StrictMode>
    <ChakraProvider>
      <Provider store={store}>
        <App />
      </Provider>
    </ChakraProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
