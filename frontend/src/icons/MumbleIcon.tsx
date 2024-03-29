import React from "react";
import type { SVGProps } from "react";

export function MumbleIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 20 20" {...props}>
      <path
        fill="currentColor"
        d="M3 10a7 7 0 0 1 14 0v1h-3.5a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5H16a2 2 0 0 0 2-2v-6a8 8 0 1 0-16 0v6a2 2 0 0 0 2 2h2.5a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5H3v-1Z"
      ></path>
    </svg>
  );
}
