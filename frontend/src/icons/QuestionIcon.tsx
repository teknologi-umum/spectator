import React from "react";
import type { SVGProps } from "react";

export function QuestionIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 16 16" {...props}>
      <g fill="none">
        <path
          d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2zm0 8.5A.75.75 0 1 0 8 12a.75.75 0 0 0 0-1.5zm0-6a2 2 0 0 0-2 2a.5.5 0 0 0 1 0a1 1 0 0 1 2 0c0 .37-.083.58-.366.898l-.116.125l-.264.27C7.712 8.36 7.5 8.768 7.5 9.5a.5.5 0 0 0 1 0c0-.37.083-.58.366-.898l.116-.125l.264-.27C9.788 7.64 10 7.232 10 6.5a2 2 0 0 0-2-2z"
          fill="currentColor"
        ></path>
      </g>
    </svg>
  );
}
