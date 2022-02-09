import React from "react";
import { TourProvider } from "@reactour/tour";
import type { StepType } from "@reactour/tour";
import type { FunctionComponent } from "react";

export default function WithTour(
  Component: FunctionComponent,
  steps: StepType[]
) {
  return function WrappedComponent() {
    return (
      <TourProvider steps={steps}>
        <Component />;
      </TourProvider>
    );
  };
}
