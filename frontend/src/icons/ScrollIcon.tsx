import React from "react";
import type { SVGProps } from "react";

export function ScrollIcon(
  props: SVGProps<SVGSVGElement>
) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 24 24" {...props}>
      <g fill="none">
        <path
          d="M15.75 2A2.25 2.25 0 0 1 18 4.25v15.5A2.25 2.25 0 0 1 15.75 22h-7.5A2.25 2.25 0 0 1 6 19.75V4.25A2.25 2.25 0 0 1 8.25 2h7.5zm-2.036 11.475L12 15.225l-1.718-1.75a.75.75 0 0 0-1.07 1.05l2.253 2.296c.294.3.777.3 1.07 0l2.25-2.296a.75.75 0 1 0-1.07-1.05zm1.072-3.954l-2.25-2.296a.75.75 0 0 0-.987-.075l-.084.075L9.212 9.52a.75.75 0 0 0 .987 1.125l.083-.074L12 8.821l1.714 1.75a.75.75 0 0 0 1.143-.965l-.071-.085l-2.25-2.296l2.25 2.296z"
          fill="currentColor"
        ></path>
      </g>
    </svg>
  );
}
