import React from "react";
import type { SVGProps } from "react";

export function SumIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 20 20" {...props}>
      <path
        fill="currentColor"
        d="M3.81 3.706a.75.75 0 0 1 .69-.456h11a.75.75 0 0 1 0 1.5H6.262l4.146 4.308a.75.75 0 0 1 .035 1.001L6.104 15.25H15.5a.75.75 0 0 1 0 1.5h-11a.75.75 0 0 1-.575-1.231L8.86 9.613L3.96 4.52a.75.75 0 0 1-.15-.814Z"
      ></path>
    </svg>
  );
}
