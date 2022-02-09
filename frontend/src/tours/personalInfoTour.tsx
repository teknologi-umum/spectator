import React from "react";
import type { StepType } from "@reactour/tour";

export const personalInfoTour: StepType[] = [
  {
    selector: "[data-tour=\"personal-info-step-1\"]",
    content: (
      <p>You can change the theme and the language of the application here.</p>
    )
  },
  {
    selector: "[data-tour=\"personal-info-step-2\"]",
    content: <p>When you&apos;re done, you can press this button to continue.</p>
  }
];
