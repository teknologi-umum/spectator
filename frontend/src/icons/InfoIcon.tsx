import React from "react";
import type { SVGProps } from "react";

export function InfoIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 12 12" {...props}>
      <g fill="none">
        <path
          d="M11 6A5 5 0 1 1 1 6a5 5 0 0 1 10 0zm-5.5.5V8a.5.5 0 0 0 1 0V6.5a.5.5 0 0 0-1 0zM6 3.75a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5z"
          fill="currentColor"
        ></path>
      </g>
    </svg>
  );
}
