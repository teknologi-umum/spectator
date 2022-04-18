import React from "react";
import type { StepType } from "@reactour/tour";
import { store } from "@/store";
import { markTourCompleted } from "@/store/slices/sessionSlice";

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
      <>
        <p>
          When you&apos;re satisfied with your choice, you can press this button
          to continue to the next question.
        </p>
        <p>
          There will be a <kbd>Previous</kbd> on the next page which you can
          press to go to the previous question.
        </p>
        <p>
          You can press the <kbd>Finish</kbd> button to finish the SAM test
          session and go to the Coding Test page.
        </p>
      </>
    ),
    action: () => {
      store.dispatch(markTourCompleted("samTest"));
    }
  }
];
