import React from "react";
import type { SVGProps } from "react";

export function CrossIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 12 12" {...props}>
      <g fill="none">
        <path
          d="M6 11A5 5 0 1 0 6 1a5 5 0 0 0 0 10zm1.854-6.854a.5.5 0 0 1 0 .708L6.707 6l1.147 1.146a.5.5 0 1 1-.708.708L6 6.707L4.854 7.854a.5.5 0 1 1-.708-.708L5.293 6L4.146 4.854a.5.5 0 1 1 .708-.708L6 5.293l1.146-1.147a.5.5 0 0 1 .708 0z"
          fill="currentColor"
        ></path>
      </g>
    </svg>
  );
}
