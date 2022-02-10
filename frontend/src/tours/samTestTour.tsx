import React from "react";
import type { StepType } from "@reactour/tour";

export const samTestTour: StepType[] = [
  {
    selector: "[data-tour=\"step-1\"]",
    content: (
      <p>
        You can click on one of these option. Your current choice is highlighted
        with a bright blue outline.
      </p>
    )
  },
  {
    selector: "[data-tour=\"step-2\"]",
    content: (
      <p>
        When you&apos;re satisfied with your choice, you can press this button
        to continue to the next question.
      </p>
    )
  },
  {
    selector: "[data-tour=\"step-3\"]",
    content: <p>You can press this button to go to the previous question.</p>
  },
  {
    selector: "[data-tour=\"step-4\"]",
    content: (
      <p>
        You can press this button to finish the SAM test session and go to the
        Coding Test page.
      </p>
    )
  }
];
