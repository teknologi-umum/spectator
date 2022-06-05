import React from "react";
import { store } from "@/store";
import { markTourCompleted } from "@/store/slices/sessionSlice";
import { TourStepsBuilder } from "./types";

export const codingTestTour: TourStepsBuilder = (t) => [
  {
    selector: "[data-tour=\"sidebar-step-1\"]",
    content: (
      <>
        <p>{t("sidebar.first.0")}</p>
        <br />
        <p>{t("sidebar.first.1")}</p>
      </>
    )
  },
  {
    selector: "[data-tour=\"sidebar-step-2\"]",
    content: <p>{t("sidebar.second")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-1\"]",
    content: <p>{t("topbar.0")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-2\"]",
    content: <p>{t("topbar.1")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-3\"]",
    content: <p>{t("topbar.2")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-4\"]",
    content: <p>{t("topbar.3")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-5\"]",
    content: <p>{t("topbar.4")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-6\"]",
    content: <p>{t("topbar.5")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-7\"]",
    content: <p>{t("topbar.6")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-8\"]",
    content: <p>{t("topbar.7")}</p>
  },
  {
    selector: "[data-tour=\"question-step-1\"]",
    content: (
      <p>{t("question.first")}</p>
    )
  },
  {
    selector: "[data-tour=\"question-step-2\"]",
    content: (
      <p>{t("question.second")}</p>
    )
  },
  {
    selector: "[data-tour=\"editor-step-1\"]",
    content: (
      <>
        <p>{t("editor.first")}</p>
        <br />
        <p>{t("editor.second")}</p>
      </>
    )
  },
  {
    selector: "[data-tour=\"scratchpad-step-1\"]",
    content: (
      <>
        <p>{t("scratchpad.first")}</p>
        <br />
        <p>{t("scratchpad.second")}</p>
      </>
    ),
    action: () => {
      store.dispatch(markTourCompleted("codingTest"));
    }
  }
];
