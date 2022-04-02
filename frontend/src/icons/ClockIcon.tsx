import React from "react";
import type { SVGProps } from "react";

export function ClockIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 12 12" {...props}>
      <g fill="none">
        <path
          d="M6 1a5 5 0 1 1 0 10A5 5 0 0 1 6 1zm-.5 2.5A.5.5 0 0 0 5 4v2.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 0-1H6V4a.5.5 0 0 0-.5-.5z"
          fill="currentColor"
        ></path>
      </g>
    </svg>
  );
}
