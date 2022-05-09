import React from "react";
import { expect, test, vi } from "vitest";
import type { FC } from "react";
import { Provider } from "react-redux";
import { fireEvent, render } from "@testing-library/react";
import { store } from "@/store";
import TopBar from "@/components/CodingTest/TopBar";
import { BrowserRouter, Route, Routes } from "react-router-dom";

vi.stubGlobal("Worker", vi.fn(() => ({
  onmessage: vi.fn(),
  postmessage: vi.fn()
})));

const reduxWrapper: FC = ({ children }) => (
  <Provider store={store}>
    <BrowserRouter>
      <Routes>
        <Route index element={children} />
      </Routes>
    </BrowserRouter>
  </Provider>
);

test("should be able to change editor language using <Select />", () => {
  const { getByTestId } = render(<TopBar bg="black" fg="white" />, {
    wrapper: reduxWrapper
  });

  const languageSelect = getByTestId("editor-language-select");

  fireEvent.change(languageSelect, { target: { value: "python" } });

  expect((languageSelect as HTMLSelectElement).value).to.equal("python");
});

test("should be able to change editor fontsize using <Select />", () => {
  const { getAllByTestId } = render(<TopBar bg="black" fg="white" />, {
    wrapper: reduxWrapper
  });

  const languageSelect = getAllByTestId("editor-fontsize-select")[0];

  fireEvent.change(languageSelect, { target: { value: 18 } });

  expect((languageSelect as HTMLSelectElement).value).to.equal("18");
});
