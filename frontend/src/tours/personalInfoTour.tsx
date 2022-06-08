import React from "react";
import { store } from "@/store";
import { markTourCompleted } from "@/store/slices/sessionSlice";
import { TourStepsBuilder } from "./types";

export const personalInfoTour: TourStepsBuilder = (t) => [
  {
    selector: "[data-tour=\"step-1\"]",
    content: <p>{t("personal_info.first")}</p>
  },
  {
    selector: "[data-tour=\"step-2\"]",
    content: <p>{t("personal_info.second")}</p>,
    action: () => {
      store.dispatch(markTourCompleted("personalInfo"));
    }
  }
];
