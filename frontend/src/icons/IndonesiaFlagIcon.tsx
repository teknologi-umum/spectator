import React from "react";
import type { SVGProps } from "react";

export function IndonesiaFlagIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 64 64" {...props}>
      <path
        d="M31.8 62c16.6 0 30-13.4 30-30h-60c0 16.6 13.4 30 30 30"
        fill="#f9f9f9"
      ></path>
      <path
        d="M31.8 2c-16.6 0-30 13.4-30 30h60c0-16.6-13.4-30-30-30"
        fill="#ed4c5c"
      ></path>
    </svg>
  );
}
