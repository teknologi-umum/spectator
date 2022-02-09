import React from "react";
import type { SVGProps } from "react";

export function CodeIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 16 16" {...props}>
      <g fill="none">
        <path
          d="M9.905 2.815a.75.75 0 0 1 .38.99l-4 9a.75.75 0 1 1-1.37-.61l4-9a.75.75 0 0 1 .99-.38zM4.498 5.19a.75.75 0 0 1 .063 1.058L3.003 8l1.558 1.752a.75.75 0 1 1-1.122.996l-2-2.25a.75.75 0 0 1 0-.996l2-2.25A.75.75 0 0 1 4.5 5.19zm7.004 0a.75.75 0 0 1 1.059.062l2 2.25a.75.75 0 0 1 0 .996l-2 2.25a.75.75 0 0 1-1.122-.996L12.996 8L11.44 6.248a.75.75 0 0 1 .063-1.058z"
          fill="currentColor"
        ></path>
      </g>
    </svg>
  );
}
