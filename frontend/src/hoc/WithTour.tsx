import React from "react";
import { TourProvider } from "@reactour/tour";
import type { StepType } from "@reactour/tour";
import type { FunctionComponent } from "react";
import theme from "@/styles/themes";
import { useColorModeValue } from "@/hooks";

export default function WithTour(
  Component: FunctionComponent,
  steps: StepType[]
) {
  const c = theme.colors;

  return function WrappedComponent() {
    const blue = useColorModeValue(c.blue[400], c.blue[300], c.blue[400]);
    const bg = useColorModeValue(c.white, c.gray[700], c.gray[800]);
    const arrowColor = useColorModeValue(c.gray[800], c.gray[400], c.gray[500]);
    const gray = useColorModeValue(c.gray[300], c.gray[600], c.gray[700]);

    return (
      <TourProvider
        steps={steps}
        styles={{
          popover: (base) => ({
            ...base,
            borderRadius: "0.25rem",
            backgroundColor: bg
          }),
          arrow: (base, prop) => ({
            ...base,
            color: prop!.disabled ? gray : arrowColor,
            "&:hover": {
              filter: "brightness(1.2)"
            }
          }),
          badge: () => ({
            display: "none"
          }),
          dot: (base, prop) => ({
            ...base,
            border: "none",
            backgroundColor: prop!.current ? blue : gray
          }),
          close: (base) => ({
            ...base,
            color: gray,
            outline: "none",
            "&:hover": {
              filter: "brightness(1.2)"
            }
          })
        }}
      >
        <Component />;
      </TourProvider>
    );
  };
}
