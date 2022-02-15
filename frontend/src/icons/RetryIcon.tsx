import React from "react";
import type { SVGProps } from "react";

export function RetryIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 16 16" {...props}>
      <g fill="none">
        <path
          d="M8 3.25A4.75 4.75 0 0 0 3.25 8a.75.75 0 0 1-1.5 0a6.25 6.25 0 1 1 10.587 4.5h.913a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75v-2.5a.75.75 0 0 1 1.5 0v.461A4.75 4.75 0 0 0 8 3.25zM5.75 8a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0zM8 7.25a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5z"
          fill="currentColor"
        ></path>
      </g>
    </svg>
  );
}
