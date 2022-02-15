import React from "react";
import type { SVGProps } from "react";

export function StopwatchIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 16 16" {...props}>
      <g fill="none">
        <path
          d="M5 1.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1-.5-.5zM7.5 15a6 6 0 1 0 0-12a6 6 0 0 0 0 12zm0-10a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-1 0v-4a.5.5 0 0 1 .5-.5zm4.953-2.358a.5.5 0 1 0-.707.707l1.403 1.403a.5.5 0 1 0 .707-.707l-1.403-1.403z"
          fill="currentColor"
        ></path>
      </g>
    </svg>
  );
}
