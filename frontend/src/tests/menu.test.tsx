import React from "react";
import { expect, test } from "vitest";
import type { FC } from "react";
import { Provider } from "react-redux";
import { fireEvent, render } from "@testing-library/react";
import { store } from "@/store";
import { Menu } from "@/components/CodingTest";

const reduxWrapper: FC = ({ children }) => (
  <Provider store={store}>{children}</Provider>
);

test("should be able to change editor language using <Select />", () => {
  const { getByTestId } = render(<Menu bg="black" fg="white" />, {
    wrapper: reduxWrapper
  });

  const languageSelect = getByTestId("editor-language-select");

  fireEvent.change(languageSelect, { target: { value: "python" } });

  expect((languageSelect as HTMLSelectElement).value).to.equal("python");
});

test("should be able to change editor fontsize using <Select />", () => {
  const { getByTestId } = render(<Menu bg="black" fg="white" />, {
    wrapper: reduxWrapper
  });

  const languageSelect = getByTestId("editor-fontsize-select");

  fireEvent.change(languageSelect, { target: { value: 18 } });

  expect((languageSelect as HTMLSelectElement).value).to.equal("18");
});
