import React from "react";
import { store } from "@/store";
import { markTourCompleted } from "@/store/slices/sessionSlice";
import { TourStepsBuilder } from "./types";

export const samTestTour: TourStepsBuilder = (t) => [
  {
    selector: "[data-tour=\"step-1\"]",
    content: <p>{t("translation.translations.tour.sam_test.first")}</p>
  },
  {
    selector: "[data-tour=\"step-2\"]",
    content: (
      <>
        <p> {t("translation.translations.tour.sam_test.second.0")} </p>
        <p
          dangerouslySetInnerHTML={t(
            "translations.translation.tour.sam_test.second.1"
          )}
        ></p>
        <p
          dangerouslySetInnerHTML={t(
            "translations.translation.tour.sam_test.second.2"
          )}
        ></p>
      </>
    ),
    action: () => {
      store.dispatch(markTourCompleted("samTest"));
    }
  }
];
