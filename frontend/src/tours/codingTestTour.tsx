import React from "react";
import { store } from "@/store";
import { markTourCompleted } from "@/store/slices/sessionSlice";
import { TourStepsBuilder } from "./types";

export const codingTestTour: TourStepsBuilder = (t) => [
  {
    selector: "[data-tour=\"sidebar-step-1\"]",
    content: (
      <>
        <p>{t("translation.translations.tour.coding_test.sidebar.first.0")}</p>
        <br />
        <p>{t("translation.translations.tour.coding_test.sidebar.first.1")}</p>
      </>
    ),
    disableActions: true
  },
  {
    selector: "[data-tour=\"sidebar-step-2\"]",
    content: <p>{t("translation.translations.tour.coding_test.sidebar.second")}</p>,
    disableActions: false
  },
  {
    selector: "[data-tour=\"topbar-step-1\"]",
    content: <p>{t("translation.translations.tour.coding_test.topbar.0")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-2\"]",
    content: <p>{t("translation.translations.tour.coding_test.topbar.1")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-3\"]",
    content: <p>{t("translation.translations.tour.coding_test.topbar.2")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-4\"]",
    content: <p>{t("translation.translations.tour.coding_test.topbar.3")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-5\"]",
    content: <p>{t("translation.translations.tour.coding_test.topbar.4")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-6\"]",
    content: <p>{t("translation.translations.tour.coding_test.topbar.5")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-7\"]",
    content: <p>{t("translation.translations.tour.coding_test.topbar.6")}</p>
  },
  {
    selector: "[data-tour=\"topbar-step-8\"]",
    content: <p>{t("translation.translations.tour.coding_test.topbar.7")}</p>
  },
  {
    selector: "[data-tour=\"question-step-1\"]",
    content: (
      <p>{t("translation.translations.tour.coding_test.question.first")}</p>
    )
  },
  {
    selector: "[data-tour=\"question-step-2\"]",
    content: (
      <p>{t("translation.translations.tour.coding_test.question.second")}</p>
    )
  },
  {
    selector: "[data-tour=\"editor-step-1\"]",
    content: (
      <>
        <p>{t("translation.translations.tour.coding_test.editor.first")}</p>
        <br />
        <p>{t("translation.translations.tour.coding_test.editor.second")}</p>
      </>
    )
  },
  {
    selector: "[data-tour=\"scratchpad-step-1\"]",
    content: (
      <>
        <p>{t("translation.translations.tour.coding_test.scratchpad.first")}</p>
        <br />
        <p>{t("translation.translations.tour.coding_test.scratchpad.second")}</p>
      </>
    ),
    action: () => {
      store.dispatch(markTourCompleted("codingTest"));
    }
  }
];
