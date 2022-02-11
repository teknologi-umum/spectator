import React from "react";
import type { SVGProps } from "react";

export function SumIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 24 24" {...props}>
      <path
        d="M18 16v2a1 1 0 0 1-1 1H6l6-7l-6-7h11a1 1 0 0 1 1 1v2"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      ></path>
    </svg>
  );
}
