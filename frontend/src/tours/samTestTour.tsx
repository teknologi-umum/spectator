import React from "react";
import type { StepType } from "@reactour/tour";

export const samTestTour: StepType[] = [
  {
    selector: "[data-tour=\"sam-test-step-1\"]",
    content: (
      <p>You can change the theme and the language of the application here.</p>
    )
  },
  {
    selector: "[data-tour=\"sam-test-step-2\"]",
    content: <p>When you&apos;re done, you can press this button to continue.</p>
  }
];
