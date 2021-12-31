import { test } from "vitest";
import type { FC } from "react";
import { Provider } from "react-redux";
import { render } from "@testing-library/react";
import { store } from "@/store";
import { Menu } from "@/components/CodingTest";

const reduxWrapper: FC = ({ children }) => (
  <Provider store={store}>{children}</Provider>
);

test("Should show error message when form is invalid", () => {
  const { container } = render(<Menu bg="black" fg="white" />, {
    wrapper: reduxWrapper
  });
  // bruh
  console.log(container);
});
